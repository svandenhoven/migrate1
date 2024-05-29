package tasks

import (
	"regexp"
	"strings"
)

// Detect a markdown codeblock
var codeBlockRegex = regexp.MustCompile("```[\\s\\S]*?```")

// Remove code blocks from a string
// In some cases, we want to skip code blocks when bulk-updating content
func ExtractCodeBlocks(content string) (string, []string) {
	codeBlocks := codeBlockRegex.FindAllString(content, -1)
	processedContent := codeBlockRegex.ReplaceAllString(content, "{CODE_BLOCK}")
	return processedContent, codeBlocks
}

// Re-add code blocks to a string
func ReinsertCodeBlocks(content string, codeBlocks []string) string {
	for _, codeBlock := range codeBlocks {
		content = strings.Replace(content, "{CODE_BLOCK}", codeBlock, 1)
	}
	return content
}
