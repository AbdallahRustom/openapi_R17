/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.502 V17.11.0; 5G System; Session Management Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.502/
 *
 * API version: 1.2.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type TransferMtDataismfRequest struct {
	JsonData     *TransferMtDataReqData `json:"jsonData,omitempty" yaml:"jsonData" bson:"jsonData,omitempty"`
	BinaryMtData []byte                 `json:"binaryMtData,omitempty" yaml:"binaryMtData" bson:"binaryMtData,omitempty"`
}
