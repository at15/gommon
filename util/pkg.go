// Package util contains helper for test
package util

func GeneratedHeader(generator string, template string) string {
	return "// Code generated by " + generator + " from " + template + " DO NOT EDIT.\n"
}
