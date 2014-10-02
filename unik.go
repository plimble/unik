package unik

import (
	"github.com/satori/go.uuid"
	"github.com/sdming/gosnow"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

type Generator interface {
	Snowflake() string
	SnowflakeID(id int) string
	BSON() string
	UUIDV1() string
	UUIDV4() string
}

type unik struct {
	snowflakes []*gosnow.SnowFlake
}

func New() Generator {
	snowflakes := make([]*gosnow.SnowFlake, 200)
	snowflakes[0], _ = gosnow.NewSnowFlake(uint32(0))
	return &unik{
		snowflakes: snowflakes,
	}
}

func (u *unik) Snowflake() string {
	result, _ := u.snowflakes[0].Next()
	return strconv.FormatInt(int64(result), 10)
}

func (u *unik) SnowflakeID(id int) string {
	if id > 200 {
		id = 0
	}

	if u.snowflakes[id] == nil {
		u.snowflakes[id], _ = gosnow.NewSnowFlake(uint32(id))
	}

	result, _ := u.snowflakes[id].Next()

	return strconv.FormatInt(int64(result), 10)
}

func (u *unik) BSON() string {
	return bson.NewObjectId().Hex()
}

func (u *unik) UUIDV1() string {
	return uuid.NewV1().String()
}

func (u *unik) UUIDV4() string {
	return uuid.NewV4().String()
}
