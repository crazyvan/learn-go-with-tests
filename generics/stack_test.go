package stack_test

import (
	"testing"

	generics "github.com/crazyvan/learn-go-with-tests/generics"
)

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		stackOfInts := new(generics.Stack[int])

		AssertTrue(t, stackOfInts.IsEmpty())

		stackOfInts.Push(123)
		AssertFalse(t, stackOfInts.IsEmpty())

		value, _ := stackOfInts.Pop()
		AssertEqual(t, value, 123)
		AssertTrue(t, stackOfInts.IsEmpty())
	})
}
