/*
 * Nudm_NIDDAU
 *
 * Nudm NIDD Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.503 Unified Data Management Services, version 17.8.0
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.503/
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package NIDDAuthentication

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

type AuthorizeTheNIDDConfigurationRequestApiService service

/*
AuthorizeTheNIDDConfigurationRequestApiService Authorize the NIDD configuration request.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param UeIdentity - Represents the scope of the UE for which the NIDD configuration are authorized. Contains the GPSI of the user or the external group ID.
 * @param AuthorizationInfo -

@return AuthorizeNiddDataResponse
*/

// AuthorizeNiddDataRequest
type AuthorizeNiddDataRequest struct {
	UeIdentity        *string
	AuthorizationInfo *models.AuthorizationInfo
}

func (r *AuthorizeNiddDataRequest) SetUeIdentity(UeIdentity string) {
	r.UeIdentity = &UeIdentity
}
func (r *AuthorizeNiddDataRequest) SetAuthorizationInfo(AuthorizationInfo models.AuthorizationInfo) {
	r.AuthorizationInfo = &AuthorizationInfo
}

type AuthorizeNiddDataResponse struct {
	UdmNiddauAuthorizationData models.UdmNiddauAuthorizationData
}

type AuthorizeNiddDataError struct {
	ProblemDetails models.ProblemDetails
}

func (a *AuthorizeTheNIDDConfigurationRequestApiService) AuthorizeNiddData(ctx context.Context, request *AuthorizeNiddDataRequest) (*AuthorizeNiddDataResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Post")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AuthorizeNiddDataResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/{ueIdentity}/authorize"
	localVarPath = strings.Replace(localVarPath, "{"+"ueIdentity"+"}", openapi.StringOfValue(*request.UeIdentity), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := strings.Join(localVarHTTPHeaderAccepts, ", ")
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = request.AuthorizationInfo

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
		err = openapi.Deserialize(&localVarReturnValue.UdmNiddauAuthorizationData, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		return &localVarReturnValue, nil
	case 400:
		var v AuthorizeNiddDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 403:
		var v AuthorizeNiddDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 404:
		var v AuthorizeNiddDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 500:
		var v AuthorizeNiddDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 501:
		var v AuthorizeNiddDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 503:
		var v AuthorizeNiddDataError
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

// AuthorizeNiddDataNiddAuthUpdateNotificationPostRequest
type AuthorizeNiddDataNiddAuthUpdateNotificationPostRequest struct {
	NiddAuthUpdateNotification *models.NiddAuthUpdateNotification
}

func (r *AuthorizeNiddDataNiddAuthUpdateNotificationPostRequest) SetNiddAuthUpdateNotification(NiddAuthUpdateNotification models.NiddAuthUpdateNotification) {
	r.NiddAuthUpdateNotification = &NiddAuthUpdateNotification
}

type AuthorizeNiddDataNiddAuthUpdateNotificationPostResponse struct {
}

type AuthorizeNiddDataNiddAuthUpdateNotificationPostError struct {
}

func (a *AuthorizeTheNIDDConfigurationRequestApiService) AuthorizeNiddDataNiddAuthUpdateNotificationPost(ctx context.Context, uri string, request *AuthorizeNiddDataNiddAuthUpdateNotificationPostRequest) (*AuthorizeNiddDataNiddAuthUpdateNotificationPostResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("POST")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AuthorizeNiddDataNiddAuthUpdateNotificationPostResponse
	)

	// create path and map variables
	localVarPath := uri

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

	// body params
	localVarPostBody = request.NiddAuthUpdateNotification

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
		return &localVarReturnValue, apiError
	}
}
