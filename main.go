package main

import (
	"fmt"
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

// AddAll .
func (h *Hit) AddAll() (err error) {
	if err = exec.Command("git", "add", ".").Run(); err != nil {
		log.Printf("git add error (%v)", err)
	}
	return
}

// Commit .
func (h *Hit) Commit() (err error) {
	if err = exec.Command("git", "commit", "-m", "docs(output) new record").Run(); err != nil {
		log.Printf("git commit -m error (%v)", err)
	}
	return
}

// Push .
func (h *Hit) Push(remote string) (err error) {
	if remote == "" {
		remote = "origin master"
	}
	if err = exec.Command("git", "push", remote).Run(); err != nil {
		log.Printf("git push %s error (%v)", remote, err)
	}
	return
}

func main() {
	filename := "output"
	hit := NewHit(filename)
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	for t := range time.NewTicker(10 * time.Second).C {
		if _, err := f.WriteString(fmt.Sprintf("%s\n", t.String())); err != nil {
			log.Printf("append to file(%s) error(%v)", filename, err)
		}
		hit.AddAll()
		fmt.Println(1)
		hit.Commit()
		fmt.Println(2)
		hit.Push("")
		fmt.Println(3)
	}
}
