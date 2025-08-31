package findparser

import (
	"reflect"
	"testing"
)

func TestFindFindFunction(t *testing.T) {
	src := `contract FindExample {
	uint[] private myCollection;
	function foo(uint x) public pure returns (int) {
		uint result = myCollection.find(x);
		return result;
	}
	function bar() public {
		// nothing
	}
}`
	result := FindFindFunction(src)
	expected := []int{4}
	if !reflect.DeepEqual(result.Lines, expected) {
		t.Errorf("FindFindFunction: got %v, want %v", result.Lines, expected)
	}
}
// func TestFindFindFunction(t *testing.T) {
// 	src := `contract FindExample {
//     uint[] private myCollection;
//     function foo(uint x) public pure returns (int) {
//         uint result = myCollection.find(x);
//         return result;
//     }
//     function bar() public {
//         // nothing
//     }
// }`
// 	result := FindFindFunction(src)
// 	expected := []int{4}
// 	if !reflect.DeepEqual(result.Lines, expected) {
// 		t.Errorf("FindFindFunction: got %v, want %v", result.Lines, expected)
// 	}
// }
