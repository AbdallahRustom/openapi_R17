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

type SmsSupport string

// List of SmsSupport
const (
	SmsSupport__3_GPP    SmsSupport = "3GPP"
	SmsSupport_NON_3_GPP SmsSupport = "NON_3GPP"
	SmsSupport_BOTH      SmsSupport = "BOTH"
	SmsSupport_NONE      SmsSupport = "NONE"
)