// Copyright 2025 Focela – MIT Licensed. See LICENSE.
// Part of the Focela Open Initiative.
// Built to empower developers and partners.

// Package consts defines system-wide constants
package consts

// System defines general purpose constants used across the application.
const (
	DemoTips              = "Sensitive Data Hidden In Demo" // Mask for sensitive data in demo mode
	NilJSONToString       = "{}"                            // Empty JSON string
	RegionSplit           = " / "                           // Region delimiter
	Unknown               = "Unknown"                       // Default for unknown values
	SuperRoleKey          = "super"                         // Admin role identifier
	MaxServeLogContentLen = 2048                            // Max service log length
)

// CRUD defines constants for pagination and sorting operations.
const (
	DefaultPage      = 10 // Items per page default
	DefaultPageSize  = 1  // Starting page number
	MaxSortIncrement = 10 // Max sort value increment
)

// TenantField defines field names used for multi-tenancy support.
const (
	TenantID   = "tenant_id"   // Tenant ID field name
	MerchantID = "merchant_id" // Merchant ID field name
	UserID     = "user_id"     // User ID field name
)
