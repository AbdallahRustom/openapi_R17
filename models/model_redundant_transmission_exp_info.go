/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.520 V17.10.0; 5G System; Network Data Analytics Services.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.520/
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// The redundant transmission experience related information. When subscribed event is  \"RED_TRANS_EXP\", the \"redTransInfos\" attribute shall be included.
type RedundantTransmissionExpInfo struct {
	SpatialValidCon *NetworkAreaInfo `json:"spatialValidCon,omitempty" yaml:"spatialValidCon" bson:"spatialValidCon,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn          string                          `json:"dnn,omitempty" yaml:"dnn" bson:"dnn,omitempty"`
	RedTransExps []RedundantTransmissionExpPerTs `json:"redTransExps" yaml:"redTransExps" bson:"redTransExps,omitempty"`
}