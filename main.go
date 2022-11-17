package main

import (
	// "fmt"
	"log"
	// "net/http"
	"os/exec"
	"time"
)

func eipsClear() error {
	cmd := exec.Command("/usr/sbin/eips", "-c")
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func eipsWrite(col, row, message string) error {
	cmd := exec.Command("/usr/sbin/eips", string(col), string(row), message)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func main() {
	log.Println("Booting up")

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, "Hello from inside a kindle!")
	// })
	//
	// log.Fatal(http.ListenAndServe(":40317", nil))

	log.Println("Clearing screen")
	err := eipsClear
	if err != nil {
		log.Println("clear exited with err")
		log.Fatal(err)
	}
	log.Println("Screen cleared")

	time.Sleep(1 * time.Second)

	log.Println("Saying hi")
	err2 := eipsWrite("5", "5", "Hello World from caius/K4-room-sign")
	if err2 != nil {
		log.Println("Hi exited with err")
		log.Fatal(err2)
	}
	log.Println("Goodbye")
}
