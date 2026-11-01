package croc

import (
	"errors"
	"testing"
)

type timeoutNetErr struct{ msg string }

func (e timeoutNetErr) Error() string   { return e.msg }
func (e timeoutNetErr) Timeout() bool   { return true }
func (e timeoutNetErr) Temporary() bool { return true }

func TestIsLocalProbeTimeout(t *testing.T) {
	if !isLocalProbeTimeout(timeoutNetErr{msg: "read tcp 198.18.0.1:1->47.109.201.184:9009: i/o timeout"}) {
		t.Fatal("net timeout should fall back to the relay")
	}
	if !isLocalProbeTimeout(errors.New("read tcp: i/o timeout")) {
		t.Fatal("i/o timeout string should fall back")
	}
	if !isLocalProbeTimeout(errors.New("context deadline exceeded")) {
		t.Fatal("deadline exceeded should fall back")
	}
	if isLocalProbeTimeout(errors.New("connection reset by peer")) {
		t.Fatal("reset must stay fatal")
	}
	if isLocalProbeTimeout(nil) {
		t.Fatal("nil is not a timeout")
	}
}
