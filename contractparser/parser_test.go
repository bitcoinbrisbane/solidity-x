package contractparser

import (
	"reflect"
	"testing"
)

func TestParseContractLayout(t *testing.T) {
	src := `contract Greeter {
    string private greeting; // State variable to store the greeting message

    // Constructor: executed only once when the contract is deployed
    constructor(string memory _initialGreeting) {
        greeting = _initialGreeting;
    }

    // Function to retrieve the current greeting
    function greet() public view returns (string memory) {
        return greeting;
    }

    // Function to update the greeting message
    function setGreeting(string memory _newGreeting) public {
        greeting = _newGreeting;
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
