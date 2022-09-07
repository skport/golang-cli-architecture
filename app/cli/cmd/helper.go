// Helper for testing

package cmd

import (
	"bytes"
	"os"
	"testing"
)

// Extract standard output and return as string type
func extractStdout(t *testing.T, fnc func()) string {
	t.Helper()

	saved := os.Stdout
	defer func() {
		os.Stdout = saved
	}()

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("fail pipe: %v", err)
	}
	os.Stdout = w

	fnc()
	w.Close()

	var buffer bytes.Buffer
	if n, err := buffer.ReadFrom(r); err != nil {
		t.Fatalf("fail read buf: %v - number: %v", err, n)
	}
	s := buffer.String()

	return s[:len(s)-1]
}
