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

type UsageMonitoringInformationDocumentApiService service

/*
UsageMonitoringInformationDocumentApiService Create a usage monitoring resource
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param UeId -
 * @param UsageMonId -
 * @param UsageMonData -

@return CreateUsageMonitoringResourceResponse
*/

// CreateUsageMonitoringResourceRequest
type CreateUsageMonitoringResourceRequest struct {
	UeId         *string
	UsageMonId   *string
	UsageMonData *models.UsageMonData
}

func (r *CreateUsageMonitoringResourceRequest) SetUeId(UeId string) {
	r.UeId = &UeId
}
func (r *CreateUsageMonitoringResourceRequest) SetUsageMonId(UsageMonId string) {
	r.UsageMonId = &UsageMonId
}
func (r *CreateUsageMonitoringResourceRequest) SetUsageMonData(UsageMonData models.UsageMonData) {
	r.UsageMonData = &UsageMonData
}

type CreateUsageMonitoringResourceResponse struct {
	Location     string
	UsageMonData models.UsageMonData
}

type CreateUsageMonitoringResourceError struct {
	ProblemDetails models.ProblemDetails
}

func (a *UsageMonitoringInformationDocumentApiService) CreateUsageMonitoringResource(ctx context.Context, request *CreateUsageMonitoringResourceRequest) (*CreateUsageMonitoringResourceResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Put")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CreateUsageMonitoringResourceResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/policy-data/ues/{ueId}/sm-data/{usageMonId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", openapi.StringOfValue(*request.UeId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"usageMonId"+"}", openapi.StringOfValue(*request.UsageMonId), -1)

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
	localVarPostBody = request.UsageMonData

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
	case 201:
		err = openapi.Deserialize(&localVarReturnValue.UsageMonData, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		localVarReturnValue.Location = localVarHTTPResponse.Header.Get("Location")
		return &localVarReturnValue, nil
	case 400:
		var v CreateUsageMonitoringResourceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 401:
		var v CreateUsageMonitoringResourceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 403:
		var v CreateUsageMonitoringResourceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 404:
		var v CreateUsageMonitoringResourceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 411:
		var v CreateUsageMonitoringResourceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 413:
		var v CreateUsageMonitoringResourceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 414:
		var v CreateUsageMonitoringResourceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 415:
		var v CreateUsageMonitoringResourceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 429:
		var v CreateUsageMonitoringResourceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 500:
		var v CreateUsageMonitoringResourceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 503:
		var v CreateUsageMonitoringResourceError
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
UsageMonitoringInformationDocumentApiService Delete a usage monitoring resource
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param UeId -
 * @param UsageMonId -

@return DeleteUsageMonitoringInformationResponse
*/

// DeleteUsageMonitoringInformationRequest
type DeleteUsageMonitoringInformationRequest struct {
	UeId       *string
	UsageMonId *string
}

func (r *DeleteUsageMonitoringInformationRequest) SetUeId(UeId string) {
	r.UeId = &UeId
}
func (r *DeleteUsageMonitoringInformationRequest) SetUsageMonId(UsageMonId string) {
	r.UsageMonId = &UsageMonId
}

type DeleteUsageMonitoringInformationResponse struct {
}

type DeleteUsageMonitoringInformationError struct {
	ProblemDetails models.ProblemDetails
}

func (a *UsageMonitoringInformationDocumentApiService) DeleteUsageMonitoringInformation(ctx context.Context, request *DeleteUsageMonitoringInformationRequest) (*DeleteUsageMonitoringInformationResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Delete")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DeleteUsageMonitoringInformationResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/policy-data/ues/{ueId}/sm-data/{usageMonId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", openapi.StringOfValue(*request.UeId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"usageMonId"+"}", openapi.StringOfValue(*request.UsageMonId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/problem+json"}

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
	case 400:
		var v DeleteUsageMonitoringInformationError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 401:
		var v DeleteUsageMonitoringInformationError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 403:
		var v DeleteUsageMonitoringInformationError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 404:
		var v DeleteUsageMonitoringInformationError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 429:
		var v DeleteUsageMonitoringInformationError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 500:
		var v DeleteUsageMonitoringInformationError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 503:
		var v DeleteUsageMonitoringInformationError
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
UsageMonitoringInformationDocumentApiService Retrieve a usage monitoring resource
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param UeId -
 * @param UsageMonId -
 * @param SuppFeat - Supported Features

@return ReadUsageMonitoringInformationResponse
*/

// ReadUsageMonitoringInformationRequest
type ReadUsageMonitoringInformationRequest struct {
	UeId       *string
	UsageMonId *string
	SuppFeat   *string
}

func (r *ReadUsageMonitoringInformationRequest) SetUeId(UeId string) {
	r.UeId = &UeId
}
func (r *ReadUsageMonitoringInformationRequest) SetUsageMonId(UsageMonId string) {
	r.UsageMonId = &UsageMonId
}
func (r *ReadUsageMonitoringInformationRequest) SetSuppFeat(SuppFeat string) {
	r.SuppFeat = &SuppFeat
}

type ReadUsageMonitoringInformationResponse struct {
	UsageMonData models.UsageMonData
}

type ReadUsageMonitoringInformationError struct {
	ProblemDetails models.ProblemDetails
}

func (a *UsageMonitoringInformationDocumentApiService) ReadUsageMonitoringInformation(ctx context.Context, request *ReadUsageMonitoringInformationRequest) (*ReadUsageMonitoringInformationResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ReadUsageMonitoringInformationResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/policy-data/ues/{ueId}/sm-data/{usageMonId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", openapi.StringOfValue(*request.UeId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"usageMonId"+"}", openapi.StringOfValue(*request.UsageMonId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.SuppFeat != nil {
		localVarQueryParams.Add("supp-feat", openapi.ParameterToString(request.SuppFeat, "multi"))
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
		err = openapi.Deserialize(&localVarReturnValue.UsageMonData, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		return &localVarReturnValue, nil
	case 204:
		return &localVarReturnValue, nil
	case 400:
		var v ReadUsageMonitoringInformationError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 401:
		var v ReadUsageMonitoringInformationError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 403:
		var v ReadUsageMonitoringInformationError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 404:
		var v ReadUsageMonitoringInformationError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 406:
		return &localVarReturnValue, nil
	case 414:
		var v ReadUsageMonitoringInformationError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 429:
		var v ReadUsageMonitoringInformationError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 500:
		var v ReadUsageMonitoringInformationError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 503:
		var v ReadUsageMonitoringInformationError
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
