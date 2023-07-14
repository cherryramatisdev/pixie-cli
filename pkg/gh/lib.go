package gh

import (
	"fmt"
	"regexp"
	"strings"

	Z "github.com/rwxrob/bonzai/z"
)

func extractBranchNumber() (string, error) {
	output := Z.Out("git", "symbolic-ref", "--short", "HEAD")
	branchName := strings.TrimSpace(string(output))

	// Define the regular expression pattern to match the branch format
	pattern := `^(feature|fix|chore)/TEC-(\d+)$`

	// Compile the regular expression pattern
	regExp, err := regexp.Compile(pattern)
	if err != nil {
		return "", fmt.Errorf("failed to compile regular expression: %v", err)
	}

	// Find the matches in the branch name
	matches := regExp.FindStringSubmatch(branchName)

	if len(matches) != 3 {
		return "", fmt.Errorf("invalid branch format")
	}

	return matches[2], nil
}
