/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.502 V17.11.0; 5G System; Session Management Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.502/
 *
 * API version: 1.2.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// EPS PDN Connection Information from H-SMF to V-SMF, or from SMF to I-SMF
type EpsPdnCnxInfo struct {
	// string with format 'bytes' as defined in OpenAPI
	PgwS8cFteid string `json:"pgwS8cFteid" yaml:"pgwS8cFteid" bson:"pgwS8cFteid,omitempty"`
	// string with format 'bytes' as defined in OpenAPI
	PgwNodeName string `json:"pgwNodeName,omitempty" yaml:"pgwNodeName" bson:"pgwNodeName,omitempty"`
	// EPS Bearer Identifier
	LinkedBearerId int32 `json:"linkedBearerId,omitempty" yaml:"linkedBearerId" bson:"linkedBearerId,omitempty"`
}
