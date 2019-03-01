// Package allows running binary as Windows service.
// Set OnExit function to execute clean up and service shutdown.
//
// Copyright (C) 2019 Winserv Authors. All Rights Reserved.
// Copyright (C) 2015 Daniel Theophanes. All Rights Reserved.

// +build windows

package winserv

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	"golang.org/x/sys/windows/svc"
)

type runner struct {
	sync.Mutex
	onExit func()
}

// Interactive determines if calling process is running interactively.
var Interactive bool

var service runner

func init() {
	var err error
	Interactive, err = svc.IsAnInteractiveSession()
	if err != nil {
		panic(err)
	}
	if Interactive {
		return
	}
	go func() {
		_ = svc.Run("", service)

		service.Lock()
		fn := service.onExit
		service.Unlock()

		// Don't hold this lock in user code.
		if fn != nil {
			fn()
		}
		os.Exit(0)
	}()
}

// OnExit shutdown function.
func OnExit(fn func()) {
	service.Lock()
	service.onExit = fn
	service.Unlock()
}

// Execute gets called if binary is executed as Windows Service.
func (runner) Execute(args []string, r <-chan svc.ChangeRequest, changes chan<- svc.Status) (bool, uint32) {
	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown
	changes <- svc.Status{State: svc.StartPending}

	changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}
	for {
		c := <-r
		switch c.Cmd {
		case svc.Interrogate:
			changes <- c.CurrentStatus
		case svc.Stop, svc.Shutdown:
			changes <- svc.Status{State: svc.StopPending}
			return false, 0
		}
	}
	return false, 0
}

// RedirectLog helps with logging into file output while debugging Windows Service.
func RedirectLog(file string) (*os.File, error) {
	f, err := os.OpenFile(filepath.Join(file), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	log.SetOutput(f)
	return f, nil
}
