package probego

import "testing"

func TestHello(t *testing.T) {
	if Hello() != "probe-go-ok" {
		t.Fatalf("Hello() = %q", Hello())
	}
}
