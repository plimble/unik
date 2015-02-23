package mocks

import "github.com/stretchr/testify/mock"

type Generator struct {
	mock.Mock
}

func NewGenerator() *Generator {
	return &Generator{}
}

func (m *Generator) Generate() string {
	ret := m.Called()

	r0 := ret.Get(0).(string)

	return r0
}
