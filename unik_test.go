package unik

import (
	"testing"
)

func TestSnowflake(t *testing.T) {
	idgen := New()
	idgen.Snowflake()
	idgen.SnowflakeID(120)
	idgen.BSON()
	idgen.UUIDV1()
	idgen.UUIDV4()
}
