package hamcrest

import (
	"strings"
)

// The only reason why this exists is to make hamcrest testable.
// Feel free to expose more funcs of the testing helper here...
type TestingT interface {
	Helper()
	Errorf(format string, args ...any)
}

func Expect(got any, t TestingT) *ExpectResult {
	t.Helper()
	return &ExpectResult{got, t}
}

type ExpectResult struct {
	Got any
	t   TestingT
}

func (e *ExpectResult) ToEqual(want any) {
	e.t.Helper()

	handled := false

	{
		gotString, gotIsString := e.Got.(string)
		if gotIsString {
			handled = true
			if gotString != want {
				e.t.Errorf("got %q, want %q", gotString, want)
			}
		}
	}

	{
		gotInt, isInt := e.Got.(int)
		if isInt {
			handled = true
			if gotInt != want {
				e.t.Errorf("got %d, want %d", gotInt, want)
			}
		}
	}

	{
		gotBool, gotIsBool := e.Got.(bool)
		if gotIsBool {
			handled = true
			if gotBool != want {
				e.t.Errorf("got %t, want %t", gotBool, want)
			}
		}
	}

	if !handled {
		e.t.Errorf("matcher for ToEqual was not implemented. Consider implementing it.")
	}

}

func (e *ExpectResult) ToContain(want any) {
	e.t.Helper()

	handled := false

	{
		gotString, gotIsString := e.Got.(string)
		wantString, wantIsString := want.(string)
		if gotIsString && wantIsString {
			handled = true
			if !strings.Contains(gotString, wantString) {
				e.t.Errorf("got %q did not contain %q", gotString, wantString)
			}
		}
	}

	if !handled {
		e.t.Errorf("matcher for ToContain was not implemented. Consider implementing it.")
	}

}

func (e *ExpectResult) ToStartWith(want any) {
	handled := false

	{
		gotString, gotIsString := e.Got.(string)
		wantString, wantIsString := want.(string)

		if gotIsString && wantIsString {
			handled = true
			if !strings.HasPrefix(gotString, wantString) {
				e.t.Errorf("got %q, want %q", gotString, want)
			}
		}
	}

	if !handled {
		e.t.Errorf("matcher for ToEqual was not implemented. Consider implementing it.")
	}
}
