/*
 * Npcf_UEPolicyControl
 *
 * UE Policy Control Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 29.525 V16.9.0; 5G System; UE Policy Control Service.
 * Url: http://www.3gpp.org/ftp/Specs/archive/29_series/29.525/
 *
 * API version: 1.1.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type UePolicyAssociationRequest struct {
	NotificationUri string `json:"notificationUri" yaml:"notificationUri" bson:"notificationUri,omitempty"`
	// Alternate or backup IPv4 Address(es) where to send Notifications.
	AltNotifIpv4Addrs []string `json:"altNotifIpv4Addrs,omitempty" yaml:"altNotifIpv4Addrs" bson:"altNotifIpv4Addrs,omitempty"`
	// Alternate or backup IPv6 Address(es) where to send Notifications.
	AltNotifIpv6Addrs []string `json:"altNotifIpv6Addrs,omitempty" yaml:"altNotifIpv6Addrs" bson:"altNotifIpv6Addrs,omitempty"`
	// Alternate or backup FQDN(s) where to send Notifications.
	AltNotifFqdns []string      `json:"altNotifFqdns,omitempty" yaml:"altNotifFqdns" bson:"altNotifFqdns,omitempty"`
	Supi          string        `json:"supi" yaml:"supi" bson:"supi,omitempty"`
	Gpsi          string        `json:"gpsi,omitempty" yaml:"gpsi" bson:"gpsi,omitempty"`
	AccessType    AccessType    `json:"accessType,omitempty" yaml:"accessType" bson:"accessType,omitempty"`
	Pei           string        `json:"pei,omitempty" yaml:"pei" bson:"pei,omitempty"`
	UserLoc       *UserLocation `json:"userLoc,omitempty" yaml:"userLoc" bson:"userLoc,omitempty"`
	TimeZone      string        `json:"timeZone,omitempty" yaml:"timeZone" bson:"timeZone,omitempty"`
	ServingPlmn   *PlmnIdNid    `json:"servingPlmn,omitempty" yaml:"servingPlmn" bson:"servingPlmn,omitempty"`
	RatType       RatType       `json:"ratType,omitempty" yaml:"ratType" bson:"ratType,omitempty"`
	GroupIds      []string      `json:"groupIds,omitempty" yaml:"groupIds" bson:"groupIds,omitempty"`
	HPcfId        string        `json:"hPcfId,omitempty" yaml:"hPcfId" bson:"hPcfId,omitempty"`
	UePolReq      string        `json:"uePolReq,omitempty" yaml:"uePolReq" bson:"uePolReq,omitempty"`
	Guami         *Guami        `json:"guami,omitempty" yaml:"guami" bson:"guami,omitempty"`
	ServiceName   ServiceName   `json:"serviceName,omitempty" yaml:"serviceName" bson:"serviceName,omitempty"`
	ServingNfId   string        `json:"servingNfId,omitempty" yaml:"servingNfId" bson:"servingNfId,omitempty"`
	Pc5Capab      Pc5Capability `json:"pc5Capab,omitempty" yaml:"pc5Capab" bson:"pc5Capab,omitempty"`
	SuppFeat      string        `json:"suppFeat" yaml:"suppFeat" bson:"suppFeat,omitempty"`
}
