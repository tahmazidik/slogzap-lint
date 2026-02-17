package core

type Kind string

const (
	KindLowercase Kind = "lowercase"
	KindCyrillic  Kind = "cyrillic"
	KindForbidden Kind = "forbidden_symbol"
	KindSensitive Kind = "sensitive"
)

type Violation struct {
	Kind    Kind
	Message string
	BadRune rune
	Keyword string
}
