/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.503 Unified Data Management Services, version 17.13.0
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.503/
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package SubscriberDataManagement

import (
	"github.com/AbdallahRustom/openapi"
	"github.com/AbdallahRustom/openapi/models"

	"context"
	"io/ioutil"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type TraceConfigurationDataRetrievalApiService service

/*
TraceConfigurationDataRetrievalApiService retrieve a UE's Trace Configuration Data
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param Supi - Identifier of the UE
 * @param SupportedFeatures - Supported Features
 * @param PlmnId - serving PLMN ID
 * @param IfNoneMatch - Validator for conditional requests, as described in RFC 7232, 3.2
 * @param IfModifiedSince - Validator for conditional requests, as described in RFC 7232, 3.3

@return GetTraceConfigDataResponse
*/

// GetTraceConfigDataRequest
type GetTraceConfigDataRequest struct {
	Supi              *string
	SupportedFeatures *string
	PlmnId            *models.PlmnId
	IfNoneMatch       *string
	IfModifiedSince   *string
}

func (r *GetTraceConfigDataRequest) SetSupi(Supi string) {
	r.Supi = &Supi
}
func (r *GetTraceConfigDataRequest) SetSupportedFeatures(SupportedFeatures string) {
	r.SupportedFeatures = &SupportedFeatures
}
func (r *GetTraceConfigDataRequest) SetPlmnId(PlmnId models.PlmnId) {
	r.PlmnId = &PlmnId
}
func (r *GetTraceConfigDataRequest) SetIfNoneMatch(IfNoneMatch string) {
	r.IfNoneMatch = &IfNoneMatch
}
func (r *GetTraceConfigDataRequest) SetIfModifiedSince(IfModifiedSince string) {
	r.IfModifiedSince = &IfModifiedSince
}

type GetTraceConfigDataResponse struct {
	CacheControl      string
	ETag              string
	LastModified      string
	TraceDataResponse models.TraceDataResponse
}

type GetTraceConfigDataError struct {
	ProblemDetails models.ProblemDetails
}

func (a *TraceConfigurationDataRetrievalApiService) GetTraceConfigData(ctx context.Context, request *GetTraceConfigDataRequest) (*GetTraceConfigDataResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GetTraceConfigDataResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/{supi}/trace-data"
	localVarPath = strings.Replace(localVarPath, "{"+"supi"+"}", openapi.StringOfValue(*request.Supi), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.SupportedFeatures != nil {
		localVarQueryParams.Add("supported-features", openapi.ParameterToString(request.SupportedFeatures, "multi"))
	}
	if request.PlmnId != nil {
		localVarQueryParams.Add("plmn-id", openapi.ParameterToString(request.PlmnId, "application/json"))
	}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := strings.Join(localVarHTTPHeaderAccepts, ", ")
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	if request.IfNoneMatch != nil {
		localVarHeaderParams["If-None-Match"] = openapi.ParameterToString(request.IfNoneMatch, "csv")
	}

	if request.IfModifiedSince != nil {
		localVarHeaderParams["If-Modified-Since"] = openapi.ParameterToString(request.IfModifiedSince, "csv")
	}

	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := openapi.CallAPI(a.client.cfg, r)
	if err != nil || localVarHTTPResponse == nil {
		return nil, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	if err != nil {
		return nil, err
	}
	err = localVarHTTPResponse.Body.Close()
	if err != nil {
		return nil, err
	}

	apiError := openapi.GenericOpenAPIError{
		RawBody:     localVarBody,
		ErrorStatus: localVarHTTPResponse.StatusCode,
	}

	switch localVarHTTPResponse.StatusCode {
	case 200:
		err = openapi.Deserialize(&localVarReturnValue.TraceDataResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		localVarReturnValue.CacheControl = localVarHTTPResponse.Header.Get("Cache-Control")
		localVarReturnValue.ETag = localVarHTTPResponse.Header.Get("ETag")
		localVarReturnValue.LastModified = localVarHTTPResponse.Header.Get("Last-Modified")
		return &localVarReturnValue, nil
	case 400:
		var v GetTraceConfigDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 404:
		var v GetTraceConfigDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 500:
		var v GetTraceConfigDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 503:
		var v GetTraceConfigDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	default:
		return nil, apiError
	}
}
