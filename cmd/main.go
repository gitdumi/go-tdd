package main

import (
	"net/http"
	"os"
	"time"

	print "github.com/gitdumi/go-tdd/pkg/depinj"
	m "github.com/gitdumi/go-tdd/pkg/mocking"
)

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	print.Greet(w, "dude")
}

func main() {
	// print.Greet(os.Stdout, "Yo")
	// log.Fatal(http.ListenAndServe(":1001", http.HandlerFunc(MyGreetHandler)))

	sleeper := &m.ConfigurableSleeper{Duration: 2 * time.Second, CustomSleep: time.Sleep}
	m.Countdown(os.Stdout, sleeper)
}
