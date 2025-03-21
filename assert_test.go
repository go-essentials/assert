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

// Quality assurance: Verify (and measure the performance) of the public API of the "assert" package.
package assert_test

import (
	"fmt"
	"testing"

	"github.com/go-essentials/assert"
)

// The testableT wraps the testing.T struct and adds a field for storing the failure message.
type testableT struct {
	testing.TB
	failureMsg string
}

// Fatalf formats args using fmt.Sprintf and stores the result in t.
func (t *testableT) Fatalf(format string, args ...any) {
	t.failureMsg = fmt.Sprintf(format, args...)
}

// UT: Compare a value against the boolean 'true' value.
func TestTrue(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	for tcName, tc := range map[string]struct {
		gInput    bool
		nameInput string
		want      string
	}{
		"When 'got' is 'false'.": {
			gInput:    false,
			nameInput: "IsDigit(\"0\")",
			want:      "IsDigit(\"0\") = false, want true",
		},
		"When 'got' is 'true'.": {
			gInput:    true,
			nameInput: "IsDigit(\"0\")",
		},
	} {
		tc := tc // Rebind 'tc'. Note: This is required to support "parallel" execution.

		// EXECUTION.
		t.Run(tcName, func(t *testing.T) {
			t.Parallel() // Enable parallel execution.

			// ARRANGE.
			testingT := &testableT{TB: t}

			// ACT.
			assert.True(testingT, tc.gInput, tc.nameInput)

			// ASSERT.
			if testingT.failureMsg != tc.want {
				t.Fatalf("Failure message = \"%s\", want \"%s\"", testingT.failureMsg, tc.want)
			}
		})
	}

	// EXECUTION.
	t.Run("When using a custom failure message.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// ARRANGE.
		testingT := &testableT{TB: t}

		// ACT.
		assert.True(testingT, false, "", "UT Failed: `IsDigit(\"0\")` - got %t, want %t.", false, true)

		// ASSERT.
		if testingT.failureMsg != "UT Failed: `IsDigit(\"0\")` - got false, want true." {
			t.Fatalf("Failure message = \"%s\", want \"%s\"", testingT.failureMsg, "UT Failed: `IsDigit(\"0\")` - got false, want true.")
		}
	})
}

// UT: Compare a value against the boolean 'false' value.
func TestFalse(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	for tcName, tc := range map[string]struct {
		gInput    bool
		nameInput string
		want      string
	}{
		"When 'got' is 'true'.": {
			gInput:    true,
			nameInput: "IsLetter(\"0\")",
			want:      "IsLetter(\"0\") = true, want false",
		},
		"When 'got' is 'false'.": {
			gInput:    false,
			nameInput: "IsLetter(\"0\")",
		},
	} {
		tc := tc // Rebind 'tc'. Note: This is required to support "parallel" execution.

		// EXECUTION.
		t.Run(tcName, func(t *testing.T) {
			t.Parallel() // Enable parallel execution.

			// ARRANGE.
			testingT := &testableT{TB: t}

			// ACT.
			assert.False(testingT, tc.gInput, tc.nameInput)

			// ASSERT.
			if testingT.failureMsg != tc.want {
				t.Fatalf("Failure message = \"%s\", want \"%s\"", testingT.failureMsg, tc.want)
			}
		})
	}

	// EXECUTION.
	t.Run("When using a custom failure message.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// ARRANGE.
		testingT := &testableT{TB: t}

		// ACT.
		assert.False(testingT, true, "", "UT Failed: `IsLetter(\"0\")` - got %t, want %t.", true, false)

		// ASSERT.
		if testingT.failureMsg != "UT Failed: `IsLetter(\"0\")` - got true, want false." {
			t.Fatalf("Failure message = \"%s\", want \"%s\"", testingT.failureMsg, "UT Failed: `IsLetter(\"0\")` - got true, want false.")
		}
	})
}

// UT: Compare 2 values for equality.
func TestEqual(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	for tcName, tc := range map[string]struct {
		gInput, wInput bool
		nameInput      string
		want           string
	}{
		"When 'got' and 'want' are different.": {
			gInput: false, wInput: true,
			nameInput: "IsDigit(\"0\")",
			want:      "IsDigit(\"0\") = false, want true",
		},
		"When 'got' and 'want' are equal.": {
			gInput: true, wInput: true,
			nameInput: "IsDigit(\"0\")",
		},
	} {
		tc := tc // Rebind 'tc'. Note: This is required to support "parallel" execution.

		// EXECUTION.
		t.Run(tcName, func(t *testing.T) {
			t.Parallel() // Enable parallel execution.

			// ARRANGE.
			testingT := &testableT{TB: t}

			// ACT.
			assert.Equal(testingT, tc.gInput, tc.wInput, tc.nameInput)

			// ASSERT.
			if testingT.failureMsg != tc.want {
				t.Fatalf("Failure message = \"%s\", want \"%s\"", testingT.failureMsg, tc.want)
			}
		})
	}

	// EXECUTION.
	t.Run("When using a custom failure message.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// ARRANGE.
		testingT := &testableT{TB: t}

		// ACT.
		assert.Equal(testingT, false, true, "", "UT Failed: `IsDigit(\"0\")` - got %t, want %t.", false, true)

		// ASSERT.
		if testingT.failureMsg != "UT Failed: `IsDigit(\"0\")` - got false, want true." {
			t.Fatalf("Failure message = \"%s\", want \"%s\"", testingT.failureMsg, "UT Failed: `IsDigit(\"0\")` - got false, want true.")
		}
	})
}

