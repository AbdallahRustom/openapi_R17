/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.520 V17.9.0; 5G System; Network Data Analytics Services.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.520/
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

type UdmEeMonitoringReport struct {
	ReferenceId              int32                          `json:"referenceId" yaml:"referenceId" bson:"referenceId,omitempty"`
	EventType                UdmEeEventType                 `json:"eventType" yaml:"eventType" bson:"eventType,omitempty"`
	Report                   *UdmEeReport                   `json:"report,omitempty" yaml:"report" bson:"report,omitempty"`
	ReachabilityForSmsReport *UdmEeReachabilityForSmsReport `json:"reachabilityForSmsReport,omitempty" yaml:"reachabilityForSmsReport" bson:"reachabilityForSmsReport,omitempty"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.
	Gpsi string `json:"gpsi,omitempty" yaml:"gpsi" bson:"gpsi,omitempty"`
	// string with format \"date-time\" as defined in OpenAPI.
	TimeStamp          *time.Time          `json:"timeStamp" yaml:"timeStamp" bson:"timeStamp,omitempty"`
	ReachabilityReport *ReachabilityReport `json:"reachabilityReport,omitempty" yaml:"reachabilityReport" bson:"reachabilityReport,omitempty"`
}
