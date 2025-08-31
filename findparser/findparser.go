package findparser

import (
	"regexp"
)

type ContractInfo struct {
	Lines    []int
}

func FindFindFunction(src string) ContractInfo {
	lines := regexp.MustCompile(`\r?\n`).Split(src, -1)
	var found []int
	for i, line := range lines {
		if regexp.MustCompile(`\.find\b`).FindString(line) != "" {
			found = append(found, i+1) // 1-based line numbers
		}
	}
	return ContractInfo{
		Lines: found,
	}
}
