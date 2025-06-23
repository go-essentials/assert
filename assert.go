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

// True compares got against the boolean 'true' value.
// If they aren't equal, tb is marked as failed, and it's execution is terminated.
func True(tb testing.TB, got bool, name string) {
	tb.Helper()

	if !got {
		tb.Fatalf("%s = %t, want %t", name, got, true)
	}
}

// Truef compares got against the boolean 'true' value.
// If they aren't equal, tb is marked as failed, and it's execution is terminated.
// The error message is formatted using fmt.Sprintf with the provided format and args.
func Truef(tb testing.TB, got bool, format string, args ...any) {
	tb.Helper()

	if !got {
		tb.Fatalf(format, args...)
	}
}

// False compares got against the boolean 'false' value.
// If they aren't equal, tb is marked as failed, and it's execution is terminated.
func False(tb testing.TB, got bool, name string) {
	tb.Helper()

	if got {
		tb.Fatalf("%s = %t, want %t", name, got, false)
	}
}

// Falsef compares got against the boolean 'false' value.
// If they aren't equal, tb is marked as failed, and it's execution is terminated.
// The error message is formatted using fmt.Sprintf with the provided format and args.
func Falsef(tb testing.TB, got bool, format string, args ...any) {
	tb.Helper()

	if got {
		tb.Fatalf(format, args...)
	}
}

// Equal compares got against want for equality.
// If they aren't equal, tb is marked as failed, and it's execution is terminated.
func Equal[V comparable](tb testing.TB, got, want V, name string) {
	tb.Helper()

	if got != want {
		tb.Fatalf("%s = %v, want %v", name, got, want)
	}
}

// Equalf compares got against want for equality.
// If they aren't equal, tb is marked as failed, and it's execution is terminated.
// The error message is formatted using fmt.Sprintf with the provided format and args.
func Equalf[V comparable](tb testing.TB, got, want V, format string, args ...any) {
	tb.Helper()

	if got != want {
		tb.Fatalf(format, args...)
	}
}

// EqualFn compares got against want for equality using a custom comparison function.
// If they aren't equal, tb is marked as failed, and it's execution is terminated.
func EqualFn[V comparable](tb testing.TB, got, want V, cmpFn func(V, V) bool, name string) {
	tb.Helper()

	if !cmpFn(got, want) {
		tb.Fatalf("%s = %v, want %v", name, got, want)
	}
}

// EqualFnf compares got against want for equality using a custom comparison function.
// If they aren't equal, tb is marked as failed, and it's execution is terminated.
// The error message is formatted using fmt.Sprintf with the provided format and args.
func EqualFnf[V comparable](tb testing.TB, got, want V, cmpFn func(V, V) bool, format string, args ...any) {
	tb.Helper()

	if !cmpFn(got, want) {
		tb.Fatalf(format, args...)
	}
}

// Nil compares got against <nil> for equality.
// If they aren't equal, tb is marked as failed, and it's execution is terminated.
func Nil(tb testing.TB, got any, name string) {
	tb.Helper()

	if got != nil {
		tb.Fatalf("%s = %v, want %v", name, got, "<nil>")
	}
}

// Nilf compares got against <nil> for equality.
// If they aren't equal, tb is marked as failed, and it's execution is terminated.
// The error message is formatted using fmt.Sprintf with the provided format and args.
func Nilf(tb testing.TB, got any, format string, args ...any) {
	tb.Helper()

	if got != nil {
		tb.Fatalf(format, args...)
	}
}

// NotNil compares got against NOT <nil> for equality.
// If they aren't equal, tb is marked as failed, and it's execution is terminated.
func NotNil(tb testing.TB, got any, name string) {
	tb.Helper()

	if got == nil {
		tb.Fatalf("%s = %v, want NOT %v", name, got, "<nil>")
	}
}

// NotNilf compares got against NOT <nil> for equality.
// If they aren't equal, tb is marked as failed, and it's execution is terminated.
// The error message is formatted using fmt.Sprintf with the provided format and args.
func NotNilf(tb testing.TB, got any, format string, args ...any) {
	tb.Helper()

	if got == nil {
		tb.Fatalf(format, args...)
	}
}

// EqualS compares got against want for equality.
// If they are not equal, tb is marked as failed, and it's execution is terminated.
func EqualS[S ~[]E, E comparable](tb testing.TB, got, want S, name string) {
	tb.Helper()

	if !slices.Equal(got, want) {
		tb.Fatalf("%s = %v, want %v", name, got, want)
	}
}

// EqualSf compares got against want for equality.
// If they are not equal, tb is marked as failed, and it's execution is terminated.
// The error message is formatted using fmt.Sprintf with the provided format and args.
func EqualSf[S ~[]E, E comparable](tb testing.TB, got, want S, format string, args ...any) {
	tb.Helper()

	if !slices.Equal(got, want) {
		tb.Fatalf(format, args...)
	}
}

// Fail marks tb as failed and terminates its execution.
func Fail(tb testing.TB, name string) {
	tb.Helper()

	tb.Fatalf("%s", name)
}

// Failf marks tb as failed and terminates its execution.
// The error message is formatted using fmt.Sprintf with the provided format and args.
func Failf(tb testing.TB, format string, args ...any) {
	tb.Helper()

	tb.Fatalf(format, args...)
}
