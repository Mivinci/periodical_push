package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"
)

// Hit .
type Hit struct {
	filename string
}

// NewHit .
func NewHit(filename string) *Hit {
	return &Hit{filename: filename}
}

// Add .
func (h *Hit) Add() (err error) {
	err = exec.Command("git", "add", ".").Run()
	log.Printf("git add error %s", err)
	return
}

func main() {
	hit := NewHit("")
	for t := range time.NewTicker(2 * time.Second).C {
		if err := ioutil.WriteFile("output", []byte(t.String()), os.ModePerm); err != nil {
			log.Println(err)
		}
		hit.Add()
	}
}