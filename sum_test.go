package demo_test

import (
	"testing"

	demo "github.com/pochanglin/github-actions-demo"
)

func TestSum(t *testing.T) {
	result := demo.Sum(1, 2)

	if result != 3 {
		t.Error("testing sum failed!")
	}
}
