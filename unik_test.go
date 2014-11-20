package unik_test

import (
	. "github.com/plimble/unik"
	"testing"
)

func TestGenerateID(t *testing.T) {
	idgen := New()
	idgen.Snowflake()
	idgen.SnowflakeID(120)
	idgen.BSON()
	idgen.UUIDV1()
	idgen.UUIDV4()
}
