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

// Parameters \"replaceableInd\" and \"rechargeableInd\" are only included if the value of Parameter \"batteryInd\" is true.
type BatteryIndication struct {
	// This IE shall indicate whether the UE is battery powered or not. true: the UE is battery powered; false or absent: the UE is not battery powered
	BatteryInd bool `json:"batteryInd,omitempty" yaml:"batteryInd" bson:"batteryInd,omitempty"`
	// This IE shall indicate whether the battery of the UE is replaceable or not. true: the battery of the UE is replaceable; false or absent: the battery of the UE is not replaceable.
	ReplaceableInd bool `json:"replaceableInd,omitempty" yaml:"replaceableInd" bson:"replaceableInd,omitempty"`
	// This IE shall indicate whether the battery of the UE is rechargeable or not. true: the battery of UE is rechargeable; false or absent: the battery of the UE is not rechargeable.
	RechargeableInd bool `json:"rechargeableInd,omitempty" yaml:"rechargeableInd" bson:"rechargeableInd,omitempty"`
}
