package unik

import (
	"encoding/base64"
	"encoding/hex"
	"github.com/plimble/utils/strings2"
	"github.com/satori/go.uuid"
	"github.com/sdming/gosnow"
	mgo "gopkg.in/mgo.v2/bson"
	"strconv"
	"strings"
)

//go:generate mockery -name Generator

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
	v1 := uuid.NewV1()
	data := v1.Bytes()
	buf := make([]byte, 32)
	hex.Encode(buf, data)
	return string2.BytesToString(buf)
}

type uuidv4 struct{}

func NewUUIDV4() Generator {
	return &uuidv4{}
}

func (g *uuidv4) Generate() string {
	v4 := uuid.NewV4()
	data := v4.Bytes()
	buf := make([]byte, 32)
	hex.Encode(buf, data)
	return string2.BytesToString(buf)
}

type uuid1base64 struct{}

func NewUUID1Base64() Generator {
	return &uuid1base64{}
}

func (g *uuid1base64) Generate() string {
	id := base64.URLEncoding.EncodeToString(uuid.NewV1().Bytes())
	id = strings.TrimRight(id, "=")
	return id
}

type bson struct{}

func NewBSON() Generator {
	return &bson{}
}

func (g *bson) Generate() string {
	return mgo.NewObjectId().Hex()
}
