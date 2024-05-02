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

type ServiceParameterDataStoreApiService service

/*
ServiceParameterDataStoreApiService Retrieve Service Parameter Data
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ServiceParamIds - Each element identifies a service.
 * @param Dnns - Each element identifies a DNN.
 * @param Snssais - Each element identifies a slice.
 * @param InternalGroupIds - Each element identifies a group of users.
 * @param Supis - Each element identifies the user.
 * @param UeIpv4s - Each element identifies the user.
 * @param UeIpv6s - Each element identifies the user.
 * @param UeMacs - Each element identifies the user.
 * @param AnyUe - Indicates whether the request is for any UE.
 * @param SuppFeat - Supported Features

@return ReadServiceParameterDataResponse
*/

// ReadServiceParameterDataRequest
type ReadServiceParameterDataRequest struct {
	ServiceParamIds  []string
	Dnns             []string
	Snssais          []models.Snssai
	InternalGroupIds []string
	Supis            []string
	UeIpv4s          []string
	UeIpv6s          []string
	UeMacs           []string
	AnyUe            *bool
	SuppFeat         *string
}

func (r *ReadServiceParameterDataRequest) SetServiceParamIds(ServiceParamIds []string) {
	r.ServiceParamIds = ServiceParamIds
}
func (r *ReadServiceParameterDataRequest) SetDnns(Dnns []string) {
	r.Dnns = Dnns
}
func (r *ReadServiceParameterDataRequest) SetSnssais(Snssais []models.Snssai) {
	r.Snssais = Snssais
}
func (r *ReadServiceParameterDataRequest) SetInternalGroupIds(InternalGroupIds []string) {
	r.InternalGroupIds = InternalGroupIds
}
func (r *ReadServiceParameterDataRequest) SetSupis(Supis []string) {
	r.Supis = Supis
}
func (r *ReadServiceParameterDataRequest) SetUeIpv4s(UeIpv4s []string) {
	r.UeIpv4s = UeIpv4s
}
func (r *ReadServiceParameterDataRequest) SetUeIpv6s(UeIpv6s []string) {
	r.UeIpv6s = UeIpv6s
}
func (r *ReadServiceParameterDataRequest) SetUeMacs(UeMacs []string) {
	r.UeMacs = UeMacs
}
func (r *ReadServiceParameterDataRequest) SetAnyUe(AnyUe bool) {
	r.AnyUe = &AnyUe
}
func (r *ReadServiceParameterDataRequest) SetSuppFeat(SuppFeat string) {
	r.SuppFeat = &SuppFeat
}

type ReadServiceParameterDataResponse struct {
	ServiceParameterData []models.ServiceParameterData
}

type ReadServiceParameterDataError struct {
	ProblemDetails models.ProblemDetails
}

func (a *ServiceParameterDataStoreApiService) ReadServiceParameterData(ctx context.Context, request *ReadServiceParameterDataRequest) (*ReadServiceParameterDataResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ReadServiceParameterDataResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/application-data/serviceParamData"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.ServiceParamIds != nil {
		if len(request.ServiceParamIds) < 1 {
			return &localVarReturnValue, openapi.ReportError("ServiceParamIds must have at least 1 elements")
		}
		localVarQueryParams.Add("service-param-ids", openapi.ParameterToString(request.ServiceParamIds, "multi"))
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
		localVarQueryParams.Add("internal-group-ids", openapi.ParameterToString(request.InternalGroupIds, "multi"))
	}
	if request.Supis != nil {
		if len(request.Supis) < 1 {
			return &localVarReturnValue, openapi.ReportError("Supis must have at least 1 elements")
		}
		localVarQueryParams.Add("supis", openapi.ParameterToString(request.Supis, "multi"))
	}
	if request.UeIpv4s != nil {
		if len(request.UeIpv4s) < 1 {
			return &localVarReturnValue, openapi.ReportError("UeIpv4s must have at least 1 elements")
		}
		localVarQueryParams.Add("ue-ipv4s", openapi.ParameterToString(request.UeIpv4s, "multi"))
	}
	if request.UeIpv6s != nil {
		if len(request.UeIpv6s) < 1 {
			return &localVarReturnValue, openapi.ReportError("UeIpv6s must have at least 1 elements")
		}
		localVarQueryParams.Add("ue-ipv6s", openapi.ParameterToString(request.UeIpv6s, "multi"))
	}
	if request.UeMacs != nil {
		if len(request.UeMacs) < 1 {
			return &localVarReturnValue, openapi.ReportError("UeMacs must have at least 1 elements")
		}
		localVarQueryParams.Add("ue-macs", openapi.ParameterToString(request.UeMacs, "multi"))
	}
	if request.AnyUe != nil {
		localVarQueryParams.Add("any-ue", openapi.ParameterToString(request.AnyUe, "multi"))
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
		err = openapi.Deserialize(&localVarReturnValue.ServiceParameterData, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		return &localVarReturnValue, nil
	case 400:
		var v ReadServiceParameterDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 401:
		var v ReadServiceParameterDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 403:
		var v ReadServiceParameterDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 404:
		var v ReadServiceParameterDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 406:
		return &localVarReturnValue, nil
	case 414:
		var v ReadServiceParameterDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 429:
		var v ReadServiceParameterDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 500:
		var v ReadServiceParameterDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 503:
		var v ReadServiceParameterDataError
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
