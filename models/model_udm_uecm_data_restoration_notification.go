/*
 * Nudm_UECM
 *
 * Nudm Context Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.503 Unified Data Management Services, version 17.13.0
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.503/
 *
 * API version: 1.2.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

// Contains identities representing those UEs potentially affected by a data-loss event at the UDR
type UdmUecmDataRestorationNotification struct {
	// string with format 'date-time' as defined in OpenAPI.
	LastReplicationTime *time.Time `json:"lastReplicationTime,omitempty" yaml:"lastReplicationTime" bson:"lastReplicationTime,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	RecoveryTime *time.Time      `json:"recoveryTime,omitempty" yaml:"recoveryTime" bson:"recoveryTime,omitempty"`
	PlmnId       *PlmnId         `json:"plmnId,omitempty" yaml:"plmnId" bson:"plmnId,omitempty"`
	SupiRanges   []SupiRange     `json:"supiRanges,omitempty" yaml:"supiRanges" bson:"supiRanges,omitempty"`
	GpsiRanges   []IdentityRange `json:"gpsiRanges,omitempty" yaml:"gpsiRanges" bson:"gpsiRanges,omitempty"`
	ResetIds     []string        `json:"resetIds,omitempty" yaml:"resetIds" bson:"resetIds,omitempty"`
	SNssaiList   []Snssai        `json:"sNssaiList,omitempty" yaml:"sNssaiList" bson:"sNssaiList,omitempty"`
	DnnList      []string        `json:"dnnList,omitempty" yaml:"dnnList" bson:"dnnList,omitempty"`
	// Identifier of a group of NFs.
	UdmGroupId string `json:"udmGroupId,omitempty" yaml:"udmGroupId" bson:"udmGroupId,omitempty"`
}
