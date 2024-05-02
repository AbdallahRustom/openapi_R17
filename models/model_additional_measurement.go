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

import (
	"time"
)

// Represents additional measurement information.
type AdditionalMeasurement struct {
	UnexpLoc      *NetworkAreaInfo          `json:"unexpLoc,omitempty" yaml:"unexpLoc" bson:"unexpLoc,omitempty"`
	UnexpFlowTeps []IpEthFlowDescription    `json:"unexpFlowTeps,omitempty" yaml:"unexpFlowTeps" bson:"unexpFlowTeps,omitempty"`
	UnexpWakes    []time.Time               `json:"unexpWakes,omitempty" yaml:"unexpWakes" bson:"unexpWakes,omitempty"`
	DdosAttack    *AddressList              `json:"ddosAttack,omitempty" yaml:"ddosAttack" bson:"ddosAttack,omitempty"`
	WrgDest       *AddressList              `json:"wrgDest,omitempty" yaml:"wrgDest" bson:"wrgDest,omitempty"`
	Circums       []CircumstanceDescription `json:"circums,omitempty" yaml:"circums" bson:"circums,omitempty"`
}
