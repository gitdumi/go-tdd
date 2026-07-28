package main

import (
	"log"
	"net/http"
	"os"

	print "github.com/gitdumi/go-tdd/pkg/depinj"
)

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	print.Greet(w, "dude")
}

func main() {
	print.Greet(os.Stdout, "Yo")

	log.Fatal(http.ListenAndServe(":1001", http.HandlerFunc(MyGreetHandler)))
}
