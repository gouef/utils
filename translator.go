package utils

type Translator interface {
	translate(message string, parameters ...interface{})
}
