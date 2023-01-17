package hamcrest

import (
	"testing"
)

type TestingTForTest struct {
	errorfCalled int
}

func (t *TestingTForTest) Helper() {
	// not used
}
func (t *TestingTForTest) Errorf(format string, args ...any) {
	t.errorfCalled++
}

func TestHamcrest(t *testing.T) {
	t.Run("ToContains works for positive cases", func(t *testing.T) {
		// positive tests
		Expect("test", t).ToContain("es")
		Expect("test", t).ToContain("test")
		Expect("test", t).ToContain("es")

		// Negative tests. Can't use t here... mocking t otherwise the test would fail
		// as we expect failure
		testingTForTest := TestingTForTest{0}
		Expect("test", &testingTForTest).ToContain("bingo")
		if testingTForTest.errorfCalled != 1 {
			t.Errorf("test seems to contain bingo - arg. That can't be the case...")
		}
	})

	t.Run("ToEquals works for string", func(t *testing.T) {
		// positive tests
		Expect("test", t).ToEqual("test")
		Expect(true, t).ToEqual(true)
		Expect(false, t).ToEqual(false)

		// Negative tests. Can't use t here... mocking t otherwise the test would fail
		// as we expect failure
		testingTForTest := TestingTForTest{0}
		Expect("test", &testingTForTest).ToEqual("dingo")
		if testingTForTest.errorfCalled != 1 {
			t.Errorf("test seems to equal dingo - arg. That can't be the case...")
		}
	})

	t.Run("ToStartWith works", func(t *testing.T) {
		// positive tests
		Expect("test", t).ToStartWith("t")
		Expect("test", t).ToStartWith("te")
		Expect("test", t).ToStartWith("test")

		// Negative tests. Can't use t here... mocking t otherwise the test would fail
		// as we expect failure
		testingTForTest := TestingTForTest{0}
		Expect("bingo", &testingTForTest).ToStartWith("dingo")
		if testingTForTest.errorfCalled != 1 {
			t.Errorf("bingo seems to start with dingo - arg. That can't be the case...")
		}
	})

}
