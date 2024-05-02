/*
 * Nudm_UEAU
 *
 * UDM UE Authentication Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.503 Unified Data Management Services, version 17.10.0
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.503/
 *
 * API version: 1.2.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type NodeType string

// List of NodeType
const (
	NodeType_AUSF            NodeType = "AUSF"
	NodeType_VLR             NodeType = "VLR"
	NodeType_SGSN            NodeType = "SGSN"
	NodeType_S_CSCF          NodeType = "S_CSCF"
	NodeType_BSF             NodeType = "BSF"
	NodeType_GAN_AAA_SERVER  NodeType = "GAN_AAA_SERVER"
	NodeType_WLAN_AAA_SERVER NodeType = "WLAN_AAA_SERVER"
	NodeType_MME             NodeType = "MME"
)
