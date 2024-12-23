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

// Package config provides configuration constants shared across the framework.
package config

// Configuration constants shared across different services and modules.
const (
	// ConfigNodeNameDatabase represents the configuration node name for database settings.
	ConfigNodeNameDatabase = "database"

	// ConfigNodeNameLogger represents the configuration node name for logger settings.
	ConfigNodeNameLogger = "logger"

	// ConfigNodeNameRedis represents the configuration node name for Redis settings.
	ConfigNodeNameRedis = "redis"

	// ConfigNodeNameViewer represents the configuration node name for viewer settings.
	ConfigNodeNameViewer = "viewer"

	// ConfigNodeNameServer represents the configuration node name for primary server settings.
	ConfigNodeNameServer = "server"

	// ConfigNodeNameServerSecondary represents the configuration node name for secondary server settings.
	ConfigNodeNameServerSecondary = "httpserver"

	// StackFilterKeyForAltura represents the stack filter key used in the Altura package.
	StackFilterKeyForAltura = "github.com/focela/altura/"
)
