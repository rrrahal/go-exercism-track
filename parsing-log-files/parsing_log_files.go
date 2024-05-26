package parsinglogfiles

import (
	"regexp"
)

func IsValidLine(text string) bool {
	start_regex := regexp.MustCompile(`\b(?:TRC|DBG|INF|WRN|ERR|FTL)\w*\b`)

	return start_regex.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<*>`)

	return re.Split(text, 3)
}

func CountQuotedPasswords(lines []string) int {

	panic("Please implement the CountQuotedPasswords function")
}

func RemoveEndOfLineText(text string) string {
	panic("Please implement the RemoveEndOfLineText function")
}

func TagWithUserName(lines []string) []string {
	panic("Please implement the TagWithUserName function")
}
