/*
 * Namf_Communication
 *
 * AMF Communication Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.518 V17.12.0; 5G System; Access and Mobility Management Services
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.518/
 *
 * API version: 1.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// UE policy delivery related N1 message notification subscription data.
type UpdpSubscriptionData struct {
	UpdpNotifySubscriptionId string `json:"updpNotifySubscriptionId" yaml:"updpNotifySubscriptionId" bson:"updpNotifySubscriptionId,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	UpdpNotifyCallbackUri string `json:"updpNotifyCallbackUri" yaml:"updpNotifyCallbackUri" bson:"updpNotifyCallbackUri,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures   string `json:"supportedFeatures,omitempty" yaml:"supportedFeatures" bson:"supportedFeatures,omitempty"`
	UpdpCallbackBinding string `json:"updpCallbackBinding,omitempty" yaml:"updpCallbackBinding" bson:"updpCallbackBinding,omitempty"`
}
