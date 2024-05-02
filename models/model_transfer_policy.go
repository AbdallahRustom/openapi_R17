/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 29.504 V16.9.0; 5G System; Unified Data Repository Services; Stage 3
 * Url: http://www.3gpp.org/ftp/Specs/archive/29_series/29.504/
 *
 * API version: 2.1.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Describes a transfer policy.
type TransferPolicy struct {
	MaxBitRateDl string `json:"maxBitRateDl,omitempty" yaml:"maxBitRateDl" bson:"maxBitRateDl,omitempty"`
	MaxBitRateUl string `json:"maxBitRateUl,omitempty" yaml:"maxBitRateUl" bson:"maxBitRateUl,omitempty"`
	// Indicates a rating group for the recommended time window.
	RatingGroup int32       `json:"ratingGroup" yaml:"ratingGroup" bson:"ratingGroup,omitempty"`
	RecTimeInt  *TimeWindow `json:"recTimeInt" yaml:"recTimeInt" bson:"recTimeInt,omitempty"`
	// Contains an identity of a transfer policy.
	TransPolicyId int32 `json:"transPolicyId" yaml:"transPolicyId" bson:"transPolicyId,omitempty"`
}