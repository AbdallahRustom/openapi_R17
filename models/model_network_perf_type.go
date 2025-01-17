/*
 * Nnwdaf_MLModelProvision
 *
 * Nnwdaf_MLModelProvision API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.520 V17.7.0; 5G System; Network Data Analytics Services.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.520/
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type NetworkPerfType string

// List of NetworkPerfType
const (
	NetworkPerfType_GNB_ACTIVE_RATIO    NetworkPerfType = "GNB_ACTIVE_RATIO"
	NetworkPerfType_GNB_COMPUTING_USAGE NetworkPerfType = "GNB_COMPUTING_USAGE"
	NetworkPerfType_GNB_MEMORY_USAGE    NetworkPerfType = "GNB_MEMORY_USAGE"
	NetworkPerfType_GNB_DISK_USAGE      NetworkPerfType = "GNB_DISK_USAGE"
	NetworkPerfType_NUM_OF_UE           NetworkPerfType = "NUM_OF_UE"
	NetworkPerfType_SESS_SUCC_RATIO     NetworkPerfType = "SESS_SUCC_RATIO"
	NetworkPerfType_HO_SUCC_RATIO       NetworkPerfType = "HO_SUCC_RATIO"
)
