/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 29.510 V16.9.0; 5G System; Network Function Repository Services; Stage 3
 * Url: http://www.3gpp.org/ftp/Specs/archive/29_series/29.510/
 *
 * API version: 1.1.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type NfStatus string

// List of NFStatus
const (
	NfStatus_REGISTERED     NfStatus = "REGISTERED"
	NfStatus_SUSPENDED      NfStatus = "SUSPENDED"
	NfStatus_UNDISCOVERABLE NfStatus = "UNDISCOVERABLE"
)