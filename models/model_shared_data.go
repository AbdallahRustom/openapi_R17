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

type SharedData struct {
	SharedDataId            string                             `json:"sharedDataId" yaml:"sharedDataId" bson:"sharedDataId,omitempty"`
	SharedAmData            *AccessAndMobilitySubscriptionData `json:"sharedAmData,omitempty" yaml:"sharedAmData" bson:"sharedAmData,omitempty"`
	SharedSmsSubsData       *SmsSubscriptionData               `json:"sharedSmsSubsData,omitempty" yaml:"sharedSmsSubsData" bson:"sharedSmsSubsData,omitempty"`
	SharedSmsMngSubsData    *SmsManagementSubscriptionData     `json:"sharedSmsMngSubsData,omitempty" yaml:"sharedSmsMngSubsData" bson:"sharedSmsMngSubsData,omitempty"`
	SharedDnnConfigurations map[string]DnnConfiguration        `json:"sharedDnnConfigurations,omitempty" yaml:"sharedDnnConfigurations" bson:"sharedDnnConfigurations,omitempty"`
	SharedTraceData         *TraceData                         `json:"sharedTraceData,omitempty" yaml:"sharedTraceData" bson:"sharedTraceData,omitempty"`
	SharedSnssaiInfos       map[string]SnssaiInfo              `json:"sharedSnssaiInfos,omitempty" yaml:"sharedSnssaiInfos" bson:"sharedSnssaiInfos,omitempty"`
	SharedVnGroupDatas      map[string]VnGroupData             `json:"sharedVnGroupDatas,omitempty" yaml:"sharedVnGroupDatas" bson:"sharedVnGroupDatas,omitempty"`
}
