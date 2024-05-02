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

// Contains the background data transfer data.
type DrBdtData struct {
	AspId       string          `json:"aspId" yaml:"aspId" bson:"aspId,omitempty"`
	TransPolicy *TransferPolicy `json:"transPolicy" yaml:"transPolicy" bson:"transPolicy,omitempty"`
	// string identifying a BDT Reference ID as defined in subclause 5.3.3 of 3GPP TS 29.154.
	BdtRefId   string           `json:"bdtRefId,omitempty" yaml:"bdtRefId" bson:"bdtRefId,omitempty"`
	NwAreaInfo *NetworkAreaInfo `json:"nwAreaInfo,omitempty" yaml:"nwAreaInfo" bson:"nwAreaInfo,omitempty"`
	NumOfUes   int32            `json:"numOfUes,omitempty" yaml:"numOfUes" bson:"numOfUes,omitempty"`
	VolPerUe   *UsageThreshold  `json:"volPerUe,omitempty" yaml:"volPerUe" bson:"volPerUe,omitempty"`
	Dnn        string           `json:"dnn,omitempty" yaml:"dnn" bson:"dnn,omitempty"`
	Snssai     *Snssai          `json:"snssai,omitempty" yaml:"snssai" bson:"snssai,omitempty"`
	// Identify a traffic descriptor as defined in Figure 5.2.2 of 3GPP TS 24.526, octets v+5 to w.
	TrafficDes string          `json:"trafficDes,omitempty" yaml:"trafficDes" bson:"trafficDes,omitempty"`
	BdtpStatus BdtPolicyStatus `json:"bdtpStatus,omitempty" yaml:"bdtpStatus" bson:"bdtpStatus,omitempty"`
	SuppFeat   string          `json:"suppFeat,omitempty" yaml:"suppFeat" bson:"suppFeat,omitempty"`
}
