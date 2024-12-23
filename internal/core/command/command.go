/*
 * FOCELA TECHNOLOGIES INTERNAL USE ONLY LICENSE AGREEMENT
 *
 * Copyright (c) 2024 Focela Technologies. All rights reserved.
 *
 * Permission is hereby granted to employees or authorized personnel of Focela
 * Technologies (the "Company") to use this software solely for internal business
 * purposes within the Company.
 *
 * For inquiries or permissions, please contact: legal@focela.com
 */

// Package command provides tools for parsing and managing command-line arguments, options, and environment variables.
package command

import (
	"os"
	"regexp"
	"strings"
)

var (
	defaultParsedArgs    = make([]string, 0)
	defaultParsedOptions = make(map[string]string)
	argumentOptionRegex  = regexp.MustCompile(`^\-{1,2}([\w\?\.\-]+)(=)?(.*)$`)
)

// Init initializes command arguments and options.
// If no arguments are provided, it defaults to os.Args.
func Init(args ...string) {
	if len(args) == 0 {
		if len(defaultParsedArgs) == 0 && len(defaultParsedOptions) == 0 {
			args = os.Args
		} else {
			return
		}
	} else {
		defaultParsedArgs = []string{}
		defaultParsedOptions = map[string]string{}
	}
	defaultParsedArgs, defaultParsedOptions = ParseUsingDefaultAlgorithm(args...)
}

// ParseUsingDefaultAlgorithm parses command-line arguments and options.
// Supports both single-dash and double-dash options.
func ParseUsingDefaultAlgorithm(args ...string) (parsedArgs []string, parsedOptions map[string]string) {
	parsedArgs = []string{}
	parsedOptions = map[string]string{}

	for i := 0; i < len(args); i++ {
		array := argumentOptionRegex.FindStringSubmatch(args[i])
		if len(array) > 2 {
			if array[2] == "=" {
				parsedOptions[array[1]] = array[3]
			} else if i < len(args)-1 && len(args[i+1]) > 0 && args[i+1][0] != '-' {
				parsedOptions[array[1]] = args[i+1]
				i++
			} else {
				parsedOptions[array[1]] = array[3]
			}
		} else {
			parsedArgs = append(parsedArgs, args[i])
		}
	}
	return
}

// GetOpt retrieves the value of a command-line option by its name.
// If the option does not exist, it returns the default value if provided.
func GetOpt(name string, def ...string) string {
	Init()
	if v, ok := defaultParsedOptions[name]; ok {
		return v
	}
	if len(def) > 0 {
		return def[0]
	}
	return ""
}

// GetOptAll retrieves all parsed options as a map.
func GetOptAll() map[string]string {
	Init()
	return defaultParsedOptions
}

// ContainsOpt checks whether an option exists in the parsed arguments.
func ContainsOpt(name string) bool {
	Init()
	_, ok := defaultParsedOptions[name]
	return ok
}

// GetArg retrieves an argument by its index.
// If the index is out of range, it returns the default value if provided.
func GetArg(index int, def ...string) string {
	Init()
	if index < len(defaultParsedArgs) {
		return defaultParsedArgs[index]
	}
	if len(def) > 0 {
		return def[0]
	}
	return ""
}

// GetArgAll retrieves all parsed arguments as a slice.
func GetArgAll() []string {
	Init()
	return defaultParsedArgs
}

// GetOptWithEnv retrieves the value of a command-line option or environment variable.
// Command-line keys are lowercase (e.g., altura.package.variable).
// Environment variable keys are uppercase (e.g., ALTURA_PACKAGE_VARIABLE).
func GetOptWithEnv(key string, def ...string) string {
	cmdKey := strings.ToLower(strings.ReplaceAll(key, "_", "."))
	if ContainsOpt(cmdKey) {
		return GetOpt(cmdKey)
	}

	envKey := strings.ToUpper(strings.ReplaceAll(key, ".", "_"))
	if r, ok := os.LookupEnv(envKey); ok {
		return r
	}

	if len(def) > 0 {
		return def[0]
	}
	return ""
}
