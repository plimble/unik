package unik

//go:generate mockery -name Generator

type Generator interface {
	Generate() string
}
