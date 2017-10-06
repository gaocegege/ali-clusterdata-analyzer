package parser

type Parser interface {
	GetFileName() string
	ParseFromFile() string
}
