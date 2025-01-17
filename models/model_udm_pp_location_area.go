/*
 * Nnwdaf_MLModelProvision
 *
 * Nnwdaf_MLModelProvision API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.520 V17.7.0; 5G System; Network Data Analytics Services.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.520/
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type UdmPpLocationArea struct {
	// Identifies a list of geographic area of the user where the UE is located.
	GeographicAreas []GeographicArea `json:"geographicAreas,omitempty" yaml:"geographicAreas" bson:"geographicAreas,omitempty"`
	// Identifies a list of civic addresses of the user where the UE is located.
	CivicAddresses []CivicAddress   `json:"civicAddresses,omitempty" yaml:"civicAddresses" bson:"civicAddresses,omitempty"`
	NwAreaInfo     *NetworkAreaInfo `json:"nwAreaInfo,omitempty" yaml:"nwAreaInfo" bson:"nwAreaInfo,omitempty"`
	UmtTime        *UmtTime         `json:"umtTime,omitempty" yaml:"umtTime" bson:"umtTime,omitempty"`
}
