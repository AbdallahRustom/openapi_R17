/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 29.504 V16.9.0; 5G System; Unified Data Repository Services; Stage 3
 * Url: http://www.3gpp.org/ftp/Specs/archive/29_series/29.504/
 *
 * API version: 2.1.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type ProblemDetails1 struct {
	// string providing an URI formatted according to IETF RFC 3986.
	Type string `json:"type,omitempty" yaml:"type" bson:"type,omitempty"`
	// A short, human-readable summary of the problem type. It should not change from occurrence to occurrence of the problem.
	Title string `json:"title,omitempty" yaml:"title" bson:"title,omitempty"`
	// The HTTP status code for this occurrence of the problem.
	Status int32 `json:"status,omitempty" yaml:"status" bson:"status,omitempty"`
	// A human-readable explanation specific to this occurrence of the problem.
	Detail string `json:"detail,omitempty" yaml:"detail" bson:"detail,omitempty"`
	// string providing an URI formatted according to IETF RFC 3986.
	Instance string `json:"instance,omitempty" yaml:"instance" bson:"instance,omitempty"`
	// A machine-readable application error cause specific to this occurrence of the problem. This IE should be present and provide application-related error information, if available.
	Cause string `json:"cause,omitempty" yaml:"cause" bson:"cause,omitempty"`
	// Description of invalid parameters, for a request rejected due to invalid parameters.
	InvalidParams []InvalidParam1 `json:"invalidParams,omitempty" yaml:"invalidParams" bson:"invalidParams,omitempty"`
}
