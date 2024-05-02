/*
 * 3gpp-pfd-management
 *
 * API for PFD management. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 29.122 V16.9.0 T8 reference point for Northbound APIs
 * Url: http://www.3gpp.org/ftp/Specs/archive/29_series/29.122/
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type PfdPfdData struct {
	// Each element uniquely external application identifier
	ExternalAppId string `json:"externalAppId" yaml:"externalAppId" bson:"externalAppId"`
	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	Self string `json:"self,omitempty" yaml:"self" bson:"self"`
	// Contains the PFDs of the external application identifier. Each PFD is identified in the map via a key containing the PFD identifier.
	Pfds map[string]Pfd `json:"pfds" yaml:"pfds" bson:"pfds"`
	// Unsigned integer identifying a period of time in units of seconds with \"nullable=true\" property.
	AllowedDelay int32 `json:"allowedDelay,omitempty" yaml:"allowedDelay" bson:"allowedDelay"`
	// Unsigned integer identifying a period of time in units of seconds with \"readOnly=true\" property.
	CachingTime int32 `json:"cachingTime,omitempty" yaml:"cachingTime" bson:"cachingTime"`
}
