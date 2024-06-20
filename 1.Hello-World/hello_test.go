// hello_test.go
package main

import "testing"


func TestHello(t * testing.T)  {
    got:= Hello("example")
    want:= "Hello, example"
    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}

// func TestHello(t *testing.T) {
//     got := Hello("arg")
//     want := "Hello, arg"
//     if got != want {
//         t.Errorf("got %q  want %q", got, want)
//     }
// }
