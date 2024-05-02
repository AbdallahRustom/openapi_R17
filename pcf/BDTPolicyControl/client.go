/*
 * Npcf_BDTPolicyControl Service API
 *
 * PCF BDT Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.554 V17.4.0; 5G System; Background Data Transfer Policy Control Service.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.554/
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package BDTPolicyControl

// APIClient manages communication with the Npcf_BDTPolicyControl Service API API v1.2.0
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services
	BDTPoliciesCollectionApi       *BDTPoliciesCollectionApiService
	IndividualBDTPolicyDocumentApi *IndividualBDTPolicyDocumentApiService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.BDTPoliciesCollectionApi = (*BDTPoliciesCollectionApiService)(&c.common)
	c.IndividualBDTPolicyDocumentApi = (*IndividualBDTPolicyDocumentApiService)(&c.common)

	return c
}
