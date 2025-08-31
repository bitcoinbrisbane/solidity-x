package contractparser

import (
	"reflect"
	"testing"
)

func TestParseContractLayout(t *testing.T) {
	src := `contract FindExample {
    uint[] private myCollection;
    function foo(uint x) public pure returns (int) {
        return 1;
    }
    function bar() public {
        // do nothing
    }
}`
	lines := make([]string, 0)
	for _, line := range splitLines(src) {
		lines = append(lines, line)
	}
	layout := ParseContractLayout(lines)

	expectedFuncs := []FunctionInfo{
		{"foo", 2, 4},
		{"bar", 5, 7},
	}
	expectedVars := []GlobalVarInfo{
		{"myCollection", "uint[]", 1, 1},
	}

	if !reflect.DeepEqual(layout.Functions, expectedFuncs) {
		t.Errorf("Functions: got %+v, want %+v", layout.Functions, expectedFuncs)
	}
	if !reflect.DeepEqual(layout.GlobalVars, expectedVars) {
		t.Errorf("GlobalVars: got %+v, want %+v", layout.GlobalVars, expectedVars)
	}
}

func splitLines(s string) []string {
	var lines []string
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			lines = append(lines, s[start:i])
			start = i + 1
		}
	}
	if start < len(s) {
		lines = append(lines, s[start:])
	}
	return lines
}
