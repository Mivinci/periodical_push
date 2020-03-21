package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

var (
	f      os.File
	t      int
	script string
)

func main() {
	flag.Parse()

	// Start task
	go periodicTask()

	shutdown()
}

func periodicTask() {
	for t := range time.NewTicker(time.Duration(t) * time.Hour).C {
		err := exec.Command(fmt.Sprintf("./%s", script)).Run()
		if err != nil {
			log.Println(err, t)
		}
	}
}

func init() {
	f, _ := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	log.SetOutput(f)
	log.Printf("Start periodic task")

	flag.StringVar(&script, "f", "", "script file")
	flag.IntVar(&t, "t", 12, "Period of the task execution")
}

func shutdown() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		fmt.Println("Shutdown safely")
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSTOP:
			f.Close()
			log.Printf("Got a signal %s", s.String())
			return
		case syscall.SIGHUP:
			// TODO reload
		default:
			return
		}
	}
}
