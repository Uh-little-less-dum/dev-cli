package pathutils

import (
	"fmt"
	"path/filepath"
	"strings"
)

// Applies the ./xyz for xyz if it does not start with a period, otheriwse returns the original string.
func ApplyLeadingSlash(s string) string {
	if strings.HasPrefix(s, ".") {
		return s
	}
	return fmt.Sprintf("./%s", s)
}

func RemoveFileExtension(s string) string {
	e := filepath.Ext(s)
	if e != "" {
		rest, val := strings.CutSuffix(s, e)
		if val {
			return rest
		}
	}
	return s
}

func AddLeadingSlashAndRemoveExtension(s string) string {
	return RemoveFileExtension(ApplyLeadingSlash(s))
}
