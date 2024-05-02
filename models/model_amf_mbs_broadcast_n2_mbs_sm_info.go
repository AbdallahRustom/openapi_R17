/*
 * Namf_MBSBroadcast
 *
 * AMF MBSBroadcast Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.518 V17.12.0; 5G System; Access and Mobility Management Services
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.518/
 *
 * API version: 1.0.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// N2 MBS Session Management information
type AmfMbsBroadcastN2MbsSmInfo struct {
	NgapIeType AmfMbsBroadcastNgapIeType `json:"ngapIeType" yaml:"ngapIeType" bson:"ngapIeType,omitempty"`
	NgapData   *RefToBinaryData          `json:"ngapData" yaml:"ngapData" bson:"ngapData,omitempty"`
	RanId      *GlobalRanNodeId          `json:"ranId,omitempty" yaml:"ranId" bson:"ranId,omitempty"`
}
