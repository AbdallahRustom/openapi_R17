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

type LocationAccuracy string

// List of LocationAccuracy
const (
	LocationAccuracy_CELL_LEVEL   LocationAccuracy = "CELL_LEVEL"
	LocationAccuracy_TA_LEVEL     LocationAccuracy = "TA_LEVEL"
	LocationAccuracy_N3_IWF_LEVEL LocationAccuracy = "N3IWF_LEVEL"
	LocationAccuracy_UE_IP        LocationAccuracy = "UE_IP"
	LocationAccuracy_UE_PORT      LocationAccuracy = "UE_PORT"
)