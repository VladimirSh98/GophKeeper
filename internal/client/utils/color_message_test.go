package utils

import (
	"bytes"
	"testing"

	"github.com/fatih/color"
	"github.com/stretchr/testify/require"
)

func TestColorMessage(t *testing.T) {
	var buf bytes.Buffer
	color.Output = &buf
	ColorMessage("Hello %s", color.FgRed, "World")
	output := buf.String()
	require.Contains(t, output, "Hello World\n")
}
