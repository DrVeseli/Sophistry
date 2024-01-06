// customlinks.go
package main

import (
	"fmt"
	"regexp"
)

func replaceCustomLinks(content []byte) []byte {
	re := regexp.MustCompile(`\[\[([^\]]+)\]\]`)
	return re.ReplaceAllFunc(content, func(match []byte) []byte {
		linkText := string(match[2 : len(match)-2]) // Extracting the content inside [[ ]]
		link := fmt.Sprintf("[%s.html](%s.html)", linkText, linkText)
		return []byte(link)
	})
}
