/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 32.291 V17.9.0: Telecommunication management; Charging management;  5G system, charging service; Stage 3.
 * Url: http://www.3gpp.org/ftp/Specs/archive/32_series/32.291/
 *
 * API version: 3.1.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

type ChfConvergedChargingChargingDataResponse struct {
	// string with format 'date-time' as defined in OpenAPI.
	InvocationTimeStamp *time.Time `json:"invocationTimeStamp" yaml:"invocationTimeStamp" bson:"invocationTimeStamp,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.
	InvocationSequenceNumber int32             `json:"invocationSequenceNumber" yaml:"invocationSequenceNumber" bson:"invocationSequenceNumber,omitempty"`
	InvocationResult         *InvocationResult `json:"invocationResult,omitempty" yaml:"invocationResult" bson:"invocationResult,omitempty"`
	SessionFailover          SessionFailover   `json:"sessionFailover,omitempty" yaml:"sessionFailover" bson:"sessionFailover,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures                    string                                             `json:"supportedFeatures,omitempty" yaml:"supportedFeatures" bson:"supportedFeatures,omitempty"`
	MultipleUnitInformation              []MultipleUnitInformation                          `json:"multipleUnitInformation,omitempty" yaml:"multipleUnitInformation" bson:"multipleUnitInformation,omitempty"`
	Triggers                             []ChfConvergedChargingTrigger                      `json:"triggers,omitempty" yaml:"triggers" bson:"triggers,omitempty"`
	PDUSessionChargingInformation        *ChfConvergedChargingPduSessionChargingInformation `json:"pDUSessionChargingInformation,omitempty" yaml:"pDUSessionChargingInformation" bson:"pDUSessionChargingInformation,omitempty"`
	RoamingQBCInformation                *ChfConvergedChargingRoamingQbcInformation         `json:"roamingQBCInformation,omitempty" yaml:"roamingQBCInformation" bson:"roamingQBCInformation,omitempty"`
	LocationReportingChargingInformation *LocationReportingChargingInformation              `json:"locationReportingChargingInformation,omitempty" yaml:"locationReportingChargingInformation" bson:"locationReportingChargingInformation,omitempty"`
}
