// Copyright (c) 2024 Focela Technologies. All rights reserved.
// This source code is governed by an MIT License.
// See LICENSE file for full terms and conditions.

// Package cmd parses command-line arguments and options.
package cmd

import (
	"os"
	"regexp"
	"strings"
)

var (
	// Caches parsed arguments and options.
	parsedArgs    = make([]string, 0)
	parsedOptions = make(map[string]string)
	// Matches command-line options with an optional value.
	optionRegex = regexp.MustCompile(`^\-{1,2}([\w\?\.\-]+)(=){0,1}(.*)$`)
)

// Init parses arguments and options. If args are empty, it defaults to os.Args.
// It avoids re-parsing if values are already set.
func Init(args ...string) {
	if len(args) == 0 {
		if len(parsedArgs) == 0 && len(parsedOptions) == 0 {
			args = os.Args
		} else {
			return
		}
	} else {
		parsedArgs = make([]string, 0)
		parsedOptions = make(map[string]string)
	}
	parsedArgs, parsedOptions = parseDefault(args...)
}

// parseDefault processes args into standalone arguments and options.
func parseDefault(args ...string) (argsOut []string, optsOut map[string]string) {
	argsOut = make([]string, 0)
	optsOut = make(map[string]string)

	for i := 0; i < len(args); {
		matches := optionRegex.FindStringSubmatch(args[i])
		if len(matches) > 2 {
			if matches[2] == "=" {
				optsOut[matches[1]] = matches[3]
			} else if i < len(args)-1 {
				if len(args[i+1]) > 0 && args[i+1][0] == '-' {
					optsOut[matches[1]] = matches[3]
				} else {
					optsOut[matches[1]] = args[i+1]
					i += 2
					continue
				}
			} else {
				optsOut[matches[1]] = matches[3]
			}
		} else {
			argsOut = append(argsOut, args[i])
		}
		i++
	}
	return
}

// GetOpt returns the value for the given option name, or the default if not found.
func GetOpt(name string, def ...string) string {
	Init()
	if v, ok := parsedOptions[name]; ok {
		return v
	}
	if len(def) > 0 {
		return def[0]
	}
	return ""
}

// GetOptAll returns all parsed options.
func GetOptAll() map[string]string {
	Init()
	return parsedOptions
}

// ContainsOpt checks if the given option exists.
func ContainsOpt(name string) bool {
	Init()
	_, ok := parsedOptions[name]
	return ok
}

// GetArg returns the argument at the specified index, or the default if out of range.
func GetArg(index int, def ...string) string {
	Init()
	if index < len(parsedArgs) {
		return parsedArgs[index]
	}
	if len(def) > 0 {
		return def[0]
	}
	return ""
}

// GetArgAll returns all parsed arguments.
func GetArgAll() []string {
	Init()
	return parsedArgs
}

// GetOptWithEnv returns the command-line option for key, or the corresponding environment variable,
// or the default if neither is found.
func GetOptWithEnv(key string, def ...string) string {
	cmdKey := strings.ToLower(strings.ReplaceAll(key, "_", "."))
	if ContainsOpt(cmdKey) {
		return GetOpt(cmdKey)
	}
	envKey := strings.ToUpper(strings.ReplaceAll(key, ".", "_"))
	if value, ok := os.LookupEnv(envKey); ok {
		return value
	}
	if len(def) > 0 {
		return def[0]
	}
	return ""
}
