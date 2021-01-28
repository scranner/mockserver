package templating

import (
	"github.com/brianvoe/gofakeit/v5"
	guuid "github.com/google/uuid"
	"strings"
)


func IsTemplateValue(value string) bool {
	return strings.HasPrefix(value, "[") && strings.HasSuffix(value,"]")
}

func ProcessTemplateVariable(value string) func() string {
	switch value {
	case "[guid]": return func() string {
		return guuid.New().String()
	}
	case "[name]": return func() string {
		return gofakeit.Name()
	}
	case "[lorum-paragraph]": return func() string {
		return gofakeit.LoremIpsumParagraph(2, 3, 30, " ")
	}
	case "[image-url]": return func() string {
		return gofakeit.ImageURL(200, 200)
	}
	default: return func() string {
		return value
	}
	}
}
