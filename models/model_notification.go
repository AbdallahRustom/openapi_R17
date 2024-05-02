/*
 * Npcf_BDTPolicyControl Service API
 *
 * PCF BDT Policy Control Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 29.554 V16.7.0; 5G System; Background Data Transfer Policy Control Service.
 * Url: http://www.3gpp.org/ftp/Specs/archive/29_series/29.554/
 *
 * API version: 1.1.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Describes a BDT notification.
type Notification struct {
	// string identifying a BDT Reference ID as defined in subclause 5.3.3 of 3GPP TS 29.154.
	BdtRefId string `json:"bdtRefId" yaml:"bdtRefId" bson:"bdtRefId,omitempty"`
	// Contains a list of the candidate transfer policies from which the AF may select a new transfer policy due to a network performance is below the criteria set by the operator.
	CandPolicies []TransferPolicy `json:"candPolicies,omitempty" yaml:"candPolicies" bson:"candPolicies,omitempty"`
	NwAreaInfo   *NetworkAreaInfo `json:"nwAreaInfo,omitempty" yaml:"nwAreaInfo" bson:"nwAreaInfo,omitempty"`
	TimeWindow   *TimeWindow      `json:"timeWindow,omitempty" yaml:"timeWindow" bson:"timeWindow,omitempty"`
}
