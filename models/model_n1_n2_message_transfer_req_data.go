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

// Data within a N1/N2 message transfer request
type N1N2MessageTransferReqData struct {
	N1MessageContainer *N1MessageContainer `json:"n1MessageContainer,omitempty" yaml:"n1MessageContainer" bson:"n1MessageContainer,omitempty"`
	N2InfoContainer    *N2InfoContainer    `json:"n2InfoContainer,omitempty" yaml:"n2InfoContainer" bson:"n2InfoContainer,omitempty"`
	MtData             *RefToBinaryData    `json:"mtData,omitempty" yaml:"mtData" bson:"mtData,omitempty"`
	SkipInd            bool                `json:"skipInd,omitempty" yaml:"skipInd" bson:"skipInd,omitempty"`
	LastMsgIndication  bool                `json:"lastMsgIndication,omitempty" yaml:"lastMsgIndication" bson:"lastMsgIndication,omitempty"`
	// Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.
	PduSessionId int32 `json:"pduSessionId,omitempty" yaml:"pduSessionId" bson:"pduSessionId,omitempty"`
	// LCS Correlation ID.
	LcsCorrelationId string `json:"lcsCorrelationId,omitempty" yaml:"lcsCorrelationId" bson:"lcsCorrelationId,omitempty"`
	// Paging Policy Indicator
	Ppi int32 `json:"ppi,omitempty" yaml:"ppi" bson:"ppi,omitempty"`
	Arp *Arp  `json:"arp,omitempty" yaml:"arp" bson:"arp,omitempty"`
	// Unsigned integer representing a 5G QoS Identifier (see clause 5.7.2.1 of 3GPP TS 23.501, within the range 0 to 255.
	Var5qi int32 `json:"5qi,omitempty" yaml:"5qi" bson:"5qi,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	N1n2FailureTxfNotifURI string          `json:"n1n2FailureTxfNotifURI,omitempty" yaml:"n1n2FailureTxfNotifURI" bson:"n1n2FailureTxfNotifURI,omitempty"`
	SmfReallocationInd     bool            `json:"smfReallocationInd,omitempty" yaml:"smfReallocationInd" bson:"smfReallocationInd,omitempty"`
	AreaOfValidity         *AreaOfValidity `json:"areaOfValidity,omitempty" yaml:"areaOfValidity" bson:"areaOfValidity,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures string     `json:"supportedFeatures,omitempty" yaml:"supportedFeatures" bson:"supportedFeatures,omitempty"`
	OldGuami          *Guami     `json:"oldGuami,omitempty" yaml:"oldGuami" bson:"oldGuami,omitempty"`
	MaAcceptedInd     bool       `json:"maAcceptedInd,omitempty" yaml:"maAcceptedInd" bson:"maAcceptedInd,omitempty"`
	ExtBufSupport     bool       `json:"extBufSupport,omitempty" yaml:"extBufSupport" bson:"extBufSupport,omitempty"`
	TargetAccess      AccessType `json:"targetAccess,omitempty" yaml:"targetAccess" bson:"targetAccess,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	NfId string `json:"nfId,omitempty" yaml:"nfId" bson:"nfId,omitempty"`
}
