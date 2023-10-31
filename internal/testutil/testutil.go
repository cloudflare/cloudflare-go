package testutil

import (
	"net/http"
	"os"
	"strconv"
	"testing"
)

func CheckTestServer(t *testing.T, url string) bool {
	if _, err := http.Get(url); err != nil {
		str := os.Getenv("SKIP_MOCK_TESTS")
		skip, err := strconv.ParseBool(str)
		if err != nil {
			panic("invalid value for flag SKIP_MOCK_TESTS \"" + str + "\", should be parsable as a bool")
		}
		if skip {
			t.Skip("The test will not run without a mock Prism server running against your OpenAPI spec")
			return false
		}
		t.Error("The test will not run without a mock Prism server running against your OpenAPI spec. You can set the environment variable SKIP_MOCK_TESTS to true to skip running any tests that require the mock server.")
		return false
	}
	return true
}
