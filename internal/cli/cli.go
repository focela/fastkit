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

// Package cli provides console operations, like options/arguments reading.
package cli

import (
	"os"
	"regexp"
	"strings"
)

// Global variables
var (
	defaultParsedArgs    []string
	defaultParsedOptions map[string]string
	argumentOptionRegex  = regexp.MustCompile(`^\-{1,2}([\w\?\.\-]+)(=){0,1}(.*)$`)
)

// Init initializes CLI argument parsing.
// It resets parsed arguments and options if they already exist.
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

// ParseUsingDefaultAlgorithm parses command-line arguments into options and arguments.
func ParseUsingDefaultAlgorithm(args ...string) (parsedArgs []string, parsedOptions map[string]string) {
	parsedArgs = []string{}
	parsedOptions = map[string]string{}

	for i := 0; i < len(args); {
		array := argumentOptionRegex.FindStringSubmatch(args[i])
		if len(array) > 2 {
			if array[2] == "=" {
				parsedOptions[array[1]] = array[3]
			} else if i < len(args)-1 && len(args[i+1]) > 0 && args[i+1][0] != '-' {
				parsedOptions[array[1]] = args[i+1]
				i += 2
				continue
			} else {
				parsedOptions[array[1]] = array[3]
			}
		} else {
			parsedArgs = append(parsedArgs, args[i])
		}
		i++
	}
	return
}

// GetOpt retrieves the value of the specified option.
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

// GetOptAll returns all parsed options as a map.
func GetOptAll() map[string]string {
	Init()
	return defaultParsedOptions
}

// ContainsOpt checks if a specific option exists in the parsed options.
func ContainsOpt(name string) bool {
	Init()
	_, ok := defaultParsedOptions[name]
	return ok
}

// GetArg retrieves the argument at the specified index.
// If the index is out of bounds, it returns the default value if provided.
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

// GetArgAll returns all parsed arguments as a slice.
func GetArgAll() []string {
	Init()
	return defaultParsedArgs
}

// GetOptWithEnv retrieves a value from the command-line options or environment variables.
// If neither exists, it returns the default value if provided.
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
