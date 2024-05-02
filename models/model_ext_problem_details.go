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

type ExtProblemDetails struct {
	Type               string          `json:"type,omitempty" yaml:"type" bson:"type,omitempty"`
	Title              string          `json:"title,omitempty" yaml:"title" bson:"title,omitempty"`
	Status             int32           `json:"status,omitempty" yaml:"status" bson:"status,omitempty"`
	Detail             string          `json:"detail,omitempty" yaml:"detail" bson:"detail,omitempty"`
	Instance           string          `json:"instance,omitempty" yaml:"instance" bson:"instance,omitempty"`
	Cause              string          `json:"cause,omitempty" yaml:"cause" bson:"cause,omitempty"`
	InvalidParams      []InvalidParam  `json:"invalidParams,omitempty" yaml:"invalidParams" bson:"invalidParams,omitempty"`
	SupportedFeatures  string          `json:"supportedFeatures,omitempty" yaml:"supportedFeatures" bson:"supportedFeatures,omitempty"`
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty" yaml:"accessTokenError" bson:"accessTokenError,omitempty"`
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty" yaml:"accessTokenRequest" bson:"accessTokenRequest,omitempty"`
	NrfId              string          `json:"nrfId,omitempty" yaml:"nrfId" bson:"nrfId,omitempty"`
	RemoteError        bool            `json:"remoteError,omitempty" yaml:"remoteError" bson:"remoteError,omitempty"`
}