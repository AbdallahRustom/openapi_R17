/*
 * Npcf_MBSPolicyControl API
 *
 * MBS Policy Control Service   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.537 V17.3.0; 5G System; Multicast/Broadcast Policy Control Services.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.537/
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Represents an MBS Media Component.
type MbsMediaComp struct {
	MbsMedCompNum int32          `json:"mbsMedCompNum" yaml:"mbsMedCompNum" bson:"mbsMedCompNum,omitempty"`
	MbsFlowDescs  []string       `json:"mbsFlowDescs,omitempty" yaml:"mbsFlowDescs" bson:"mbsFlowDescs,omitempty"`
	MbsSdfResPrio ReservPriority `json:"mbsSdfResPrio,omitempty" yaml:"mbsSdfResPrio" bson:"mbsSdfResPrio,omitempty"`
	MbsMediaInfo  *MbsMediaInfo  `json:"mbsMediaInfo,omitempty" yaml:"mbsMediaInfo" bson:"mbsMediaInfo,omitempty"`
	QosRef        string         `json:"qosRef,omitempty" yaml:"qosRef" bson:"qosRef,omitempty"`
	MbsQoSReq     *MbsQoSReq     `json:"mbsQoSReq,omitempty" yaml:"mbsQoSReq" bson:"mbsQoSReq,omitempty"`
}
