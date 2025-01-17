/*
 * 3gpp-as-session-with-qos
 *
 * API for setting us an AS session with required QoS.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.122 V17.9.0 T8 reference point for Northbound APIs
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.122/
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Represents a QoS monitoring report.
type QosMonitoringReport struct {
	UlDelays []int32 `json:"ulDelays,omitempty" yaml:"ulDelays" bson:"ulDelays,omitempty"`
	DlDelays []int32 `json:"dlDelays,omitempty" yaml:"dlDelays" bson:"dlDelays,omitempty"`
	RtDelays []int32 `json:"rtDelays,omitempty" yaml:"rtDelays" bson:"rtDelays,omitempty"`
	// Represents the packet delay measurement failure indicator.
	Pdmf bool `json:"pdmf,omitempty" yaml:"pdmf" bson:"pdmf,omitempty"`
}
