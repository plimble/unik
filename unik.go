package unik

import (
	"strconv"

	"github.com/satori/go.uuid"
	"github.com/tinode/snowflake"
)

var GenFunc func() string
var GenUInt64Func func() uint64
var sf *snowflake.SnowFlake

func init() {
	Snowflake(1)
}

func Gen() string {
	return GenFunc()
}

func GenUInt64() uint64 {
	return GenUInt64Func()
}

func UUIDV1() {
	GenFunc = func() string {
		return uuid.NewV1().String()
	}
}

func UUIDV4() {
	GenFunc = func() string {
		return uuid.NewV4().String()
	}
}

func Snowflake(machine uint32) {
	sf, _ = snowflake.NewSnowFlake(machine)
	GenFunc = func() string {
		idInt, _ := sf.Next()
		return strconv.FormatUint(idInt, 10)
	}

	GenUInt64Func = func() uint64 {
		idInt, _ := sf.Next()
		return idInt
	}
}

func Mock(id string) {
	GenFunc = func() string {
		return id
	}
}
