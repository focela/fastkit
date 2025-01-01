// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package debug provides utilities for managing debug mode and retrieving caller details.
package debug

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

const (
	maxCallerDepth = 1000            // Maximum stack depth for caller tracing.
	stackFilterKey = "/debug/gdebug" // Filter key for stack trace filtering.
)

var (
	goRootForFilter  = runtime.GOROOT() // Used for stack filtering.
	binaryVersion    string             // Current binary version (uint64 hex).
	binaryVersionMd5 string             // Current binary version (MD5 hash).
	selfPath         string             // Absolute path to the current running binary.
)

// Initializes debug variables like `selfPath` and normalizes `goRootForFilter`.
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

// Caller retrieves the function name, file path, and line number of the caller.
func Caller(skip ...int) (function string, path string, line int) {
	return CallerWithFilter(nil, skip...)
}

// CallerWithFilter retrieves the caller's details with optional path filtering.
func CallerWithFilter(filters []string, skip ...int) (function string, path string, line int) {
	number := 0
	if len(skip) > 0 {
		number = skip[0]
	}

	_, _, line, start := callerFromIndex(filters)
	if start == -1 {
		return "", "", -1
	}

	for i := start + number; i < maxCallerDepth; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		if filterFileByFilters(file, filters) {
			continue
		}

		fn := runtime.FuncForPC(reflect.ValueOf(i).Pointer())
		if fn == nil {
			function = "unknown"
		} else {
			function = fn.Name()
		}
		return function, file, line
	}
	return "", "", -1
}

// callerFromIndex finds the initial valid caller index, skipping filters.
func callerFromIndex(filters []string) (pc uintptr, file string, line int, index int) {
	for index = 0; index < maxCallerDepth; index++ {
		pc, file, line, ok := runtime.Caller(index)
		if !ok {
			break
		}
		if filterFileByFilters(file, filters) {
			continue
		}
		if index > 0 {
			index--
		}
		return pc, file, line, index
	}
	return 0, "", -1, -1
}

// filterFileByFilters filters stack trace files based on filters.
func filterFileByFilters(file string, filters []string) bool {
	if file == "" || strings.Contains(file, stackFilterKey) {
		return true
	}
	for _, filter := range filters {
		if filter != "" && strings.Contains(file, filter) {
			return true
		}
	}
	if goRootForFilter != "" && strings.HasPrefix(file, goRootForFilter) {
		separator := file[len(goRootForFilter)]
		if separator == filepath.Separator || separator == '\\' || separator == '/' {
			return true
		}
	}
	return false
}

// CallerPackage retrieves the package name of the caller.
func CallerPackage() string {
	function, _, _ := Caller()
	return getPackageFromCallerFunction(function)
}

// getPackageFromCallerFunction extracts the package name from the function path.
func getPackageFromCallerFunction(function string) string {
	indexSplit := strings.LastIndexByte(function, '/')
	if indexSplit == -1 {
		return function[:strings.IndexByte(function, '.')]
	}
	leftPart := function[:indexSplit+1]
	rightPart := function[indexSplit+1:]
	indexDot := strings.IndexByte(rightPart, '.')
	if indexDot >= 0 {
		rightPart = rightPart[:indexDot]
	}
	return leftPart + rightPart
}

// CallerFunction retrieves the function name of the caller.
func CallerFunction() string {
	function, _, _ := Caller()
	function = function[strings.LastIndexByte(function, '/')+1:]
	function = function[strings.IndexByte(function, '.')+1:]
	return function
}

// CallerFilePath retrieves the file path of the caller.
func CallerFilePath() string {
	_, path, _ := Caller()
	return path
}

// CallerDirectory retrieves the directory of the caller.
func CallerDirectory() string {
	_, path, _ := Caller()
	return filepath.Dir(path)
}

// CallerFileLine retrieves the file path and line number of the caller.
func CallerFileLine() string {
	_, path, line := Caller()
	return fmt.Sprintf(`%s:%d`, path, line)
}

// CallerFileLineShort retrieves the file name and line number of the caller.
func CallerFileLineShort() string {
	_, path, line := Caller()
	return fmt.Sprintf(`%s:%d`, filepath.Base(path), line)
}

// FuncPath retrieves the full function path of the given function.
func FuncPath(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

// FuncName retrieves the function name of the given function.
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
