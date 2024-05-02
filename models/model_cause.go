/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 29.502 V16.9.0; 5G System; Session Management Services; Stage 3
 * Url: http://www.3gpp.org/ftp/Specs/archive/29_series/29.502/
 *
 * API version: 1.1.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type Cause string

// List of Cause
const (
	Cause_REL_DUE_TO_HO                     Cause = "REL_DUE_TO_HO"
	Cause_EPS_FALLBACK                      Cause = "EPS_FALLBACK"
	Cause_REL_DUE_TO_UP_SEC                 Cause = "REL_DUE_TO_UP_SEC"
	Cause_DNN_CONGESTION                    Cause = "DNN_CONGESTION"
	Cause_S_NSSAI_CONGESTION                Cause = "S_NSSAI_CONGESTION"
	Cause_REL_DUE_TO_REACTIVATION           Cause = "REL_DUE_TO_REACTIVATION"
	Cause__5_G_AN_NOT_RESPONDING            Cause = "5G_AN_NOT_RESPONDING"
	Cause_REL_DUE_TO_SLICE_NOT_AVAILABLE    Cause = "REL_DUE_TO_SLICE_NOT_AVAILABLE"
	Cause_REL_DUE_TO_DUPLICATE_SESSION_ID   Cause = "REL_DUE_TO_DUPLICATE_SESSION_ID"
	Cause_PDU_SESSION_STATUS_MISMATCH       Cause = "PDU_SESSION_STATUS_MISMATCH"
	Cause_HO_FAILURE                        Cause = "HO_FAILURE"
	Cause_INSUFFICIENT_UP_RESOURCES         Cause = "INSUFFICIENT_UP_RESOURCES"
	Cause_PDU_SESSION_HANDED_OVER           Cause = "PDU_SESSION_HANDED_OVER"
	Cause_PDU_SESSION_RESUMED               Cause = "PDU_SESSION_RESUMED"
	Cause_CN_ASSISTED_RAN_PARAMETER_TUNING  Cause = "CN_ASSISTED_RAN_PARAMETER_TUNING"
	Cause_ISMF_CONTEXT_TRANSFER             Cause = "ISMF_CONTEXT_TRANSFER"
	Cause_SMF_CONTEXT_TRANSFER              Cause = "SMF_CONTEXT_TRANSFER"
	Cause_REL_DUE_TO_PS_TO_CS_HO            Cause = "REL_DUE_TO_PS_TO_CS_HO"
	Cause_REL_DUE_TO_SUBSCRIPTION_CHANGE    Cause = "REL_DUE_TO_SUBSCRIPTION_CHANGE"
	Cause_HO_CANCEL                         Cause = "HO_CANCEL"
	Cause_REL_DUE_TO_SLICE_NOT_AUTHORIZED   Cause = "REL_DUE_TO_SLICE_NOT_AUTHORIZED"
	Cause_PDU_SESSION_HAND_OVER_FAILURE     Cause = "PDU_SESSION_HAND_OVER_FAILURE"
	Cause_DDN_FAILURE_STATUS                Cause = "DDN_FAILURE_STATUS"
	Cause_REL_DUE_TO_CP_ONLY_NOT_APPLICABLE Cause = "REL_DUE_TO_CP_ONLY_NOT_APPLICABLE"
	Cause_NOT_SUPPORTED_WITH_ISMF           Cause = "NOT_SUPPORTED_WITH_ISMF"
	Cause_CHANGED_ANCHOR_SMF                Cause = "CHANGED_ANCHOR_SMF"
	Cause_CHANGED_INTERMEDIATE_SMF          Cause = "CHANGED_INTERMEDIATE_SMF"
)
