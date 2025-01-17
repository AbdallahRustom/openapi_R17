/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.502 V17.11.0; 5G System; Session Management Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.502/
 *
 * API version: 1.2.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Data within Update SM Context Request
type SmfPduSessionSmContextUpdateData struct {
	// String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.
	Pei string `json:"pei,omitempty" yaml:"pei" bson:"pei,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	ServingNfId        string          `json:"servingNfId,omitempty" yaml:"servingNfId" bson:"servingNfId,omitempty"`
	Guami              *Guami          `json:"guami,omitempty" yaml:"guami" bson:"guami,omitempty"`
	ServingNetwork     *PlmnIdNid      `json:"servingNetwork,omitempty" yaml:"servingNetwork" bson:"servingNetwork,omitempty"`
	BackupAmfInfo      []BackupAmfInfo `json:"backupAmfInfo,omitempty" yaml:"backupAmfInfo" bson:"backupAmfInfo,omitempty"`
	AnType             AccessType      `json:"anType,omitempty" yaml:"anType" bson:"anType,omitempty"`
	AdditionalAnType   AccessType      `json:"additionalAnType,omitempty" yaml:"additionalAnType" bson:"additionalAnType,omitempty"`
	AnTypeToReactivate AccessType      `json:"anTypeToReactivate,omitempty" yaml:"anTypeToReactivate" bson:"anTypeToReactivate,omitempty"`
	RatType            RatType         `json:"ratType,omitempty" yaml:"ratType" bson:"ratType,omitempty"`
	PresenceInLadn     PresenceState   `json:"presenceInLadn,omitempty" yaml:"presenceInLadn" bson:"presenceInLadn,omitempty"`
	UeLocation         *UserLocation   `json:"ueLocation,omitempty" yaml:"ueLocation" bson:"ueLocation,omitempty"`
	// String with format \"time-numoffset\" optionally appended by \"daylightSavingTime\", where  - \"time-numoffset\" shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \"daylightSavingTime\" shall represent the adjustment that has been made and shall be    encoded as \"+1\" or \"+2\" for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time.
	UeTimeZone         string           `json:"ueTimeZone,omitempty" yaml:"ueTimeZone" bson:"ueTimeZone,omitempty"`
	AddUeLocation      *UserLocation    `json:"addUeLocation,omitempty" yaml:"addUeLocation" bson:"addUeLocation,omitempty"`
	UpCnxState         UpCnxState       `json:"upCnxState,omitempty" yaml:"upCnxState" bson:"upCnxState,omitempty"`
	HoState            HoState          `json:"hoState,omitempty" yaml:"hoState" bson:"hoState,omitempty"`
	ToBeSwitched       bool             `json:"toBeSwitched,omitempty" yaml:"toBeSwitched" bson:"toBeSwitched,omitempty"`
	FailedToBeSwitched bool             `json:"failedToBeSwitched,omitempty" yaml:"failedToBeSwitched" bson:"failedToBeSwitched,omitempty"`
	N1SmMsg            *RefToBinaryData `json:"n1SmMsg,omitempty" yaml:"n1SmMsg" bson:"n1SmMsg,omitempty"`
	N2SmInfo           *RefToBinaryData `json:"n2SmInfo,omitempty" yaml:"n2SmInfo" bson:"n2SmInfo,omitempty"`
	N2SmInfoType       N2SmInfoType     `json:"n2SmInfoType,omitempty" yaml:"n2SmInfoType" bson:"n2SmInfoType,omitempty"`
	TargetId           *NgRanTargetId   `json:"targetId,omitempty" yaml:"targetId" bson:"targetId,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	TargetServingNfId string `json:"targetServingNfId,omitempty" yaml:"targetServingNfId" bson:"targetServingNfId,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	SmContextStatusUri    string                             `json:"smContextStatusUri,omitempty" yaml:"smContextStatusUri" bson:"smContextStatusUri,omitempty"`
	DataForwarding        bool                               `json:"dataForwarding,omitempty" yaml:"dataForwarding" bson:"dataForwarding,omitempty"`
	N9ForwardingTunnel    *TunnelInfo                        `json:"n9ForwardingTunnel,omitempty" yaml:"n9ForwardingTunnel" bson:"n9ForwardingTunnel,omitempty"`
	N9DlForwardingTnlList []IndirectDataForwardingTunnelInfo `json:"n9DlForwardingTnlList,omitempty" yaml:"n9DlForwardingTnlList" bson:"n9DlForwardingTnlList,omitempty"`
	N9UlForwardingTnlList []IndirectDataForwardingTunnelInfo `json:"n9UlForwardingTnlList,omitempty" yaml:"n9UlForwardingTnlList" bson:"n9UlForwardingTnlList,omitempty"`
	N9DlForwardingTunnel  *TunnelInfo                        `json:"n9DlForwardingTunnel,omitempty" yaml:"n9DlForwardingTunnel" bson:"n9DlForwardingTunnel,omitempty"`
	// indicating a time in seconds.
	N9InactivityTimer int32              `json:"n9InactivityTimer,omitempty" yaml:"n9InactivityTimer" bson:"n9InactivityTimer,omitempty"`
	EpsBearerSetup    []string           `json:"epsBearerSetup,omitempty" yaml:"epsBearerSetup" bson:"epsBearerSetup,omitempty"`
	RevokeEbiList     []int32            `json:"revokeEbiList,omitempty" yaml:"revokeEbiList" bson:"revokeEbiList,omitempty"`
	Release           bool               `json:"release,omitempty" yaml:"release" bson:"release,omitempty"`
	Cause             SmfPduSessionCause `json:"cause,omitempty" yaml:"cause" bson:"cause,omitempty"`
	NgApCause         *NgApCause         `json:"ngApCause,omitempty" yaml:"ngApCause" bson:"ngApCause,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Var5gMmCauseValue  int32                     `json:"5gMmCauseValue,omitempty" yaml:"5gMmCauseValue" bson:"5gMmCauseValue,omitempty"`
	SNssai             *Snssai                   `json:"sNssai,omitempty" yaml:"sNssai" bson:"sNssai,omitempty"`
	TraceData          *TraceData                `json:"traceData,omitempty" yaml:"traceData" bson:"traceData,omitempty"`
	EpsInterworkingInd EpsInterworkingIndication `json:"epsInterworkingInd,omitempty" yaml:"epsInterworkingInd" bson:"epsInterworkingInd,omitempty"`
	AnTypeCanBeChanged bool                      `json:"anTypeCanBeChanged,omitempty" yaml:"anTypeCanBeChanged" bson:"anTypeCanBeChanged,omitempty"`
	N2SmInfoExt1       *RefToBinaryData          `json:"n2SmInfoExt1,omitempty" yaml:"n2SmInfoExt1" bson:"n2SmInfoExt1,omitempty"`
	N2SmInfoTypeExt1   N2SmInfoType              `json:"n2SmInfoTypeExt1,omitempty" yaml:"n2SmInfoTypeExt1" bson:"n2SmInfoTypeExt1,omitempty"`
	MaReleaseInd       MaReleaseIndication       `json:"maReleaseInd,omitempty" yaml:"maReleaseInd" bson:"maReleaseInd,omitempty"`
	MaNwUpgradeInd     bool                      `json:"maNwUpgradeInd,omitempty" yaml:"maNwUpgradeInd" bson:"maNwUpgradeInd,omitempty"`
	MaRequestInd       bool                      `json:"maRequestInd,omitempty" yaml:"maRequestInd" bson:"maRequestInd,omitempty"`
	ExemptionInd       *ExemptionInd             `json:"exemptionInd,omitempty" yaml:"exemptionInd" bson:"exemptionInd,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures     string            `json:"supportedFeatures,omitempty" yaml:"supportedFeatures" bson:"supportedFeatures,omitempty"`
	MoExpDataCounter      *MoExpDataCounter `json:"moExpDataCounter,omitempty" yaml:"moExpDataCounter" bson:"moExpDataCounter,omitempty"`
	ExtendedNasSmTimerInd bool              `json:"extendedNasSmTimerInd,omitempty" yaml:"extendedNasSmTimerInd" bson:"extendedNasSmTimerInd,omitempty"`
	// string with format 'bytes' as defined in OpenAPI
	ForwardingFTeid                      string                    `json:"forwardingFTeid,omitempty" yaml:"forwardingFTeid" bson:"forwardingFTeid,omitempty"`
	ForwardingBearerContexts             []string                  `json:"forwardingBearerContexts,omitempty" yaml:"forwardingBearerContexts" bson:"forwardingBearerContexts,omitempty"`
	DdnFailureSubs                       *DdnFailureSubs           `json:"ddnFailureSubs,omitempty" yaml:"ddnFailureSubs" bson:"ddnFailureSubs,omitempty"`
	SkipN2PduSessionResRelInd            bool                      `json:"skipN2PduSessionResRelInd,omitempty" yaml:"skipN2PduSessionResRelInd" bson:"skipN2PduSessionResRelInd,omitempty"`
	SecondaryRatUsageDataReportContainer []string                  `json:"secondaryRatUsageDataReportContainer,omitempty" yaml:"secondaryRatUsageDataReportContainer" bson:"secondaryRatUsageDataReportContainer,omitempty"`
	SmPolicyNotifyInd                    bool                      `json:"smPolicyNotifyInd,omitempty" yaml:"smPolicyNotifyInd" bson:"smPolicyNotifyInd,omitempty"`
	PcfUeCallbackInfo                    *PcfUeCallbackInfo        `json:"pcfUeCallbackInfo,omitempty" yaml:"pcfUeCallbackInfo" bson:"pcfUeCallbackInfo,omitempty"`
	SatelliteBackhaulCat                 SatelliteBackhaulCategory `json:"satelliteBackhaulCat,omitempty" yaml:"satelliteBackhaulCat" bson:"satelliteBackhaulCat,omitempty"`
}
