package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/tadvi/winserv"
)

func main() {
	path, _ := os.Executable()
	dir := filepath.Dir(path)

	f, err := os.OpenFile(filepath.Join(dir, "log.txt"), os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return
	}
	log.SetOutput(f)

	done := make(chan bool, 1)

	winserv.OnExit(func() {
		// This runs in separate goroutine from main.
		// It's easy to get into race conditions here.
		log.Println("Exit")
		f.Sync()
		f.Close()
		done <- true
	})

	log.Println("Started with", dir)

	select {
	case <-done:
		break
	}
}
