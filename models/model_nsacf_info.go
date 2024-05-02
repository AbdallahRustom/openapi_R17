/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.520 V17.9.0; 5G System; Network Data Analytics Services.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.520/
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Information of a NSACF NF Instance
type NsacfInfo struct {
	NsacfCapability *NsacfCapability `json:"nsacfCapability" yaml:"nsacfCapability" bson:"nsacfCapability,omitempty"`
	TaiList         []Tai            `json:"taiList,omitempty" yaml:"taiList" bson:"taiList,omitempty"`
	TaiRangeList    []TaiRange       `json:"taiRangeList,omitempty" yaml:"taiRangeList" bson:"taiRangeList,omitempty"`
	NsacSaiList     []string         `json:"nsacSaiList,omitempty" yaml:"nsacSaiList" bson:"nsacSaiList,omitempty"`
}
