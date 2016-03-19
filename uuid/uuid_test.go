package uuid

import (
	"testing"
)

func TestV1(t *testing.T) {
	UUIDV1()
	uuid := Gen()
	if len(uuid) != 36 {
		t.Error("Error v1")
	}
}

func TestV4(t *testing.T) {
	UUIDV4()
	uuid := Gen()
	if len(uuid) != 36 {
		t.Error("Error v4")
	}
}
