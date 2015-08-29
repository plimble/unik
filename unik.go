package unik

//go:generate mockery -name Generator -output mock_unik

type Generator interface {
	Generate() string
}
