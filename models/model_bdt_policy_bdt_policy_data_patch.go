/*
 * Npcf_BDTPolicyControl Service API
 *
 * PCF BDT Policy Control Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 29.554 V16.7.0; 5G System; Background Data Transfer Policy Control Service.
 * Url: http://www.3gpp.org/ftp/Specs/archive/29_series/29.554/
 *
 * API version: 1.1.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// A JSON Merge Patch body schema containing modification instruction to be performed on the bdtPolData attribute of the BdtPolicy data structure to select a transfer policy. Adds selTransPolicyId to BdtPolicyData data structure.
type BdtPolicyBdtPolicyDataPatch struct {
	// Contains an identity (i.e. transPolicyId value) of the selected transfer policy. If the BdtNotification_5G feature is supported value 0 indicates that no transfer policy is selected.
	SelTransPolicyId int32 `json:"selTransPolicyId" yaml:"selTransPolicyId" bson:"selTransPolicyId,omitempty"`
}
