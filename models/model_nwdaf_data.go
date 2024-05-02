/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.512 V17.11.0; 5G System; Session Management Policy Control Service.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.512/
 *
 * API version: 1.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Indicates the list of Analytic ID(s) per NWDAF instance ID used for the PDU Session consumed by the SMF.
type NwdafData struct {
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	NwdafInstanceId string       `json:"nwdafInstanceId" yaml:"nwdafInstanceId" bson:"nwdafInstanceId,omitempty"`
	NwdafEvents     []NwdafEvent `json:"nwdafEvents,omitempty" yaml:"nwdafEvents" bson:"nwdafEvents,omitempty"`
}
