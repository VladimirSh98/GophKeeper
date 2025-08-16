package utils

import "github.com/fatih/color"

func MessageWithoutArgs(message string, intColor color.Attribute) {
	msg := color.New(intColor).PrintlnFunc()
	msg(message)
}

func MessageWithArgs(message string, intColor color.Attribute, args ...interface{}) {
	msg := color.New(intColor).PrintfFunc()
	msg(message+"\n", args...)
}
