/*
 * LMF Location
 *
 * LMF Location Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.572 V17.9.0; 5G System; Location Management Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.572/
 *
 * API version: 1.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Horizontal and vertical velocity with speed uncertainty.
type HorizontalWithVerticalVelocityAndUncertainty struct {
	// Indicates value of horizontal speed.
	HSpeed float32 `json:"hSpeed" yaml:"hSpeed" bson:"hSpeed,omitempty"`
	// Indicates value of angle.
	Bearing int32 `json:"bearing" yaml:"bearing" bson:"bearing,omitempty"`
	// Indicates value of vertical speed.
	VSpeed     float32           `json:"vSpeed" yaml:"vSpeed" bson:"vSpeed,omitempty"`
	VDirection VerticalDirection `json:"vDirection" yaml:"vDirection" bson:"vDirection,omitempty"`
	// Indicates value of speed uncertainty.
	HUncertainty float32 `json:"hUncertainty" yaml:"hUncertainty" bson:"hUncertainty,omitempty"`
	// Indicates value of speed uncertainty.
	VUncertainty float32 `json:"vUncertainty" yaml:"vUncertainty" bson:"vUncertainty,omitempty"`
}
