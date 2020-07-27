package main

import (
	"net"
	"strconv"
	"strings"
)

// Analyze summarizes a stream of input without modifying user files.
func Analyze(input string) map[string]interface{} {

	valid := []string{}
	invalid := []string{}
	for _, line := range strings.Split(input, "\n") {
		item := strings.TrimSpace(line)
		if item == "" {
			continue
		}
		host, portText, err := net.SplitHostPort(item)
		if err != nil {
			invalid = append(invalid, item)
			continue
		}
		port, err := strconv.Atoi(portText)
		if err != nil || port < 1 || port > 65535 || host == "" {
			invalid = append(invalid, item)
			continue
		}
		valid = append(valid, item)
	}
	return map[string]interface{}{"valid": valid, "invalid": invalid, "checked": len(valid) + len(invalid)}
}
