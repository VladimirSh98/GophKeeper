package utils

import "github.com/fatih/color"

// ColorMessage print message with args
func ColorMessage(message string, intColor color.Attribute, args ...interface{}) {
	msg := color.New(intColor).PrintfFunc()
	msg(message+"\n", args...)
}
