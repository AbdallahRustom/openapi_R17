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

type UserConsentSubscriptionDataRetrievalApiService service

/*
UserConsentSubscriptionDataRetrievalApiService retrieve a UE's User Consent Subscription Data
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param Supi - Identifier of the UE
 * @param SupportedFeatures - Supported Features
 * @param UcPurpose - User consent purpose
 * @param IfNoneMatch - Validator for conditional requests, as described in RFC 7232, 3.2
 * @param IfModifiedSince - Validator for conditional requests, as described in RFC 7232, 3.3

@return GetUcDataResponse
*/

// GetUcDataRequest
type GetUcDataRequest struct {
	Supi              *string
	SupportedFeatures *string
	UcPurpose         *models.UcPurpose
	IfNoneMatch       *string
	IfModifiedSince   *string
}

func (r *GetUcDataRequest) SetSupi(Supi string) {
	r.Supi = &Supi
}
func (r *GetUcDataRequest) SetSupportedFeatures(SupportedFeatures string) {
	r.SupportedFeatures = &SupportedFeatures
}
func (r *GetUcDataRequest) SetUcPurpose(UcPurpose models.UcPurpose) {
	r.UcPurpose = &UcPurpose
}
func (r *GetUcDataRequest) SetIfNoneMatch(IfNoneMatch string) {
	r.IfNoneMatch = &IfNoneMatch
}
func (r *GetUcDataRequest) SetIfModifiedSince(IfModifiedSince string) {
	r.IfModifiedSince = &IfModifiedSince
}

type GetUcDataResponse struct {
	CacheControl       string
	ETag               string
	LastModified       string
	UcSubscriptionData models.UcSubscriptionData
}

type GetUcDataError struct {
	ProblemDetails models.ProblemDetails
}

func (a *UserConsentSubscriptionDataRetrievalApiService) GetUcData(ctx context.Context, request *GetUcDataRequest) (*GetUcDataResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GetUcDataResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/{supi}/uc-data"
	localVarPath = strings.Replace(localVarPath, "{"+"supi"+"}", openapi.StringOfValue(*request.Supi), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.SupportedFeatures != nil {
		localVarQueryParams.Add("supported-features", openapi.ParameterToString(request.SupportedFeatures, "multi"))
	}
	if request.UcPurpose != nil {
		localVarQueryParams.Add("uc-purpose", openapi.ParameterToString(request.UcPurpose, "multi"))
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
		err = openapi.Deserialize(&localVarReturnValue.UcSubscriptionData, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		localVarReturnValue.CacheControl = localVarHTTPResponse.Header.Get("Cache-Control")
		localVarReturnValue.ETag = localVarHTTPResponse.Header.Get("ETag")
		localVarReturnValue.LastModified = localVarHTTPResponse.Header.Get("Last-Modified")
		return &localVarReturnValue, nil
	case 400:
		var v GetUcDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 404:
		var v GetUcDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 500:
		var v GetUcDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 503:
		var v GetUcDataError
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
