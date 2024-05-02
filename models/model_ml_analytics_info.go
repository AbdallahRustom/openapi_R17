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

// ML Analytics Filter information supported by the Nnwdaf_MLModelProvision service
type MlAnalyticsInfo struct {
	MlAnalyticsIds   []NwdafEvent `json:"mlAnalyticsIds,omitempty" yaml:"mlAnalyticsIds" bson:"mlAnalyticsIds,omitempty"`
	SnssaiList       []Snssai     `json:"snssaiList,omitempty" yaml:"snssaiList" bson:"snssaiList,omitempty"`
	TrackingAreaList []Tai        `json:"trackingAreaList,omitempty" yaml:"trackingAreaList" bson:"trackingAreaList,omitempty"`
}
