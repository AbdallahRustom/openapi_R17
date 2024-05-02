/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 29.504 V16.9.0; 5G System; Unified Data Repository Services; Stage 3
 * Url: http://www.3gpp.org/ftp/Specs/archive/29_series/29.504/
 *
 * API version: 2.1.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type AuthType string

// List of AuthType
const (
	AuthType__5_G_AKA      AuthType = "5G_AKA"
	AuthType_EAP_AKA_PRIME AuthType = "EAP_AKA_PRIME"
	AuthType_EAP_TLS       AuthType = "EAP_TLS"
)