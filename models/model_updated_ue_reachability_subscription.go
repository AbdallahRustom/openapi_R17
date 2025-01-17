/*
 * Nhss_imsSDM
 *
 * Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * Source file: 3GPP TS 29.562 HSS Services for interworking with IMS, version 16.8.0
 * Url: http://www.3gpp.org/ftp/Specs/archive/29_series/29.562/
 *
 * API version: 1.0.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models
import (
	"time"
)



type UpdatedUeReachabilitySubscription struct {
	Expiry *time.Time `json:"expiry" yaml:"expiry" bson:"expiry"`
}
