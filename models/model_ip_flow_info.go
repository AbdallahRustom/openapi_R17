/*
 * Npcf_EventExposure
 *
 * PCF Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.523 V17.7.0; 5G System; Policy Control Event Exposure Service; Stage 3.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.523/
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Identifies an UL/DL IP flow.
type IpFlowInfo struct {
	IpFlows    []string `json:"ipFlows,omitempty" yaml:"ipFlows" bson:"ipFlows,omitempty"`
	FlowNumber int32    `json:"flowNumber" yaml:"flowNumber" bson:"flowNumber,omitempty"`
}
