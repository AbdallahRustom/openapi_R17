/*
 * Nhss_imsSDM
 *
 * Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.562 HSS Services for interworking with IMS, version 17.6.0
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.562/
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type IdentityType string

// List of IdentityType
const (
	IdentityType_DISTINCT_IMPU   IdentityType = "DISTINCT_IMPU"
	IdentityType_DISTINCT_PSI    IdentityType = "DISTINCT_PSI"
	IdentityType_WILDCARDED_IMPU IdentityType = "WILDCARDED_IMPU"
	IdentityType_WILDCARDED_PSI  IdentityType = "WILDCARDED_PSI"
)
