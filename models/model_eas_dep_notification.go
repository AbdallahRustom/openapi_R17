/*
 * Nnef_EASDeployment
 *
 * NEF EAS Deployment service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.591 V17.6.0; 5G System; Network Exposure Function Southbound Services; Stage 3.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.591/
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Represents the EAS Deployment Notifcation.
type EasDepNotification struct {
	EasDepInfo *EasDeployInfoData `json:"easDepInfo" yaml:"easDepInfo" bson:"easDepInfo,omitempty"`
	EventId    EasEvent           `json:"eventId" yaml:"eventId" bson:"eventId,omitempty"`
}
