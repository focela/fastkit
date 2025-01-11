// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package debugger provides utilities for debugging, including logging and tracking application state.
package debugger

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

// Constants for stack and filter configuration.
const (
	maxCallerDepth = 1000
	stackFilterKey = "/debugger/debugger"
)

// Global variables for path and version information.
var (
	goRootForFilter  = runtime.GOROOT()
	binaryVersion    string
	binaryVersionMd5 string
	selfPath         string
)

// Initialize global variables.
func init() {
	if goRootForFilter != "" {
		goRootForFilter = strings.ReplaceAll(goRootForFilter, "\\", "/")
	}

	selfPath, _ = exec.LookPath(os.Args[0])
	if selfPath != "" {
		selfPath, _ = filepath.Abs(selfPath)
	}
	if selfPath == "" {
		selfPath, _ = filepath.Abs(os.Args[0])
	}
}

// Caller returns the function name, file path, and line number of the caller.
func Caller(skip ...int) (function string, path string, line int) {
	return CallerWithFilter(nil, skip...)
}

// CallerWithFilter returns caller information while applying file path filters.
func CallerWithFilter(filters []string, skip ...int) (function string, path string, line int) {
	number := 0
	if len(skip) > 0 {
		number = skip[0]
	}

	pc, file, line, start := callerFromIndex(filters)
	if start != -1 {
		var ok bool
		for i := start + number; i < maxCallerDepth; i++ {
			if i != start {
				pc, file, line, ok = runtime.Caller(i)
				if !ok {
					break
				}
			}
			if filterFileByFilters(file, filters) {
				continue
			}

			function = "unknown"
			if fn := runtime.FuncForPC(pc); fn != nil {
				function = fn.Name()
			}
			return function, file, line
		}
	}
	return "", "", -1
}

// callerFromIndex determines the starting index for the caller stack trace.
func callerFromIndex(filters []string) (pc uintptr, file string, line int, index int) {
	var ok bool
	for index = 0; index < maxCallerDepth; index++ {
		if pc, file, line, ok = runtime.Caller(index); ok {
			if filterFileByFilters(file, filters) {
				continue
			}
			if index > 0 {
				index--
			}
			return
		}
	}
	return 0, "", -1, -1
}

// filterFileByFilters filters files based on the provided filters and stack configuration.
func filterFileByFilters(file string, filters []string) (filtered bool) {
	if file == "" || strings.Contains(file, stackFilterKey) {
		return true
	}

	for _, filter := range filters {
		if filter != "" && strings.Contains(file, filter) {
			return true
		}
	}

	if goRootForFilter != "" && strings.HasPrefix(file, goRootForFilter) {
		fileSeparator := file[len(goRootForFilter)]
		if fileSeparator == filepath.Separator || fileSeparator == '\\' || fileSeparator == '/' {
			return true
		}
	}
	return false
}

// CallerPackage returns the package name of the caller.
func CallerPackage() string {
	function, _, _ := Caller()
	return getPackageFromCallerFunction(function)
}

// getPackageFromCallerFunction extracts the package name from the function name.
func getPackageFromCallerFunction(function string) string {
	indexSlash := strings.LastIndexByte(function, '/')
	if indexSlash == -1 {
		return function[:strings.IndexByte(function, '.')]
	}
	leftPart := function[:indexSlash+1]
	rightPart := function[indexSlash+1:]
	indexDot := strings.IndexByte(rightPart, '.')
	if indexDot >= 0 {
		rightPart = rightPart[:indexDot]
	}
	return leftPart + rightPart
}

// CallerFunction returns the function name of the caller.
func CallerFunction() string {
	function, _, _ := Caller()
	function = function[strings.LastIndexByte(function, '/')+1:]
	function = function[strings.IndexByte(function, '.')+1:]
	return function
}

// CallerFilePath returns the file path of the caller.
func CallerFilePath() string {
	_, path, _ := Caller()
	return path
}

// CallerDirectory returns the directory of the caller.
func CallerDirectory() string {
	_, path, _ := Caller()
	return filepath.Dir(path)
}

// CallerFileLine returns the file path with the line number of the caller.
func CallerFileLine() string {
	_, path, line := Caller()
	return fmt.Sprintf("%s:%d", path, line)
}

// CallerFileLineShort returns the file name with the line number of the caller.
func CallerFileLineShort() string {
	_, path, line := Caller()
	return fmt.Sprintf("%s:%d", filepath.Base(path), line)
}

// FuncPath returns the complete function path of the provided function.
func FuncPath(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

// FuncName returns the function name of the provided function.
func FuncName(f interface{}) string {
	path := FuncPath(f)
	if path == "" {
		return ""
	}
	index := strings.LastIndexByte(path, '/')
	if index < 0 {
		index = strings.LastIndexByte(path, '\\')
	}
	return path[index+1:]
}
