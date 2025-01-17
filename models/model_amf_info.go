/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 29.510 V16.9.0; 5G System; Network Function Repository Services; Stage 3
 * Url: http://www.3gpp.org/ftp/Specs/archive/29_series/29.510/
 *
 * API version: 1.1.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Information of an AMF NF Instance
type AmfInfo struct {
	AmfSetId             string              `json:"amfSetId" yaml:"amfSetId" bson:"amfSetId,omitempty"`
	AmfRegionId          string              `json:"amfRegionId" yaml:"amfRegionId" bson:"amfRegionId,omitempty"`
	GuamiList            []Guami             `json:"guamiList" yaml:"guamiList" bson:"guamiList,omitempty"`
	TaiList              []Tai               `json:"taiList,omitempty" yaml:"taiList" bson:"taiList,omitempty"`
	TaiRangeList         []TaiRange          `json:"taiRangeList,omitempty" yaml:"taiRangeList" bson:"taiRangeList,omitempty"`
	BackupInfoAmfFailure []Guami             `json:"backupInfoAmfFailure,omitempty" yaml:"backupInfoAmfFailure" bson:"backupInfoAmfFailure,omitempty"`
	BackupInfoAmfRemoval []Guami             `json:"backupInfoAmfRemoval,omitempty" yaml:"backupInfoAmfRemoval" bson:"backupInfoAmfRemoval,omitempty"`
	N2InterfaceAmfInfo   *N2InterfaceAmfInfo `json:"n2InterfaceAmfInfo,omitempty" yaml:"n2InterfaceAmfInfo" bson:"n2InterfaceAmfInfo,omitempty"`
}
