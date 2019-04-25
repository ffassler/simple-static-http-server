package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func managePanic(t *testing.T) {
	if r := recover(); r != nil {
		t.Failed()
		t.Errorf("panic during execution: %v", r)
	}
}

func TestGetServeFilename(t *testing.T) {
	defer managePanic(t)

	directory := "/dir"
	path := "/path"

	filename := GetServeFilename(directory, path)

	assert.Equal(t, "/dir/path", filename)
}

func TestGetServeFilenameNoPath(t *testing.T) {
	defer managePanic(t)

	directory := "/dir"
	path := ""

	filename := GetServeFilename(directory, path)

	assert.Equal(t, "/dir/", filename)
}
