package main

import (
	"errors"
	"fmt"
	"sync"
)

var (
	ErrInsufficient = errors.New("insufficient balance")
	ErrConflict     = errors.New("optimistic lock conflict: version mismatch")
)

type BankAccount struct {
	Balance int64
	Version int64 // tracks how many times this account has been modified. used by optimistic locking to detect concurrent changes
	mu      sync.Mutex
}

// Transfer uses pessimistic locking — it assumes conflicts are likely,
// so it locks both accounts BEFORE reading or modifying any data.
// No other goroutine can touch these accounts until we're done.
// Downside: if many goroutines call this, they all queue up waiting for locks.
func (b *BankAccount) Transfer(to *BankAccount, amount int64) error {
	// We need to lock BOTH accounts because we're modifying both balances.
	// Problem: if goroutine A locks alice then bob, and goroutine B locks bob then alice,
	// they deadlock (each waiting for the other's lock forever).
	// Solution: always lock in a deterministic order based on memory address.
	// This guarantees every goroutine acquires locks in the same sequence.
	first, second := b, to
	if fmt.Sprintf("%p", b) > fmt.Sprintf("%p", to) {
		first, second = to, b
	}

	first.mu.Lock()
	defer first.mu.Unlock()
	second.mu.Lock()
	defer second.mu.Unlock()

	// At this point, we hold exclusive access to both accounts.
	// No other goroutine can read or write these balances.
	if b.Balance < amount {
		return ErrInsufficient
	}

	b.Balance -= amount
	to.Balance += amount
	return nil
}

// TransferOptimistic uses optimistic locking — it assumes conflicts are rare.
// Instead of holding locks during the entire operation, it:
//  1. Reads the current state (brief lock)
//  2. Does validation WITHOUT any lock held (other goroutines proceed freely)
//  3. Re-acquires locks and checks if the data changed since we read it
//
// If someone modified the accounts between step 1 and step 3, the version numbers
// won't match, and we return ErrConflict so the caller can retry.
// This is faster than pessimistic locking when conflicts are rare because
// the lock is held for microseconds (just reads/writes), not during validation.
func (b *BankAccount) TransferOptimistic(to *BankAccount, amount int64) error {
	// Step 1: READ — take a snapshot of the current state.
	// We lock briefly just to get a consistent read, then immediately release.
	// This is the key difference from pessimistic: we DON'T hold the lock while thinking.
	b.mu.Lock()
	srcBalance := b.Balance
	srcVersion := b.Version
	b.mu.Unlock()

	to.mu.Lock()
	dstVersion := to.Version
	to.mu.Unlock()

	// Step 2: VALIDATE — no locks held here.
	// While we're doing this check, other goroutines are free to modify accounts.
	// That's intentional — we'll detect their changes via version numbers in step 3.
	if srcBalance < amount {
		return ErrInsufficient
	}

	// Step 3: COMMIT — re-acquire locks and verify nothing changed.
	// Same lock ordering as pessimistic to prevent deadlocks.
	first, second := b, to
	if fmt.Sprintf("%p", b) > fmt.Sprintf("%p", to) {
		first, second = to, b
	}
	first.mu.Lock()
	defer first.mu.Unlock()
	second.mu.Lock()
	defer second.mu.Unlock()

	// This is the "optimistic check": compare current versions against our snapshot.
	// If another goroutine modified either account between step 1 and now,
	// the version will have incremented, and this check fails.
	if b.Version != srcVersion || to.Version != dstVersion {
		return ErrConflict
	}

	// Versions match — no one else touched these accounts. Safe to commit.
	b.Balance -= amount
	b.Version++ // increment so other goroutines' optimistic checks will detect our write
	to.Balance += amount
	to.Version++
	return nil
}

// TransferWithRetry wraps TransferOptimistic with a retry loop.
// This is the standard pattern: optimistic locking expects occasional conflicts,
// so callers must be prepared to retry. If conflicts are too frequent,
// this retry loop becomes expensive and pessimistic locking would be better.
func (b *BankAccount) TransferWithRetry(to *BankAccount, amount int64, maxRetries int) error {
	for attempt := 1; attempt <= maxRetries; attempt++ {
		err := b.TransferOptimistic(to, amount)
		if err == nil {
			return nil
		}
		if errors.Is(err, ErrConflict) {
			fmt.Printf("  ⚠ conflict on attempt %d, retrying...\n", attempt)
			continue
		}
		return err // real error like insufficient funds — retrying won't help
	}
	return fmt.Errorf("optimistic transfer failed after %d retries", maxRetries)
}

func main() {
	alice := &BankAccount{Balance: 1000}
	bob := &BankAccount{Balance: 500}

	fmt.Println("=== Starting Balances ===")
	fmt.Printf("Alice: %d, Bob: %d (total: %d)\n\n", alice.Balance, bob.Balance, alice.Balance+bob.Balance)

	// Pessimistic: 20 goroutines all try to transfer $10 from alice to bob.
	// Each goroutine blocks until it acquires both locks. No retries needed,
	// but throughput is limited because only one transfer runs at a time.
	fmt.Println("=== Pessimistic Locking ===")
	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = alice.Transfer(bob, 10)
		}()
	}
	wg.Wait()

	fmt.Printf("Alice: %d, Bob: %d (total: %d)\n\n",
		alice.Balance, bob.Balance, alice.Balance+bob.Balance)

	// Reset balances for the optimistic test
	alice.Balance = 1000
	alice.Version = 0
	bob.Balance = 500
	bob.Version = 0

	// Optimistic: same 20 goroutines, but now they don't hold locks during validation.
	// Multiple goroutines read the same version, race to commit, but only ONE wins per version.
	// Losers get ErrConflict and retry. Watch the output — you'll see retries happening.
	// Despite retries, the total (alice + bob) is always 1500.
	fmt.Println("=== Optimistic Locking ===")
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := alice.TransferWithRetry(bob, 10, 10)
			if err != nil {
				fmt.Printf("  ❌ transfer failed: %v\n", err)
			}
		}()
	}
	wg.Wait()

	fmt.Printf("Alice: %d, Bob: %d (total: %d)\n",
		alice.Balance, bob.Balance, alice.Balance+bob.Balance)
}
