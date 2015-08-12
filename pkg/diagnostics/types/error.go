package types

import (
	"fmt"

	"github.com/openshift/origin/pkg/diagnostics/log"
)

// DiagnosticError is an error created by the diagnostic framework and has a little
// more info than a regular error to make them easier to identify in the receiver.
type DiagnosticError struct {
	ID         string
	LogMessage *log.Message
	Cause      error
}

// Error() method means it conforms to the error interface.
func (e DiagnosticError) Error() string {
	if e.LogMessage != nil {
		return fmt.Sprintf("%v", e.LogMessage)
	}
	if e.Cause != nil {
		return e.Cause.Error()
	}
	return e.ID
}

// Easily determine if an error is in fact a Diagnostic error
func IsDiagnosticError(e error) bool {
	_, ok := e.(DiagnosticError)
	return ok
}

// Is the error a diagnostics error that matches the given ID?
func MatchesDiagError(err error, id string) bool {
	if derr, ok := err.(DiagnosticError); ok && derr.ID == id {
		return true
	}
	return false
}
