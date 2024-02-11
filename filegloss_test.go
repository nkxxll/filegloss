package filegloss

import (
	"testing"
)

func TestPrintInfo(t *testing.T) {
    exp := ""
    f := New("./fixtures/helloworld")
    res := f.InfoString()
    if exp != res {
        t.Fatalf("PrintInfo strings don't match")
    }
}
