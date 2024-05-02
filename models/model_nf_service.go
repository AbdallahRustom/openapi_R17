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

import (
	"time"
)

// Information of a given NF Service Instance; it is part of the NFProfile of an NF Instance
type NfService struct {
	ServiceInstanceId string             `json:"serviceInstanceId" yaml:"serviceInstanceId" bson:"serviceInstanceId,omitempty"`
	ServiceName       ServiceName        `json:"serviceName" yaml:"serviceName" bson:"serviceName,omitempty"`
	Versions          []NfServiceVersion `json:"versions" yaml:"versions" bson:"versions,omitempty"`
	Scheme            UriScheme          `json:"scheme" yaml:"scheme" bson:"scheme,omitempty"`
	NfServiceStatus   NfServiceStatus    `json:"nfServiceStatus" yaml:"nfServiceStatus" bson:"nfServiceStatus,omitempty"`
	// Fully Qualified Domain Name
	Fqdn string `json:"fqdn,omitempty" yaml:"fqdn" bson:"fqdn,omitempty"`
	// Fully Qualified Domain Name
	InterPlmnFqdn                    string                            `json:"interPlmnFqdn,omitempty" yaml:"interPlmnFqdn" bson:"interPlmnFqdn,omitempty"`
	IpEndPoints                      []IpEndPoint                      `json:"ipEndPoints,omitempty" yaml:"ipEndPoints" bson:"ipEndPoints,omitempty"`
	ApiPrefix                        string                            `json:"apiPrefix,omitempty" yaml:"apiPrefix" bson:"apiPrefix,omitempty"`
	DefaultNotificationSubscriptions []DefaultNotificationSubscription `json:"defaultNotificationSubscriptions,omitempty" yaml:"defaultNotificationSubscriptions" bson:"defaultNotificationSubscriptions,omitempty"`
	AllowedPlmns                     []PlmnId                          `json:"allowedPlmns,omitempty" yaml:"allowedPlmns" bson:"allowedPlmns,omitempty"`
	AllowedSnpns                     []PlmnIdNid                       `json:"allowedSnpns,omitempty" yaml:"allowedSnpns" bson:"allowedSnpns,omitempty"`
	AllowedNfTypes                   []NfType                          `json:"allowedNfTypes,omitempty" yaml:"allowedNfTypes" bson:"allowedNfTypes,omitempty"`
	AllowedNfDomains                 []string                          `json:"allowedNfDomains,omitempty" yaml:"allowedNfDomains" bson:"allowedNfDomains,omitempty"`
	AllowedNssais                    []ExtSnssai                       `json:"allowedNssais,omitempty" yaml:"allowedNssais" bson:"allowedNssais,omitempty"`
	AllowedOperationsPerNfType       map[string][]string               `json:"allowedOperationsPerNfType,omitempty" yaml:"allowedOperationsPerNfType" bson:"allowedOperationsPerNfType,omitempty"`
	AllowedOperationsPerNfInstance   map[string][]string               `json:"allowedOperationsPerNfInstance,omitempty" yaml:"allowedOperationsPerNfInstance" bson:"allowedOperationsPerNfInstance,omitempty"`
	Priority                         int32                             `json:"priority,omitempty" yaml:"priority" bson:"priority,omitempty"`
	Capacity                         int32                             `json:"capacity,omitempty" yaml:"capacity" bson:"capacity,omitempty"`
	Load                             int32                             `json:"load,omitempty" yaml:"load" bson:"load,omitempty"`
	LoadTimeStamp                    *time.Time                        `json:"loadTimeStamp,omitempty" yaml:"loadTimeStamp" bson:"loadTimeStamp,omitempty"`
	RecoveryTime                     *time.Time                        `json:"recoveryTime,omitempty" yaml:"recoveryTime" bson:"recoveryTime,omitempty"`
	SupportedFeatures                string                            `json:"supportedFeatures,omitempty" yaml:"supportedFeatures" bson:"supportedFeatures,omitempty"`
	NfServiceSetIdList               []string                          `json:"nfServiceSetIdList,omitempty" yaml:"nfServiceSetIdList" bson:"nfServiceSetIdList,omitempty"`
	SNssais                          []ExtSnssai                       `json:"sNssais,omitempty" yaml:"sNssais" bson:"sNssais,omitempty"`
	PerPlmnSnssaiList                []PlmnSnssai                      `json:"perPlmnSnssaiList,omitempty" yaml:"perPlmnSnssaiList" bson:"perPlmnSnssaiList,omitempty"`
	// Vendor ID of the NF Service instance (Private Enterprise Number assigned by IANA)
	VendorId                        string                             `json:"vendorId,omitempty" yaml:"vendorId" bson:"vendorId,omitempty"`
	SupportedVendorSpecificFeatures map[string][]VendorSpecificFeature `json:"supportedVendorSpecificFeatures,omitempty" yaml:"supportedVendorSpecificFeatures" bson:"supportedVendorSpecificFeatures,omitempty"`
	Oauth2Required                  bool                               `json:"oauth2Required,omitempty" yaml:"oauth2Required" bson:"oauth2Required,omitempty"`
}