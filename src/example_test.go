package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExampleSuccess(t *testing.T) {
	result, err := example("hoge")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if result != 1 {
		t.Fatal("failed test")
	}
}

// func TestExampleFailed(t *testing.T) {
//     result, err := example("fuga")
//     if err == nil {
//         t.Fatal("failed test")
//     }
//     if result != 0 {
//         t.Fatal("failed test")
//     }
// }

func TestSum(t *testing.T) {
	actual := sum(3, 5)
	expected := 8

	assert.Equal(t, expected, actual)
}
