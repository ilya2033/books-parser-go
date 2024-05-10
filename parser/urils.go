package parser

import (
	"regexp"
	"strings"
)

func RemoveScripts(html string) string {
	scriptRe := regexp.MustCompile(`(?s)<script.*<\/script>`)
	endScriptRe := regexp.MustCompile(`</script>`)

	result := scriptRe.ReplaceAllStringFunc(html, func(s string) string {
		index := endScriptRe.FindStringIndex(s)
		return strings.TrimSpace(s[index[1]:])
	})
	return result
}
