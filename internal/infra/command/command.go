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
	"regexp"
)

// Global variables for storing parsed arguments and options.
var (
	defaultParsedArgs    []string
	defaultParsedOptions map[string]string
	argumentOptionRegex  = regexp.MustCompile(`^\-{1,2}([\w\?\.\-]+)(=)?(.*)$`)
)
