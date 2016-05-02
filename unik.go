package unik

import (
	"strconv"

	"github.com/satori/go.uuid"
	"github.com/tinode/snowflake"
)

var GenFunc func() string
var GenInt64Func func() int64
var sf *snowflake.SnowFlake

func init() {
	Snowflake(1)
}

func Gen() string {
	return GenFunc()
}

func GenInt64() int64 {
	return GenInt64Func()
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

	GenInt64Func = func() int64 {
		idInt, _ := sf.Next()
		return int64(idInt)
	}
}

func Mock(id string) {
	GenFunc = func() string {
		return id
	}
}
