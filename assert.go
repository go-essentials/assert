// =====================================================================================================================
// == LICENSE:       Copyright (c) 2025 Kevin De Coninck
// ==
// ==                Permission is hereby granted, free of charge, to any person
// ==                obtaining a copy of this software and associated documentation
// ==                files (the "Software"), to deal in the Software without
// ==                restriction, including without limitation the rights to use,
// ==                copy, modify, merge, publish, distribute, sublicense, and/or sell
// ==                copies of the Software, and to permit persons to whom the
// ==                Software is furnished to do so, subject to the following
// ==                conditions:
// ==
// ==                The above copyright notice and this permission notice shall be
// ==                included in all copies or substantial portions of the Software.
// ==
// ==                THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// ==                EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// ==                OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// ==                NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// ==                HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// ==                WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// ==                FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// ==                OTHER DEALINGS IN THE SOFTWARE.
// =====================================================================================================================

// Package assert defines functions for making assertions in Go's standard testing framework.
package assert

import (
	"slices"
	"testing"
)

const (
	boolFailureMsgTemplate   = "%s = %t, want %t"     // Default template to use for the 'True' and 'False' functions.
	nilFailureMessage        = "%s = %v, want %v"     // Default template to use for the 'Nil' function.
	notNilFailureMessage     = "%s = %v, want NOT %v" // Default template to use for the 'NotNil' function.
	equalFailureMsgTemplate  = "%s = %v, want %v"     // Default template to use for the 'Equal' function.
	equalSFailureMsgTemplate = "%s = %v, want %v"     // Default template to use for the 'EqualS' function.
)

// True compares got against the boolean 'true' value.
// If they aren't equal, tb is marked as failed, and it's execution is terminated.
func True(tb testing.TB, got bool, name string, msg ...any) {
	if !got {
		tb.Helper()

		failTB(tb, got, true, name, boolFailureMsgTemplate, msg...)
	}
}

// False compares got against the boolean 'false' value.
// If they aren't equal, tb is marked as failed, and it's execution is terminated.
func False(tb testing.TB, got bool, name string, msg ...any) {
	if got {
		tb.Helper()

		failTB(tb, got, false, name, boolFailureMsgTemplate, msg...)
	}
}

// Equal compares got against want for equality.
// If they aren't equal, tb is marked as failed, and it's execution is terminated.
func Equal[V comparable](tb testing.TB, got, want V, name string, msg ...any) {
	if got != want {
		tb.Helper()

		failTB(tb, got, want, name, equalFailureMsgTemplate, msg...)
	}
}

// Nil compares got against <nil> for equality.
// If they aren't equal, tb is marked as failed, and it's execution is terminated.
func Nil(tb testing.TB, got any, name string, msg ...any) {
	if got != nil {
		tb.Helper()

		failTB(tb, got, nil, name, nilFailureMessage, msg...)
	}
}

// NotNil compares got against NOT <nil> for equality.
// If they aren't equal, tb is marked as failed, and it's execution is terminated.
func NotNil(tb testing.TB, got any, name string, msg ...any) {
	if got == nil {
		tb.Helper()

		failTB(tb, got, nil, name, notNilFailureMessage, msg...)
	}
}

// EqualS compares got against want for equality.
// If they are not equal, t is marked as failed, and it's execution is terminated.
func EqualS[S ~[]E, E comparable](t testing.TB, got, want S, name string, msg ...any) {
	if !slices.Equal(got, want) {
		t.Helper()

		failTB(t, got, want, name, equalSFailureMsgTemplate, msg...)
	}
}

// Marks tb as failed and terminates its execution.
func failTB[V any](tb testing.TB, got, want V, name, template string, msg ...any) {
	if name != "" {
		tb.Fatalf(template, name, got, want)
	} else {
		tb.Fatalf(msg[0].(string), msg[1:]...)
	}
}
