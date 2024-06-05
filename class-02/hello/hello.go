package hello

import (
	"fmt"
	"strings"
)

func Say(names []string) string {
	if len(names) == 0 {
		names = []string{"world"}
	}

	return fmt.Sprintf("Hello %s", strings.Join(names, ", "))
}
