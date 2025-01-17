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

type FailureCause string

// List of FailureCause
const (
	FailureCause_PCC_RULE_EVENT       FailureCause = "PCC_RULE_EVENT"
	FailureCause_PCC_QOS_FLOW_EVENT   FailureCause = "PCC_QOS_FLOW_EVENT"
	FailureCause_RULE_PERMANENT_ERROR FailureCause = "RULE_PERMANENT_ERROR"
	FailureCause_RULE_TEMPORARY_ERROR FailureCause = "RULE_TEMPORARY_ERROR"
	FailureCause_POL_DEC_ERROR        FailureCause = "POL_DEC_ERROR"
)
