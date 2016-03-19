package uuid

import (
	"github.com/satori/go.uuid"
)

var GenFunc func() string

func init() {
	UUIDV1()
}

func Gen() string {
	return GenFunc()
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

func Mock(id string) {
	GenFunc = func() string {
		return id
	}
}
