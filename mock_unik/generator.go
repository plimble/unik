package mock_unik

import "github.com/stretchr/testify/mock"

type MockGenerator struct {
	mock.Mock
}

func NewMockGenerator() *MockGenerator {
	return &MockGenerator{}
}

func (m *MockGenerator) Generate() string {
	ret := m.Called()

	r0 := ret.Get(0).(string)

	return r0
}
