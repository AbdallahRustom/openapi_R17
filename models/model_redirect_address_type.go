/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 29.512 V16.10.0; 5G System; Session Management Policy Control Service.
 * Url: http://www.3gpp.org/ftp/Specs/archive/29_series/29.512/
 *
 * API version: 1.1.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type RedirectAddressType string

// List of RedirectAddressType
const (
	RedirectAddressType_IPV4_ADDR RedirectAddressType = "IPV4_ADDR"
	RedirectAddressType_IPV6_ADDR RedirectAddressType = "IPV6_ADDR"
	RedirectAddressType_URL       RedirectAddressType = "URL"
	RedirectAddressType_SIP_URI   RedirectAddressType = "SIP_URI"
)