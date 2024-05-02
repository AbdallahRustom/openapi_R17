/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.504 V17.12.0; 5G System; Unified Data Repository Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.504/
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package DataRepository

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

type NIDDAuthorizationInfoDocumentApiService service

/*
NIDDAuthorizationInfoDocumentApiService Create NIDD Authorization Info
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param UeId -
 * @param NiddAuthorizationInfo -

@return CreateNIDDAuthorizationInfoResponse
*/

// CreateNIDDAuthorizationInfoRequest
type CreateNIDDAuthorizationInfoRequest struct {
	UeId                  *string
	NiddAuthorizationInfo *models.NiddAuthorizationInfo
}

func (r *CreateNIDDAuthorizationInfoRequest) SetUeId(UeId string) {
	r.UeId = &UeId
}
func (r *CreateNIDDAuthorizationInfoRequest) SetNiddAuthorizationInfo(NiddAuthorizationInfo models.NiddAuthorizationInfo) {
	r.NiddAuthorizationInfo = &NiddAuthorizationInfo
}

type CreateNIDDAuthorizationInfoResponse struct {
	NiddAuthorizationInfo models.NiddAuthorizationInfo
}

type CreateNIDDAuthorizationInfoError struct {
}

func (a *NIDDAuthorizationInfoDocumentApiService) CreateNIDDAuthorizationInfo(ctx context.Context, request *CreateNIDDAuthorizationInfoRequest) (*CreateNIDDAuthorizationInfoResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Put")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CreateNIDDAuthorizationInfoResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/{ueId}/context-data/nidd-authorizations"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", openapi.StringOfValue(*request.UeId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := strings.Join(localVarHTTPHeaderAccepts, ", ")
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = request.NiddAuthorizationInfo

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
	case 204:
		return &localVarReturnValue, nil
	case 201:
		err = openapi.Deserialize(&localVarReturnValue.NiddAuthorizationInfo, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		return &localVarReturnValue, nil
	default:
		return nil, apiError
	}
}

/*
NIDDAuthorizationInfoDocumentApiService Retrieve NIDD Authorization Info
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param UeId -

@return GetNiddAuthorizationInfoResponse
*/

// GetNiddAuthorizationInfoRequest
type GetNiddAuthorizationInfoRequest struct {
	UeId *string
}

func (r *GetNiddAuthorizationInfoRequest) SetUeId(UeId string) {
	r.UeId = &UeId
}

type GetNiddAuthorizationInfoResponse struct {
	NiddAuthorizationInfo models.NiddAuthorizationInfo
}

type GetNiddAuthorizationInfoError struct {
}

func (a *NIDDAuthorizationInfoDocumentApiService) GetNiddAuthorizationInfo(ctx context.Context, request *GetNiddAuthorizationInfoRequest) (*GetNiddAuthorizationInfoResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GetNiddAuthorizationInfoResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/{ueId}/context-data/nidd-authorizations"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", openapi.StringOfValue(*request.UeId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := strings.Join(localVarHTTPHeaderAccepts, ", ")
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
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
		err = openapi.Deserialize(&localVarReturnValue.NiddAuthorizationInfo, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		return &localVarReturnValue, nil
	default:
		return nil, apiError
	}
}

/*
NIDDAuthorizationInfoDocumentApiService Modify NIDD Authorization Info
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param UeId -
 * @param PatchItem -
 * @param SupportedFeatures - Features required to be supported by the target NF

@return ModifyNiddAuthorizationInfoResponse
*/

// ModifyNiddAuthorizationInfoRequest
type ModifyNiddAuthorizationInfoRequest struct {
	UeId              *string
	PatchItem         []models.PatchItem
	SupportedFeatures *string
}

func (r *ModifyNiddAuthorizationInfoRequest) SetUeId(UeId string) {
	r.UeId = &UeId
}
func (r *ModifyNiddAuthorizationInfoRequest) SetPatchItem(PatchItem []models.PatchItem) {
	r.PatchItem = PatchItem
}
func (r *ModifyNiddAuthorizationInfoRequest) SetSupportedFeatures(SupportedFeatures string) {
	r.SupportedFeatures = &SupportedFeatures
}

type ModifyNiddAuthorizationInfoResponse struct {
	PatchResult models.PatchResult
}

type ModifyNiddAuthorizationInfoError struct {
	ProblemDetails models.ProblemDetails
}

func (a *NIDDAuthorizationInfoDocumentApiService) ModifyNiddAuthorizationInfo(ctx context.Context, request *ModifyNiddAuthorizationInfoRequest) (*ModifyNiddAuthorizationInfoResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Patch")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ModifyNiddAuthorizationInfoResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/{ueId}/context-data/nidd-authorizations"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", openapi.StringOfValue(*request.UeId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.SupportedFeatures != nil {
		localVarQueryParams.Add("supported-features", openapi.ParameterToString(request.SupportedFeatures, "multi"))
	}
	localVarHTTPContentTypes := []string{"application/json-patch+json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := strings.Join(localVarHTTPHeaderAccepts, ", ")
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = request.PatchItem

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
	case 204:
		return &localVarReturnValue, nil
	case 200:
		err = openapi.Deserialize(&localVarReturnValue.PatchResult, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		return &localVarReturnValue, nil
	case 403:
		var v ModifyNiddAuthorizationInfoError
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

/*
NIDDAuthorizationInfoDocumentApiService Delete NIDD Authorization Info
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param UeId -

@return RemoveNiddAuthorizationInfoResponse
*/

// RemoveNiddAuthorizationInfoRequest
type RemoveNiddAuthorizationInfoRequest struct {
	UeId *string
}

func (r *RemoveNiddAuthorizationInfoRequest) SetUeId(UeId string) {
	r.UeId = &UeId
}

type RemoveNiddAuthorizationInfoResponse struct {
}

type RemoveNiddAuthorizationInfoError struct {
}

func (a *NIDDAuthorizationInfoDocumentApiService) RemoveNiddAuthorizationInfo(ctx context.Context, request *RemoveNiddAuthorizationInfoRequest) (*RemoveNiddAuthorizationInfoResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Delete")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  RemoveNiddAuthorizationInfoResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/{ueId}/context-data/nidd-authorizations"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", openapi.StringOfValue(*request.UeId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := strings.Join(localVarHTTPHeaderAccepts, ", ")
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
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
	case 204:
		return &localVarReturnValue, nil
	default:
		return nil, apiError
	}
}