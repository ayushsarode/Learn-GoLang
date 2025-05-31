package main

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"strings"
)

func hashString(input string) string {
	sha := sha1.New()
	io.WriteString(sha, input)
	return fmt.Sprintf("%x", sha.Sum(nil))
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter password: ")

	pass1, _ := reader.ReadString('\n');
	pass1 = strings.TrimSpace(pass1)


	fmt.Print("Re-enter password: ")
	pass2, _ := reader.ReadString('\n')
	pass2 = strings.TrimSpace(pass2)



	hash1 := hashString(pass1)
	hash2 := hashString(pass2)

	

	

	if hash1 == hash2 {
		fmt.Printf("Correct password\n")
		
	
	} else {
		fmt.Println("Password do not match")
		
	}

	fmt.Printf("% x\n", hash1)
	fmt.Printf("% x", hash2)

}