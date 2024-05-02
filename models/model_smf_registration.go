/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.504 V17.12.0; 5G System; Unified Data Repository Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.504/
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

type SmfRegistration struct {
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	SmfInstanceId string `json:"smfInstanceId" yaml:"smfInstanceId" bson:"smfInstanceId,omitempty"`
	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.
	SmfSetId string `json:"smfSetId,omitempty" yaml:"smfSetId" bson:"smfSetId,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures string `json:"supportedFeatures,omitempty" yaml:"supportedFeatures" bson:"supportedFeatures,omitempty"`
	// Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.
	PduSessionId int32   `json:"pduSessionId" yaml:"pduSessionId" bson:"pduSessionId,omitempty"`
	SingleNssai  *Snssai `json:"singleNssai" yaml:"singleNssai" bson:"singleNssai,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn               string `json:"dnn,omitempty" yaml:"dnn" bson:"dnn,omitempty"`
	EmergencyServices bool   `json:"emergencyServices,omitempty" yaml:"emergencyServices" bson:"emergencyServices,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	PcscfRestorationCallbackUri string  `json:"pcscfRestorationCallbackUri,omitempty" yaml:"pcscfRestorationCallbackUri" bson:"pcscfRestorationCallbackUri,omitempty"`
	PlmnId                      *PlmnId `json:"plmnId" yaml:"plmnId" bson:"plmnId,omitempty"`
	// Fully Qualified Domain Name
	PgwFqdn   string           `json:"pgwFqdn,omitempty" yaml:"pgwFqdn" bson:"pgwFqdn,omitempty"`
	PgwIpAddr *UdmSdmIpAddress `json:"pgwIpAddr,omitempty" yaml:"pgwIpAddr" bson:"pgwIpAddr,omitempty"`
	EpdgInd   bool             `json:"epdgInd,omitempty" yaml:"epdgInd" bson:"epdgInd,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	DeregCallbackUri   string             `json:"deregCallbackUri,omitempty" yaml:"deregCallbackUri" bson:"deregCallbackUri,omitempty"`
	RegistrationReason RegistrationReason `json:"registrationReason,omitempty" yaml:"registrationReason" bson:"registrationReason,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	RegistrationTime *time.Time   `json:"registrationTime,omitempty" yaml:"registrationTime" bson:"registrationTime,omitempty"`
	ContextInfo      *ContextInfo `json:"contextInfo,omitempty" yaml:"contextInfo" bson:"contextInfo,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	PcfId string `json:"pcfId,omitempty" yaml:"pcfId" bson:"pcfId,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	DataRestorationCallbackUri string   `json:"dataRestorationCallbackUri,omitempty" yaml:"dataRestorationCallbackUri" bson:"dataRestorationCallbackUri,omitempty"`
	ResetIds                   []string `json:"resetIds,omitempty" yaml:"resetIds" bson:"resetIds,omitempty"`
	UdrRestartInd              bool     `json:"udrRestartInd,omitempty" yaml:"udrRestartInd" bson:"udrRestartInd,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	LastSynchronizationTime *time.Time `json:"lastSynchronizationTime,omitempty" yaml:"lastSynchronizationTime" bson:"lastSynchronizationTime,omitempty"`
}
