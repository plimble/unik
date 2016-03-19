package unik

import (
	"testing"
)

func TestUUIDV1(t *testing.T) {
	UUIDV1()
	id := Gen()

	if len(id) != 36 {
		t.Error("Error UUIDV1")
	}
}

func TestUUIDV4(t *testing.T) {
	UUIDV4()
	id := Gen()

	if len(id) != 36 {
		t.Error("Error UUIDV4")
	}
}

func TestSnowflake(t *testing.T) {
	Snowflake(1)
	id := Gen()

	if len(id) != 18 {
		t.Error("Error Snowflake")
	}
}
