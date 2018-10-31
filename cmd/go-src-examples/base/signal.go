// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package base

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// Interrupted is closed when the go command receives an interrupt signal.
var Interrupted = make(chan struct{})
var signalsToIgnore = []os.Signal{os.Interrupt, syscall.SIGQUIT}

// SignalTrace is the signal to send to make a Go program
// crash with a stack trace.
var SignalTrace os.Signal = syscall.SIGQUIT

// processSignals setups signal handler.
func processSignals() {
	sig := make(chan os.Signal)
	signal.Notify(sig, signalsToIgnore...)
	go func() {
		<-sig
		close(Interrupted)
	}()
}

var onceProcessSignals sync.Once

// StartSigHandlers starts the signal handlers.
func StartSigHandlers() {
	onceProcessSignals.Do(processSignals)
}
