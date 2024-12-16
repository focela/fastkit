/*
 * Copyright © 2024 Focela Technologies. All rights reserved.
 *
 * This source code is provided for viewing purposes only. Copying, modification,
 * distribution, or use of this code is strictly prohibited without explicit
 * written permission from Focela Technologies.
 *
 * This code is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
 * either express or implied. For more information, see the LICENSE file or
 * contact legal@focela.com.
 */

// Package command provides utilities for console operations like options/arguments parsing.
package command

import (
	"os"
	"regexp"
	"strings"
)

// Global variables for storing parsed arguments and options.
var (
	defaultParsedArgs    []string
	defaultParsedOptions map[string]string
	argumentOptionRegex  = regexp.MustCompile(`^\-{1,2}([\w\?\.\-]+)(=)?(.*)$`)
)

// Init initializes defaultParsedArgs and defaultParsedOptions.
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
	// Parse arguments using the default algorithm.
	defaultParsedArgs, defaultParsedOptions = parseUsingDefaultAlgorithm(args...)
}

// parseUsingDefaultAlgorithm parses command-line arguments into args and options.
func parseUsingDefaultAlgorithm(args ...string) ([]string, map[string]string) {
	parsedArgs := []string{}
	parsedOptions := make(map[string]string)

	for i := 0; i < len(args); {
		matches := argumentOptionRegex.FindStringSubmatch(args[i])
		if len(matches) > 2 {
			key := matches[1]
			if matches[2] == "=" {
				parsedOptions[key] = matches[3]
			} else if i < len(args)-1 && !strings.HasPrefix(args[i+1], "-") {
				parsedOptions[key] = args[i+1]
				i += 2
				continue
			} else {
				parsedOptions[key] = matches[3]
			}
		} else {
			parsedArgs = append(parsedArgs, args[i])
		}
		i++
	}
	return parsedArgs, parsedOptions
}
