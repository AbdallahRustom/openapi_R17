/*
 * Npcf_AMPolicyControl
 *
 * Access and Mobility Policy Control Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 29.507 V16.9.0; 5G System; Access and Mobility Policy Control Service.
 * Url: http://www.3gpp.org/ftp/Specs/archive/29_series/29.507/
 *
 * API version: 1.1.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type AmPolicyAssociationUpdateRequest struct {
	NotificationUri string `json:"notificationUri,omitempty" yaml:"notificationUri" bson:"notificationUri,omitempty"`
	// Alternate or backup IPv4 Address(es) where to send Notifications.
	AltNotifIpv4Addrs []string `json:"altNotifIpv4Addrs,omitempty" yaml:"altNotifIpv4Addrs" bson:"altNotifIpv4Addrs,omitempty"`
	// Alternate or backup IPv6 Address(es) where to send Notifications.
	AltNotifIpv6Addrs []string `json:"altNotifIpv6Addrs,omitempty" yaml:"altNotifIpv6Addrs" bson:"altNotifIpv6Addrs,omitempty"`
	// Alternate or backup FQDN(s) where to send Notifications.
	AltNotifFqdns []string `json:"altNotifFqdns,omitempty" yaml:"altNotifFqdns" bson:"altNotifFqdns,omitempty"`
	// Request Triggers that the NF service consumer observes.
	Triggers      []PcfAmPolicyControlRequestTrigger `json:"triggers,omitempty" yaml:"triggers" bson:"triggers,omitempty"`
	ServAreaRes   *ServiceAreaRestriction            `json:"servAreaRes,omitempty" yaml:"servAreaRes" bson:"servAreaRes,omitempty"`
	WlServAreaRes *WirelineServiceAreaRestriction    `json:"wlServAreaRes,omitempty" yaml:"wlServAreaRes" bson:"wlServAreaRes,omitempty"`
	Rfsp          int32                              `json:"rfsp,omitempty" yaml:"rfsp" bson:"rfsp,omitempty"`
	SmfSelInfo    *SmfSelectionData                  `json:"smfSelInfo,omitempty" yaml:"smfSelInfo" bson:"smfSelInfo,omitempty"`
	UeAmbr        *Ambr                              `json:"ueAmbr,omitempty" yaml:"ueAmbr" bson:"ueAmbr,omitempty"`
	// Map of PRA status information.
	PraStatuses map[string]PresenceInfo `json:"praStatuses,omitempty" yaml:"praStatuses" bson:"praStatuses,omitempty"`
	UserLoc     *UserLocation           `json:"userLoc,omitempty" yaml:"userLoc" bson:"userLoc,omitempty"`
	// array of allowed S-NSSAIs for the 3GPP access.
	AllowedSnssais []Snssai `json:"allowedSnssais,omitempty" yaml:"allowedSnssais" bson:"allowedSnssais,omitempty"`
	// mapping of each S-NSSAI of the Allowed NSSAI to the corresponding S-NSSAI of the HPLMN.
	MappingSnssais []MappingOfSnssai `json:"mappingSnssais,omitempty" yaml:"mappingSnssais" bson:"mappingSnssais,omitempty"`
	AccessTypes    []AccessType      `json:"accessTypes,omitempty" yaml:"accessTypes" bson:"accessTypes,omitempty"`
	RatTypes       []RatType         `json:"ratTypes,omitempty" yaml:"ratTypes" bson:"ratTypes,omitempty"`
	// array of allowed S-NSSAIs for the Non-3GPP access.
	N3gAllowedSnssais []Snssai   `json:"n3gAllowedSnssais,omitempty" yaml:"n3gAllowedSnssais" bson:"n3gAllowedSnssais,omitempty"`
	TraceReq          *TraceData `json:"traceReq,omitempty" yaml:"traceReq" bson:"traceReq,omitempty"`
	Guami             *Guami     `json:"guami,omitempty" yaml:"guami" bson:"guami,omitempty"`
}
