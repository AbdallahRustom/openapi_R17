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

// Data within an AMF Status Change Subscription request and response
type AmfCommunicationSubscriptionData struct {
	// String providing an URI formatted according to RFC 3986.
	AmfStatusUri string  `json:"amfStatusUri" yaml:"amfStatusUri" bson:"amfStatusUri,omitempty"`
	GuamiList    []Guami `json:"guamiList,omitempty" yaml:"guamiList" bson:"guamiList,omitempty"`
}
