package math

import "testing"

// func TestAdd(t *testing.T) {
// 	got := Add(4,6)
// 	want := 10

// 	if got != want {
// 		t.Errorf("got %q, wanted %q", got, want)
// 	}
// }

// type addTest struct {
// 	arg1, arg2, expected int
// }


// //table-driven test
// var addTests = []addTest{
// 	addTest{2, 3, 5},
// 	addTest{5, 10, 15},
// 	addTest{2, 6, 8},
// 	addTest{1, 8, 9},
// }



// func TestAddTableDriven(t *testing.T) {
// 	for _, test := range addTests {
// 		if output := Add(test.arg1, test.arg2); output != test.expected {
// 			t.Errorf("Output %q not equal to expected %q", output, test.expected)
// 		}
// 	}
// }

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(4, 6)
	}
}



