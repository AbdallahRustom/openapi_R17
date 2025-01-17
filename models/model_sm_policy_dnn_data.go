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

// Contains the SM policy data for a given DNN (and S-NSSAI).
type SmPolicyDnnData struct {
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn             string   `json:"dnn" yaml:"dnn" bson:"dnn,omitempty"`
	AllowedServices []string `json:"allowedServices,omitempty" yaml:"allowedServices" bson:"allowedServices,omitempty"`
	SubscCats       []string `json:"subscCats,omitempty" yaml:"subscCats" bson:"subscCats,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	GbrUl string `json:"gbrUl,omitempty" yaml:"gbrUl" bson:"gbrUl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	GbrDl               string `json:"gbrDl,omitempty" yaml:"gbrDl" bson:"gbrDl,omitempty"`
	AdcSupport          bool   `json:"adcSupport,omitempty" yaml:"adcSupport" bson:"adcSupport,omitempty"`
	SubscSpendingLimits bool   `json:"subscSpendingLimits,omitempty" yaml:"subscSpendingLimits" bson:"subscSpendingLimits,omitempty"`
	// Represents information that identifies which IP pool or external server is used to allocate the IP address.
	Ipv4Index int32 `json:"ipv4Index,omitempty" yaml:"ipv4Index" bson:"ipv4Index,omitempty"`
	// Represents information that identifies which IP pool or external server is used to allocate the IP address.
	Ipv6Index int32                `json:"ipv6Index,omitempty" yaml:"ipv6Index" bson:"ipv6Index,omitempty"`
	Offline   bool                 `json:"offline,omitempty" yaml:"offline" bson:"offline,omitempty"`
	Online    bool                 `json:"online,omitempty" yaml:"online" bson:"online,omitempty"`
	ChfInfo   *ChargingInformation `json:"chfInfo,omitempty" yaml:"chfInfo" bson:"chfInfo,omitempty"`
	// A reference to the UsageMonDataLimit or UsageMonData instancesfor this DNN and SNSSAI that may also include the related monitoring key(s). The key of the map is the limit identifier.
	RefUmDataLimitIds map[string]*LimitIdToMonitoringKey `json:"refUmDataLimitIds,omitempty" yaml:"refUmDataLimitIds" bson:"refUmDataLimitIds,omitempty"`
	MpsPriority       bool                               `json:"mpsPriority,omitempty" yaml:"mpsPriority" bson:"mpsPriority,omitempty"`
	McsPriority       bool                               `json:"mcsPriority,omitempty" yaml:"mcsPriority" bson:"mcsPriority,omitempty"`
	ImsSignallingPrio bool                               `json:"imsSignallingPrio,omitempty" yaml:"imsSignallingPrio" bson:"imsSignallingPrio,omitempty"`
	MpsPriorityLevel  int32                              `json:"mpsPriorityLevel,omitempty" yaml:"mpsPriorityLevel" bson:"mpsPriorityLevel,omitempty"`
	McsPriorityLevel  int32                              `json:"mcsPriorityLevel,omitempty" yaml:"mcsPriorityLevel" bson:"mcsPriorityLevel,omitempty"`
	// Contains Presence reporting area information. The praId attribute within the PresenceInfo data type is the key of the map.
	PraInfos map[string]PresenceInfo `json:"praInfos,omitempty" yaml:"praInfos" bson:"praInfos,omitempty"`
	// Identifies transfer policies of background data transfer. Any string value can be used as a key of the map.
	BdtRefIds         map[string]*string `json:"bdtRefIds,omitempty" yaml:"bdtRefIds" bson:"bdtRefIds,omitempty"`
	LocRoutNotAllowed bool               `json:"locRoutNotAllowed,omitempty" yaml:"locRoutNotAllowed" bson:"locRoutNotAllowed,omitempty"`
}
