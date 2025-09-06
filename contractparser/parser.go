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
		// Simple global var detection (not in function, ends with ';')
		if !inFunc && len(trim) > 0 && trim[len(trim)-1] == ';' && (containsType(trim)) {
			layout.GlobalVars = append(layout.GlobalVars, GlobalVarInfo{
				Name:      extractVarName(trim),
				Type:      extractVarType(trim),
				StartLine: i,
				EndLine:   i,
			})
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

// Dummy helpers for type/var extraction (improve as needed)
func containsType(line string) bool {
	return (len(line) > 3 && (line[:3] == "int" || line[:4] == "uint"))
}

func extractVarName(line string) string {
	// crude: last word before ;
	parts := []rune(line)
	end := len(parts) - 1
	for end > 0 && (parts[end] == ';' || parts[end] == ' ') {
		end--
	}
	start := end
	for start > 0 && parts[start] != ' ' {
		start--
	}
	return string(parts[start+1 : end+1])
}

func extractVarType(line string) string {
	// crude: first word
	for i := 0; i < len(line); i++ {
		if line[i] == ' ' {
			return line[:i]
		}
	}
	return line
}
