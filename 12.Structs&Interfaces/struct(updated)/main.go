package main

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	port string
}

type User struct {
	ID int
	Name string
}

func (s *Server) RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to server running on http://localhost%s", s.port )
}

func (s *Server) UserHadler(w http.ResponseWriter, r *http.Request) {
	user := User{ID: 1, Name: "ayush"}

	fmt.Fprintf(w, "User: %d, %s", user.ID, user.Name)
}

func main() {
	srv := &Server{port: ":8080"}
	http.HandleFunc("/", srv.RootHandler)
	http.HandleFunc("/user", srv.UserHadler)

	log.Printf("Server starting on %s", srv.port)

	log.Fatal(http.ListenAndServe(srv.port, nil))
}






























// type laptop struct {
//     cpu string
//     ram int
//     storage int
//     manufacturer string
// }

// pointer reciever
// func (l *laptop) upgradeStorage(size int, ramSize int) {
//     l.storage += size
// 	l.ram += ramSize
// }



// func main() {
//     thinkpad := laptop{"i5-1240p", 16, 256, "Lenovo"}
//     fmt.Println(thinkpad.storage)
//     thinkpad.upgradeStorage(100, 8)
//     fmt.Println(thinkpad.storage)
// 	fmt.Println(thinkpad.ram)

// }



// 1. Value Receiver
// Copy of the struct is passed.
// Changes made inside the method do not affect the original struct.
// Used when:
// The method does not modify the struct.
// Struct is small (no performance overhead).


// 2. Pointer Receiver
// The method receives a pointer to the struct (*Struct).
// Changes made affect the original struct.
// Used when:
// You need to modify the struct.
// Struct is large (avoid copying).