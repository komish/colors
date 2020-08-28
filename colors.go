package colors

type ColorName struct {
	EnglishName string
	SpanishName string
}

func (c *ColorName) InEnglish() string {
	return c.EnglishName
}

func (c *ColorName) InSpanish() string {
	return c.SpanishName
}

type Color interface {
	InEnglish() string
	InSpanish() string
}
