// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package errors provides rich functionalities to manipulate errors.
package errors

import (
	"bytes"
	"container/list"
	"fmt"
	"runtime"
	"strings"

	"github.com/focela/loom/internal/config"
	"github.com/focela/loom/internal/core"
)

// stackInfo manages stack info of certain error.
type stackInfo struct {
	Index   int        // Index of current error in the stack trace.
	Message string     // Error message.
	Lines   *list.List // List of stack lines for this error.
}

// stackLine manages each line info of stack.
type stackLine struct {
	Function string // Function name, including full package path.
	FileLine string // File name and line number.
}

// Stack returns the error stack trace as a string.
func (err *Error) Stack() string {
	if err == nil {
		return ""
	}

	var (
		loop             = err
		index            = 1
		infos            []*stackInfo
		isStackModeBrief = core.IsStackModeBrief()
	)

	for loop != nil {
		info := &stackInfo{
			Index:   index,
			Message: fmt.Sprintf("%v", loop),
		}
		index++
		infos = append(infos, info)
		loopLinesOfStackInfo(loop.stack, info, isStackModeBrief)

		if nestedErr, ok := loop.error.(*Error); ok {
			loop = nestedErr
		} else if loop.error != nil {
			infos = append(infos, &stackInfo{
				Index:   index,
				Message: loop.error.Error(),
			})
			break
		} else {
			break
		}
	}

	filterLinesOfStackInfos(infos)
	return formatStackInfos(infos)
}

// loopLinesOfStackInfo iterates through stack info lines and extracts stack trace details.
func loopLinesOfStackInfo(st stack, info *stackInfo, isStackModeBrief bool) {
	if st == nil {
		return
	}

	for _, p := range st {
		if fn := runtime.FuncForPC(p - 1); fn != nil {
			file, line := fn.FileLine(p - 1)

			// Skip stack paths based on configuration
			if isStackModeBrief && strings.Contains(file, config.StackFilterKeyForLoom) {
				continue
			}
			if strings.Contains(file, stackFilterKeyLocal) || strings.Contains(file, "<") {
				continue
			}
			if goRootForFilter != "" && strings.HasPrefix(file, goRootForFilter) {
				continue
			}

			// Add stack line
			if info.Lines == nil {
				info.Lines = list.New()
			}
			info.Lines.PushBack(&stackLine{
				Function: fn.Name(),
				FileLine: fmt.Sprintf(`%s:%d`, file, line),
			})
		}
	}
}

// filterLinesOfStackInfos removes duplicate lines across multiple error stacks.
func filterLinesOfStackInfos(infos []*stackInfo) {
	set := make(map[string]struct{})

	for i := len(infos) - 1; i >= 0; i-- {
		info := infos[i]
		if info.Lines == nil {
			continue
		}

		var removes []*list.Element
		for e := info.Lines.Front(); e != nil; e = e.Next() {
			line := e.Value.(*stackLine)
			if _, exists := set[line.FileLine]; exists {
				removes = append(removes, e)
			} else {
				set[line.FileLine] = struct{}{}
			}
		}

		// Remove duplicates
		for _, e := range removes {
			info.Lines.Remove(e)
		}
	}
}

// formatStackInfos formats the error stack trace into a readable string.
func formatStackInfos(infos []*stackInfo) string {
	buffer := bytes.NewBuffer(nil)

	for i, info := range infos {
		buffer.WriteString(fmt.Sprintf("%d. %s\n", i+1, info.Message))
		if info.Lines != nil && info.Lines.Len() > 0 {
			formatStackLines(buffer, info.Lines)
		}
	}
	return buffer.String()
}

// formatStackLines formats and appends stack lines to the buffer.
func formatStackLines(buffer *bytes.Buffer, lines *list.List) {
	space := "  "
	for i, e := 0, lines.Front(); e != nil; i, e = i+1, e.Next() {
		line := e.Value.(*stackLine)

		if i >= 9 {
			space = " "
		}

		buffer.WriteString(fmt.Sprintf(
			"   %d).%s%s\n        %s\n",
			i+1, space, line.Function, line.FileLine,
		))
	}
}
