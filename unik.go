package unik

import (
	"github.com/satori/go.uuid"
	"github.com/sdming/gosnow"
	mgo "gopkg.in/mgo.v2/bson"
	"strconv"
)

//go:generate mockgen -destination=mock.go -package=unik github.com/plimble/unik Generator

type Generator interface {
	Generate() string
}

type snowflake struct {
	gosnow *gosnow.SnowFlake
}

func NewSnowflake(index int) Generator {
	gosnow, err := gosnow.NewSnowFlake(uint32(index))
	if err != nil {
		panic(err)
	}
	return &snowflake{gosnow}
}

func (g *snowflake) Generate() string {
	result, _ := g.gosnow.Next()
	return strconv.FormatInt(int64(result), 10)
}

type uuidv1 struct{}

func NewUUIDV1() Generator {
	return &uuidv1{}
}

func (g *uuidv1) Generate() string {
	return uuid.NewV1().String()
}

type uuidv4 struct{}

func NewUUIDV4() Generator {
	return &uuidv4{}
}

func (g *uuidv4) Generate() string {
	return uuid.NewV4().String()
}

type bson struct{}

func NewBSON() Generator {
	return &bson{}
}

func (g *bson) Generate() string {
	return mgo.NewObjectId().Hex()
}
