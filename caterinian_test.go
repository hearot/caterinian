package caterinian

import "testing"

const (
	InvalidNumber = 743
	ValidNumber = 742
)

func TestIsCaterinian(t *testing.T) {
	if !IsCaterinian(ValidNumber) {
		t.Fatalf("IsCaterinian(%d) returned false even if it is a caterinian number", ValidNumber)
	}
	if IsCaterinian(InvalidNumber) {
		t.Fatalf("IsCaterinian(%d) returned true even if it is not a caterinian number", InvalidNumber)
	}
}