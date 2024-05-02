/*
 * Nudsf_Timer
 *
 * Nudsf Timer Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.598 UDSF Services, V17.6.0.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.598/
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Timer

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

type TimerUpdateApiService service

/*
TimerUpdateApiService Timer modification
update a specific Timer
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param RealmId - Identifier of the Realm
 * @param StorageId - Identifier of the Storage
 * @param TimerId - Identifier of the Timer
 * @param PatchItem - Timer data to patch
 * @param SupportedFeatures - Features required to be supported by the target NF

@return UpdateTimerResponse
*/

// UpdateTimerRequest
type UpdateTimerRequest struct {
	RealmId           *string
	StorageId         *string
	TimerId           *string
	PatchItem         []models.PatchItem
	SupportedFeatures *string
}

func (r *UpdateTimerRequest) SetRealmId(RealmId string) {
	r.RealmId = &RealmId
}
func (r *UpdateTimerRequest) SetStorageId(StorageId string) {
	r.StorageId = &StorageId
}
func (r *UpdateTimerRequest) SetTimerId(TimerId string) {
	r.TimerId = &TimerId
}
func (r *UpdateTimerRequest) SetPatchItem(PatchItem []models.PatchItem) {
	r.PatchItem = PatchItem
}
func (r *UpdateTimerRequest) SetSupportedFeatures(SupportedFeatures string) {
	r.SupportedFeatures = &SupportedFeatures
}

type UpdateTimerResponse struct {
	PatchResult models.PatchResult
}

type UpdateTimerError struct {
	ProblemDetails models.ProblemDetails
}

func (a *TimerUpdateApiService) UpdateTimer(ctx context.Context, request *UpdateTimerRequest) (*UpdateTimerResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Patch")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  UpdateTimerResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/{realmId}/{storageId}/timers/{timerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"realmId"+"}", openapi.StringOfValue(*request.RealmId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"storageId"+"}", openapi.StringOfValue(*request.StorageId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"timerId"+"}", openapi.StringOfValue(*request.TimerId), -1)

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
	case 200:
		err = openapi.Deserialize(&localVarReturnValue.PatchResult, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		return &localVarReturnValue, nil
	case 204:
		return &localVarReturnValue, nil
	case 400:
		var v UpdateTimerError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 401:
		var v UpdateTimerError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 403:
		var v UpdateTimerError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 404:
		var v UpdateTimerError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 500:
		var v UpdateTimerError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 503:
		var v UpdateTimerError
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