// UT: Compare a value against '<nil>'.
func TestNil(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	for tcName, tc := range map[string]struct {
		gInput    any
		nameInput string
		want      string
	}{
		"When 'got' is a number.": {
			gInput:    1,
			nameInput: "Update()",
			want:      "Update() = 1, want <nil>",
		},
		"When 'got' is '<nil>'.": {
			gInput:    nil,
			nameInput: "Update()",
		},
	} {
		tc := tc // Rebind 'tc'. Note: This is required to support "parallel" execution.

		// EXECUTION.
		t.Run(tcName, func(t *testing.T) {
			t.Parallel() // Enable parallel execution.

			// ARRANGE.
			testingT := &testableT{TB: t}

			// ACT.
			assert.Nil(testingT, tc.gInput, tc.nameInput)

			// ASSERT.
			if testingT.failureMsg != tc.want {
				t.Fatalf("Failure message = \"%s\", want \"%s\"", testingT.failureMsg, tc.want)
			}
		})
	}

	// EXECUTION.
	t.Run("When using a custom failure message.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// ARRANGE.
		testingT := &testableT{TB: t}

		// ACT.
		assert.Nil(testingT, 1, "", "UT Failed: `Update()` - got %v, want <nil>.", 1)

		// ASSERT.
		if testingT.failureMsg != "UT Failed: `Update()` - got 1, want <nil>." {
			t.Fatalf("Failure message = \"%s\", want \"%s\"", testingT.failureMsg, "UT Failed: `Update()` - got 1, want <nil>.")
		}
	})
}

// UT: Compare a value against NOT '<nil>'.
func TestNotNil(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	for tcName, tc := range map[string]struct {
		gInput    any
		nameInput string
		want      string
	}{
		"When 'got' is <nil>.": {
			gInput:    nil,
			nameInput: "Update()",
			want:      "Update() = <nil>, want NOT <nil>",
		},
		"When 'got' is a number.": {
			gInput:    1,
			nameInput: "Update()",
		},
	} {
		tc := tc // Rebind 'tc'. Note: This is required to support "parallel" execution.

		// EXECUTION.
		t.Run(tcName, func(t *testing.T) {
			t.Parallel() // Enable parallel execution.

			// ARRANGE.
			testingT := &testableT{TB: t}

			// ACT.
			assert.NotNil(testingT, tc.gInput, tc.nameInput)

			// ASSERT.
			if testingT.failureMsg != tc.want {
				t.Fatalf("Failure message = \"%s\", want \"%s\"", testingT.failureMsg, tc.want)
			}
		})
	}

	// EXECUTION.
	t.Run("When using a custom failure message.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// ARRANGE.
		testingT := &testableT{TB: t}

		// ACT.
		assert.NotNil(testingT, nil, "", "UT Failed: `Update()` - got <nil>, want NOT <nil>.")

		// ASSERT.
		if testingT.failureMsg != "UT Failed: `Update()` - got <nil>, want NOT <nil>." {
			t.Fatalf("Failure message = \"%s\", want \"%s\"", testingT.failureMsg, "UT Failed: `Update()` - got <nil>, want NOT <nil>.")
		}
	})
}

// UT: Compare 2 values for equality.
func TestEqualS(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	for tcName, tc := range map[string]struct {
		gInput, wInput []int
		nameInput      string
		want           string
	}{
		"When 'got' and 'want' are different.": {
			gInput:    []int{1, 2, 3},
			wInput:    []int{3, 2, 1},
			nameInput: "Right([1 2 3], [3 2 1])",
			want:      "Right([1 2 3], [3 2 1]) = [1 2 3], want [3 2 1]",
		},
		"When 'got' and 'want' are equal.": {
			gInput:    []int{1, 2, 3},
			wInput:    []int{1, 2, 3},
			nameInput: "Right([1 2 3], [3 2 1])",
		},
	} {
		tc := tc // Rebind 'tc'. Note: This is required to support "parallel" execution.

		// EXECUTION.
		t.Run(tcName, func(t *testing.T) {
			t.Parallel() // Enable parallel execution.

			// ARRANGE.
			testingT := &testableT{TB: t}

			// ACT.
			assert.EqualS(testingT, tc.gInput, tc.wInput, tc.nameInput)

			// ASSERT.
			if testingT.failureMsg != tc.want {
				t.Fatalf("Failure message = \"%s\", want \"%s\"", testingT.failureMsg, tc.want)
			}
		})
	}

	// EXECUTION.
	t.Run("When using a custom failure message.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// ARRANGE.
		testingT := &testableT{TB: t}
		got, want := []int{1, 2, 3}, []int{3, 2, 1}

		// ACT.
		assert.EqualS(testingT, got, want, "", "UT Failed: `Right([1 2 3], [3 2 1])` - got %v, want %v.", got, want)

		// ASSERT.
		if testingT.failureMsg != "UT Failed: `Right([1 2 3], [3 2 1])` - got [1 2 3], want [3 2 1]." {
			t.Fatalf("Failure message = \"%s\", want \"%s\"", testingT.failureMsg, "UT Failed: `Right([1 2 3], [3 2 1])` - got [1 2 3], want [3 2 1].")
		}
	})
}
