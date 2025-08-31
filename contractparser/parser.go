package contractparser

import "strings"

type FunctionInfo struct {
	Name      string
	StartLine int
	EndLine   int
}

type GlobalVarInfo struct {
	Name      string
	Type      string
	StartLine int
	EndLine   int
}

type ContractLayout struct {
	Functions   []FunctionInfo
	GlobalVars  []GlobalVarInfo
}

// ParseContractLayout scans each line and records start/end lines of functions and global vars
func ParseContractLayout(lines []string) ContractLayout {
	var layout ContractLayout
	inFunc := false
	funcStart := 0
	funcName := ""
	braceCount := 0

	for i, line := range lines {
		trim := strings.TrimLeft(line, " \t")
		if !inFunc && len(trim) > 8 && trim[:8] == "function" {
			inFunc = true
			funcStart = i
			braceCount = 0
			// Extract function name
			name := ""
			for j := 8; j < len(trim); j++ {
				if trim[j] == ' ' || trim[j] == '(' {
					name = trim[8:j]
					break
				}
			}
			funcName = name
		}
		if inFunc {
			braceCount += countChar(line, '{')
			braceCount -= countChar(line, '}')
			if braceCount == 0 {
				layout.Functions = append(layout.Functions, FunctionInfo{
					Name:      funcName,
					StartLine: funcStart,
					EndLine:   i,
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
