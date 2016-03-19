package snowflake

import (
	"github.com/tinode/snowflake"
	"strconv"
)

var GenFunc func() string
var sf *snowflake.SnowFlake

func init() {
	Snowflake(1)
}

func Gen() string {
	return GenFunc()
}

func Snowflake(machine uint32) {
	sf, _ = snowflake.NewSnowFlake(machine)
	GenFunc = func() string {
		idInt, _ := sf.Next()
		return strconv.FormatUint(idInt, 10)
	}
}

func Mock(id string) {
	GenFunc = func() string {
		return id
	}
}
