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

type InfluenceDataStoreApiService service

/*
InfluenceDataStoreApiService Retrieve Traffic Influence Data
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param InfluenceIds - Each element identifies a service.
 * @param Dnns - Each element identifies a DNN.
 * @param Snssais - Each element identifies a slice.
 * @param InternalGroupIds - Each element identifies a group of users.
 * @param Supis - Each element identifies the user.
 * @param SuppFeat - Supported Features

@return ReadInfluenceDataResponse
*/

// ReadInfluenceDataRequest
type ReadInfluenceDataRequest struct {
	InfluenceIds     []string
	Dnns             []string
	Snssais          []models.Snssai
	InternalGroupIds []string
	Supis            []string
	SuppFeat         *string
}

func (r *ReadInfluenceDataRequest) SetInfluenceIds(InfluenceIds []string) {
	r.InfluenceIds = InfluenceIds
}
func (r *ReadInfluenceDataRequest) SetDnns(Dnns []string) {
	r.Dnns = Dnns
}
func (r *ReadInfluenceDataRequest) SetSnssais(Snssais []models.Snssai) {
	r.Snssais = Snssais
}
func (r *ReadInfluenceDataRequest) SetInternalGroupIds(InternalGroupIds []string) {
	r.InternalGroupIds = InternalGroupIds
}
func (r *ReadInfluenceDataRequest) SetSupis(Supis []string) {
	r.Supis = Supis
}
func (r *ReadInfluenceDataRequest) SetSuppFeat(SuppFeat string) {
	r.SuppFeat = &SuppFeat
}

type ReadInfluenceDataResponse struct {
	TrafficInfluData []models.TrafficInfluData
}

type ReadInfluenceDataError struct {
	ProblemDetails models.ProblemDetails
}

func (a *InfluenceDataStoreApiService) ReadInfluenceData(ctx context.Context, request *ReadInfluenceDataRequest) (*ReadInfluenceDataResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ReadInfluenceDataResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/application-data/influenceData"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.InfluenceIds != nil {
		if len(request.InfluenceIds) < 1 {
			return &localVarReturnValue, openapi.ReportError("InfluenceIds must have at least 1 elements")
		}
		localVarQueryParams.Add("influence-Ids", openapi.ParameterToString(request.InfluenceIds, "multi"))
	}
	if request.Dnns != nil {
		if len(request.Dnns) < 1 {
			return &localVarReturnValue, openapi.ReportError("Dnns must have at least 1 elements")
		}
		localVarQueryParams.Add("dnns", openapi.ParameterToString(request.Dnns, "multi"))
	}
	if request.Snssais != nil {
		if len(request.Snssais) < 1 {
			return &localVarReturnValue, openapi.ReportError("Snssais must have at least 1 elements")
		}
		localVarQueryParams.Add("snssais", openapi.ParameterToString(request.Snssais, "application/json"))
	}
	if request.InternalGroupIds != nil {
		if len(request.InternalGroupIds) < 1 {
			return &localVarReturnValue, openapi.ReportError("InternalGroupIds must have at least 1 elements")
		}
		localVarQueryParams.Add("internal-Group-Ids", openapi.ParameterToString(request.InternalGroupIds, "multi"))
	}
	if request.Supis != nil {
		if len(request.Supis) < 1 {
			return &localVarReturnValue, openapi.ReportError("Supis must have at least 1 elements")
		}
		localVarQueryParams.Add("supis", openapi.ParameterToString(request.Supis, "multi"))
	}
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
		err = openapi.Deserialize(&localVarReturnValue.TrafficInfluData, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		return &localVarReturnValue, nil
	case 400:
		var v ReadInfluenceDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 401:
		var v ReadInfluenceDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 403:
		var v ReadInfluenceDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 404:
		var v ReadInfluenceDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 406:
		return &localVarReturnValue, nil
	case 414:
		var v ReadInfluenceDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 429:
		var v ReadInfluenceDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 500:
		var v ReadInfluenceDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 503:
		var v ReadInfluenceDataError
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