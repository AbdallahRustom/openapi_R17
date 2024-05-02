/*
 * Npcf_PolicyAuthorization Service API
 *
 * PCF Policy Authorization Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 29.514 V16.10.0; 5G System; Policy Authorization Service;Stage 3.
 * Url: http://www.3gpp.org/ftp/Specs/archive/29_series/29.514/
 *
 * API version: 1.1.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

// describes the notification of a matched event
type EventsNotification struct {
	AccessType                AccessType                             `json:"accessType,omitempty" yaml:"accessType" bson:"accessType,omitempty"`
	AddAccessInfo             *AdditionalAccessInfo                  `json:"addAccessInfo,omitempty" yaml:"addAccessInfo" bson:"addAccessInfo,omitempty"`
	RelAccessInfo             *AdditionalAccessInfo                  `json:"relAccessInfo,omitempty" yaml:"relAccessInfo" bson:"relAccessInfo,omitempty"`
	AnChargAddr               *AccNetChargingAddress                 `json:"anChargAddr,omitempty" yaml:"anChargAddr" bson:"anChargAddr,omitempty"`
	AnChargIds                []AccessNetChargingIdentifier          `json:"anChargIds,omitempty" yaml:"anChargIds" bson:"anChargIds,omitempty"`
	AnGwAddr                  *AnGwAddress                           `json:"anGwAddr,omitempty" yaml:"anGwAddr" bson:"anGwAddr,omitempty"`
	EvSubsUri                 string                                 `json:"evSubsUri" yaml:"evSubsUri" bson:"evSubsUri,omitempty"`
	EvNotifs                  []AfEventNotification                  `json:"evNotifs" yaml:"evNotifs" bson:"evNotifs,omitempty"`
	FailedResourcAllocReports []ResourcesAllocationInfo              `json:"failedResourcAllocReports,omitempty" yaml:"failedResourcAllocReports" bson:"failedResourcAllocReports,omitempty"`
	SuccResourcAllocReports   []ResourcesAllocationInfo              `json:"succResourcAllocReports,omitempty" yaml:"succResourcAllocReports" bson:"succResourcAllocReports,omitempty"`
	NoNetLocSupp              NetLocAccessSupport                    `json:"noNetLocSupp,omitempty" yaml:"noNetLocSupp" bson:"noNetLocSupp,omitempty"`
	OutOfCredReports          []OutOfCreditInformation               `json:"outOfCredReports,omitempty" yaml:"outOfCredReports" bson:"outOfCredReports,omitempty"`
	PlmnId                    *PlmnIdNid                             `json:"plmnId,omitempty" yaml:"plmnId" bson:"plmnId,omitempty"`
	QncReports                []PolicyAuthQosNotificationControlInfo `json:"qncReports,omitempty" yaml:"qncReports" bson:"qncReports,omitempty"`
	QosMonReports             []QosMonitoringReport                  `json:"qosMonReports,omitempty" yaml:"qosMonReports" bson:"qosMonReports,omitempty"`
	// Contains the RAN and/or NAS release cause.
	RanNasRelCauses     []RanNasRelCause           `json:"ranNasRelCauses,omitempty" yaml:"ranNasRelCauses" bson:"ranNasRelCauses,omitempty"`
	RatType             RatType                    `json:"ratType,omitempty" yaml:"ratType" bson:"ratType,omitempty"`
	UeLoc               *UserLocation              `json:"ueLoc,omitempty" yaml:"ueLoc" bson:"ueLoc,omitempty"`
	UeLocTime           *time.Time                 `json:"ueLocTime,omitempty" yaml:"ueLocTime" bson:"ueLocTime,omitempty"`
	UeTimeZone          string                     `json:"ueTimeZone,omitempty" yaml:"ueTimeZone" bson:"ueTimeZone,omitempty"`
	UsgRep              *AccumulatedUsage          `json:"usgRep,omitempty" yaml:"usgRep" bson:"usgRep,omitempty"`
	TsnBridgeManCont    *BridgeManagementContainer `json:"tsnBridgeManCont,omitempty" yaml:"tsnBridgeManCont" bson:"tsnBridgeManCont,omitempty"`
	TsnPortManContDstt  *PortManagementContainer   `json:"tsnPortManContDstt,omitempty" yaml:"tsnPortManContDstt" bson:"tsnPortManContDstt,omitempty"`
	TsnPortManContNwtts []PortManagementContainer  `json:"tsnPortManContNwtts,omitempty" yaml:"tsnPortManContNwtts" bson:"tsnPortManContNwtts,omitempty"`
}