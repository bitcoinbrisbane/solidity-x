package contractparser

import (
	"regexp"
	"strings"
)

type Statements struct {
	Lines []string
}

type FunctionInfo struct {
	Name       string
	StartLine  int
	EndLine    int
	Visibility string
	// Args       []string
}

type GlobalVarInfo struct {
	Name      string
	Type      string
	StartLine int
	EndLine   int
}

type ContractLayout struct {
	Functions  []FunctionInfo
	GlobalVars []GlobalVarInfo
}

const FUNCTION_REGEX = `function\s+([a-zA-Z_][a-zA-Z0-9_]*)\s*\([^)]*\)\s*(?:(public|private|internal|external)\s+)?(?:(?:pure|view|payable|nonpayable)\s+)?(?:returns\s*\([^)]*\)\s*)?(?:override\s+)?(?:virtual\s+)?\{`
const GLOBAL_VAR_REGEX = `^\s*(?:(?:uint(?:\d+)?|int(?:\d+)?|bool|address|string|bytes(?:\d+)?|mapping\s*\([^)]+\)|[a-zA-Z_][a-zA-Z0-9_]*(?:\[\])*)\s+)(?:(public|private|internal)\s+)?(?:(constant|immutable)\s+)?([a-zA-Z_][a-zA-Z0-9_]*)\s*(?:=\s*[^;]+)?\s*;`

// ParseContractLayout scans each line and records start/end lines of functions and global vars
func ParseContractLayout(lines []string) ContractLayout {
	var layout ContractLayout
	inFunc := false
	funcStart := 0
	funcName := ""
	visibility := ""
	braceCount := 0
	const functionLength = 8

	for i, line := range lines {
		trim := strings.TrimLeft(line, " \t")
		if !inFunc && len(trim) > functionLength && trim[:functionLength] == "function" {
			inFunc = true
			funcStart = i
			braceCount = 0
			// Extract function name and visibility using regex
			funcNameRegex := regexp.MustCompile(FUNCTION_REGEX)
			matches := funcNameRegex.FindStringSubmatch(trim)
			if len(matches) > 1 {
				funcName = matches[1]
			}
			visibility = "internal" // default visibility
			if len(matches) > 2 && matches[2] != "" {
				visibility = matches[2]
			}
		}
		if inFunc {
			braceCount += countChar(line, '{')
			braceCount -= countChar(line, '}')
			if braceCount == 0 {
				layout.Functions = append(layout.Functions, FunctionInfo{
					Name:       funcName,
					StartLine:  funcStart,
					EndLine:    i,
					Visibility: visibility,
				})
				inFunc = false
			}
		}
		// Global var detection using regex
		if !inFunc {
			globalVarRegex := regexp.MustCompile(GLOBAL_VAR_REGEX)
			matches := globalVarRegex.FindStringSubmatch(line)
			if len(matches) > 3 {
				varType := strings.TrimSpace(matches[0][:strings.LastIndex(matches[0], matches[3])])
				varType = strings.Fields(varType)[0] // Get first word as type
				layout.GlobalVars = append(layout.GlobalVars, GlobalVarInfo{
					Name:      matches[3],
					Type:      varType,
					StartLine: i,
					EndLine:   i,
				})
			}
		}
	}
	return layout
}

func countChar(s string, c byte) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			count++
		}
	}
	return count
}
