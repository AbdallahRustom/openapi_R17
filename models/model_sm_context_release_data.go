/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 29.502 V16.9.0; 5G System; Session Management Services; Stage 3
 * Url: http://www.3gpp.org/ftp/Specs/archive/29_series/29.502/
 *
 * API version: 1.1.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type SmContextReleaseData struct {
	Cause             Cause            `json:"cause,omitempty" yaml:"cause" bson:"cause,omitempty"`
	NgApCause         *NgApCause       `json:"ngApCause,omitempty" yaml:"ngApCause" bson:"ngApCause,omitempty"`
	Var5gMmCauseValue int32            `json:"5gMmCauseValue,omitempty" yaml:"5gMmCauseValue" bson:"5gMmCauseValue,omitempty"`
	UeLocation        *UserLocation    `json:"ueLocation,omitempty" yaml:"ueLocation" bson:"ueLocation,omitempty"`
	UeTimeZone        string           `json:"ueTimeZone,omitempty" yaml:"ueTimeZone" bson:"ueTimeZone,omitempty"`
	AddUeLocation     *UserLocation    `json:"addUeLocation,omitempty" yaml:"addUeLocation" bson:"addUeLocation,omitempty"`
	VsmfReleaseOnly   bool             `json:"vsmfReleaseOnly,omitempty" yaml:"vsmfReleaseOnly" bson:"vsmfReleaseOnly,omitempty"`
	N2SmInfo          *RefToBinaryData `json:"n2SmInfo,omitempty" yaml:"n2SmInfo" bson:"n2SmInfo,omitempty"`
	N2SmInfoType      N2SmInfoType     `json:"n2SmInfoType,omitempty" yaml:"n2SmInfoType" bson:"n2SmInfoType,omitempty"`
	IsmfReleaseOnly   bool             `json:"ismfReleaseOnly,omitempty" yaml:"ismfReleaseOnly" bson:"ismfReleaseOnly,omitempty"`
}
