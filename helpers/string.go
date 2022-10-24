package helpers

import (
	"github.com/iancoleman/strcase"
	"strings"
)

func ToCamelWithoutDot(snakeStr string) string {
	withs := strings.Split(snakeStr, ".")
	for i := range withs {
		withs[i] = strcase.ToCamel(withs[i])
	}
	return strings.Join(withs, ".")
}
