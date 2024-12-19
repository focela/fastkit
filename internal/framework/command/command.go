// Copyright (c) 2024 Focela Technologies
// Internal Use Only - Unauthorized use prohibited
// Contact: legal@focela.com

package command

import (
	"os"
	"regexp"
	"strings"
)

// Default variables for parsed arguments and options.
var (
	defaultParsedArgs    = make([]string, 0)
	defaultParsedOptions = make(map[string]string)
	argumentOptionRegex  = regexp.MustCompile(`^\-{1,2}([\w\?\.\-]+)(=)?(.*)$`)
)

// Init initializes the argument and option parsing process.
// If no arguments are provided, it uses `os.Args` as default input.
func Init(args ...string) {
	if len(args) == 0 {
		if len(defaultParsedArgs) == 0 && len(defaultParsedOptions) == 0 {
			args = os.Args
		} else {
			return
		}
	} else {
		// Reset default parsed arguments and options.
		defaultParsedArgs = make([]string, 0)
		defaultParsedOptions = make(map[string]string)
	}
	defaultParsedArgs, defaultParsedOptions = ParseUsingDefaultAlgorithm(args...)
}

// ParseUsingDefaultAlgorithm parses arguments into positional arguments and options.
// Example:
//
//	Input: ["-a", "--key=value", "arg1"]
//	Output: parsedArgs = ["arg1"], parsedOptions = {"a": "", "key": "value"}
func ParseUsingDefaultAlgorithm(args ...string) (parsedArgs []string, parsedOptions map[string]string) {
	parsedArgs = make([]string, 0)
	parsedOptions = make(map[string]string)

	for i := 0; i < len(args); {
		matches := argumentOptionRegex.FindStringSubmatch(args[i])
		if len(matches) > 2 {
			if matches[2] == "=" {
				parsedOptions[matches[1]] = matches[3]
			} else if i < len(args)-1 && len(args[i+1]) > 0 && args[i+1][0] != '-' {
				parsedOptions[matches[1]] = args[i+1]
				i += 2
				continue
			} else {
				parsedOptions[matches[1]] = matches[3]
			}
		} else {
			parsedArgs = append(parsedArgs, args[i])
		}
		i++
	}
	return
}

// GetOpt retrieves the value of the option with the specified `name`.
// Returns the default value `def` if the option does not exist.
func GetOpt(name string, def ...string) string {
	Init()
	if value, exists := defaultParsedOptions[name]; exists {
		return value
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

// ContainsOpt checks if an option with the specified `name` exists.
func ContainsOpt(name string) bool {
	Init()
	_, exists := defaultParsedOptions[name]
	return exists
}

// GetArg retrieves the argument at the specified `index`.
// Returns the default value `def` if the index is out of range.
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
// - Command-line arguments are in lowercase format (e.g., `key.option`).
// - Environment variables are in uppercase format (e.g., `KEY_OPTION`).
// Returns the default value `def` if neither exists.
func GetOptWithEnv(key string, def ...string) string {
	cmdKey := strings.ToLower(strings.ReplaceAll(key, "_", "."))
	if ContainsOpt(cmdKey) {
		return GetOpt(cmdKey)
	}

	envKey := strings.ToUpper(strings.ReplaceAll(key, ".", "_"))
	if value, exists := os.LookupEnv(envKey); exists {
		return value
	}

	if len(def) > 0 {
		return def[0]
	}
	return ""
}
