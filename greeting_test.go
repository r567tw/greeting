package greeting

import (
    "testing"
)

// TestHelloName calls greeting.Hello with a name, checking
// for a valid return value.
func TestHello(t *testing.T) {
    want := "Hello"
    msg := Hello()
    if msg != want{
        t.Fatalf("Error")
    }
}

// TestHi calls greeting.Hi with a name, checking
func TestHi(t *testing.T) {
    want := "Hi"
    msg := Hi()
    if msg != want{
        t.Fatalf("Error")
    }
}
