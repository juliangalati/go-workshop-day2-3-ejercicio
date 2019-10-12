package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestPrintBSTInOrder(t *testing.T) {
	t.Run("One number", func(t *testing.T) {
		checkBuildAndPrintBST(t, []int{4}, "4")
	})

	t.Run("Duplicated", func(t *testing.T) {
		checkBuildAndPrintBST(t, []int{4, 5, 6, 4, 4}, "4 4 4 5 6")
	})

	t.Run("Negatives", func(t *testing.T) {
		checkBuildAndPrintBST(t, []int{-4, -5, -6}, "-6 -5 -4")
	})

	t.Run("Mixed numbers", func(t *testing.T) {
		checkBuildAndPrintBST(t, []int{4, 99990, 67, 333, -90, 0, 9, 8, 7, 101, -888, 890, 33, 98, 54, 674, 456}, "-888 -90 0 4 7 8 9 33 54 67 98 101 333 456 674 890 99990")
	})
}

func checkBuildAndPrintBST(t *testing.T, numbers []int, want string) {
	t.Helper()
	buffer := bytes.Buffer{}
	BuildAndPrintBST(&buffer, numbers)

	got := strings.TrimSpace(buffer.String())

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
