/*
 * Namf_Communication
 *
 * AMF Communication Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.518 V17.12.0; 5G System; Access and Mobility Management Services
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.518/
 *
 * API version: 1.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type N2InfoNotifyRequest struct {
	JsonData                *N2InformationNotification `json:"jsonData,omitempty" yaml:"jsonData" bson:"jsonData,omitempty"`
	BinaryDataN1Message     []byte                     `json:"binaryDataN1Message,omitempty" yaml:"binaryDataN1Message" bson:"binaryDataN1Message,omitempty"`
	BinaryDataN2Information []byte                     `json:"binaryDataN2Information,omitempty" yaml:"binaryDataN2Information" bson:"binaryDataN2Information,omitempty"`
}
