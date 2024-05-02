/*
 * Nudm_EE
 *
 * Nudm Event Exposure Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 29.503 Unified Data Management Services, version 16.9.0
 * Url: http://www.3gpp.org/ftp/Specs/archive/29_series/29.503/
 *
 * API version: 1.1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type CreatedEeSubscription struct {
	EeSubscription *EeSubscription    `json:"eeSubscription" yaml:"eeSubscription" bson:"eeSubscription,omitempty"`
	NumberOfUes    int32              `json:"numberOfUes,omitempty" yaml:"numberOfUes" bson:"numberOfUes,omitempty"`
	EventReports   []MonitoringReport `json:"eventReports,omitempty" yaml:"eventReports" bson:"eventReports,omitempty"`
	EpcStatusInd   bool               `json:"epcStatusInd,omitempty" yaml:"epcStatusInd" bson:"epcStatusInd,omitempty"`
}
