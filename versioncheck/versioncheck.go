package versioncheck

import (
	"errors"
	"regexp"
	"strconv"
)

// CheckSolidityVersion parses the contract and returns an error if version < 8
func CheckSolidityVersion(src string) error {
	re := regexp.MustCompile(`pragma solidity \^([0-9]+)\.([0-9]+)\.([0-9]+);`)
	matches := re.FindStringSubmatch(src)
	if len(matches) < 4 {
		return errors.New("pragma solidity version not found")
	}
	major, _ := strconv.Atoi(matches[1])
	if major < 8 {
		return errors.New("Solidity version must be 8 or higher")
	}
	return nil
}
