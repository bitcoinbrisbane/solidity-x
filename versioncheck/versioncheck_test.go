package versioncheck

import (
	"testing"
)

func TestCheckSolidityVersion(t *testing.T) {
	cases := []struct {
		src     string
		hasErr  bool
	}{
		{"pragma solidity ^0.7.0;", true},
		{"pragma solidity ^8.0.0;", false},
		{"pragma solidity ^9.1.2;", false},
		{"// SPDX-License-Identifier: MIT\npragma solidity ^7.9.0;", true},
		{"// SPDX-License-Identifier: MIT\npragma solidity ^8.1.0;", false},
		{"// SPDX-License-Identifier: MIT\n", true},
	}
	for _, c := range cases {
		err := CheckSolidityVersion(c.src)
		if (err != nil) != c.hasErr {
			t.Errorf("CheckSolidityVersion(%q) error = %v, wantErr %v", c.src, err, c.hasErr)
		}
	}
}
