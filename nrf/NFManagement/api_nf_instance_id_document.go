/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.510 V17.12.0; 5G System; Network Function Repository Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.510/
 *
 * API version: 1.2.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package NFManagement

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

type NFInstanceIDDocumentApiService service

/*
NFInstanceIDDocumentApiService Deregisters a given NF Instance
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param NfInstanceID - Unique ID of the NF Instance to deregister

@return DeregisterNFInstanceResponse
*/

// DeregisterNFInstanceRequest
type DeregisterNFInstanceRequest struct {
	NfInstanceID *string
}

func (r *DeregisterNFInstanceRequest) SetNfInstanceID(NfInstanceID string) {
	r.NfInstanceID = &NfInstanceID
}

type DeregisterNFInstanceResponse struct {
}

type DeregisterNFInstanceError struct {
	Location         string
	ProblemDetails   models.ProblemDetails
	RedirectResponse models.RedirectResponse
}

func (a *NFInstanceIDDocumentApiService) DeregisterNFInstance(ctx context.Context, request *DeregisterNFInstanceRequest) (*DeregisterNFInstanceResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Delete")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DeregisterNFInstanceResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/nf-instances/{nfInstanceID}"
	localVarPath = strings.Replace(localVarPath, "{"+"nfInstanceID"+"}", openapi.StringOfValue(*request.NfInstanceID), -1)

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
	case 307:
		var v DeregisterNFInstanceError
		err = openapi.Deserialize(&v.RedirectResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		v.Location = localVarHTTPResponse.Header.Get("Location")
		apiError.ErrorModel = v
		return nil, apiError
	case 308:
		var v DeregisterNFInstanceError
		err = openapi.Deserialize(&v.RedirectResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		v.Location = localVarHTTPResponse.Header.Get("Location")
		apiError.ErrorModel = v
		return nil, apiError
	case 400:
		var v DeregisterNFInstanceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 401:
		var v DeregisterNFInstanceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 403:
		var v DeregisterNFInstanceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 404:
		var v DeregisterNFInstanceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 411:
		var v DeregisterNFInstanceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 429:
		var v DeregisterNFInstanceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 500:
		var v DeregisterNFInstanceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 501:
		var v DeregisterNFInstanceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 503:
		var v DeregisterNFInstanceError
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
NFInstanceIDDocumentApiService Read the profile of a given NF Instance
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param NfInstanceID - Unique ID of the NF Instance
 * @param RequesterFeatures - Features supported by the NF Service Consumer

@return GetNFInstanceResponse
*/

// GetNFInstanceRequest
type GetNFInstanceRequest struct {
	NfInstanceID      *string
	RequesterFeatures *string
}

func (r *GetNFInstanceRequest) SetNfInstanceID(NfInstanceID string) {
	r.NfInstanceID = &NfInstanceID
}
func (r *GetNFInstanceRequest) SetRequesterFeatures(RequesterFeatures string) {
	r.RequesterFeatures = &RequesterFeatures
}

type GetNFInstanceResponse struct {
	ETag                     string
	NrfNfManagementNfProfile models.NrfNfManagementNfProfile
}

type GetNFInstanceError struct {
	Location         string
	ProblemDetails   models.ProblemDetails
	RedirectResponse models.RedirectResponse
}

func (a *NFInstanceIDDocumentApiService) GetNFInstance(ctx context.Context, request *GetNFInstanceRequest) (*GetNFInstanceResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GetNFInstanceResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/nf-instances/{nfInstanceID}"
	localVarPath = strings.Replace(localVarPath, "{"+"nfInstanceID"+"}", openapi.StringOfValue(*request.NfInstanceID), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.RequesterFeatures != nil {
		localVarQueryParams.Add("requester-features", openapi.ParameterToString(request.RequesterFeatures, "multi"))
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
		err = openapi.Deserialize(&localVarReturnValue.NrfNfManagementNfProfile, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		localVarReturnValue.ETag = localVarHTTPResponse.Header.Get("ETag")
		return &localVarReturnValue, nil
	case 307:
		var v GetNFInstanceError
		err = openapi.Deserialize(&v.RedirectResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		v.Location = localVarHTTPResponse.Header.Get("Location")
		apiError.ErrorModel = v
		return nil, apiError
	case 308:
		var v GetNFInstanceError
		err = openapi.Deserialize(&v.RedirectResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		v.Location = localVarHTTPResponse.Header.Get("Location")
		apiError.ErrorModel = v
		return nil, apiError
	case 400:
		var v GetNFInstanceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 401:
		var v GetNFInstanceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 403:
		var v GetNFInstanceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 404:
		var v GetNFInstanceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 406:
		return &localVarReturnValue, nil
	case 411:
		var v GetNFInstanceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 413:
		var v GetNFInstanceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 415:
		var v GetNFInstanceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 429:
		var v GetNFInstanceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 500:
		var v GetNFInstanceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 501:
		var v GetNFInstanceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 503:
		var v GetNFInstanceError
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
NFInstanceIDDocumentApiService Register a new NF Instance
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param NfInstanceID - Unique ID of the NF Instance to register
 * @param NrfNfManagementNfProfile -
 * @param ContentEncoding - Content-Encoding, described in IETF RFC 7231
 * @param AcceptEncoding - Accept-Encoding, described in IETF RFC 7231

@return RegisterNFInstanceResponse
*/

// RegisterNFInstanceRequest
type RegisterNFInstanceRequest struct {
	NfInstanceID             *string
	NrfNfManagementNfProfile *models.NrfNfManagementNfProfile
	ContentEncoding          *string
	AcceptEncoding           *string
}

func (r *RegisterNFInstanceRequest) SetNfInstanceID(NfInstanceID string) {
	r.NfInstanceID = &NfInstanceID
}
func (r *RegisterNFInstanceRequest) SetNrfNfManagementNfProfile(NrfNfManagementNfProfile models.NrfNfManagementNfProfile) {
	r.NrfNfManagementNfProfile = &NrfNfManagementNfProfile
}
func (r *RegisterNFInstanceRequest) SetContentEncoding(ContentEncoding string) {
	r.ContentEncoding = &ContentEncoding
}
func (r *RegisterNFInstanceRequest) SetAcceptEncoding(AcceptEncoding string) {
	r.AcceptEncoding = &AcceptEncoding
}

type RegisterNFInstanceResponse struct {
	AcceptEncoding           string
	ContentEncoding          string
	ETag                     string
	Location                 string
	NrfNfManagementNfProfile models.NrfNfManagementNfProfile
}

type RegisterNFInstanceError struct {
	Location         string
	ProblemDetails   models.ProblemDetails
	RedirectResponse models.RedirectResponse
}

func (a *NFInstanceIDDocumentApiService) RegisterNFInstance(ctx context.Context, request *RegisterNFInstanceRequest) (*RegisterNFInstanceResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Put")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  RegisterNFInstanceResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/nf-instances/{nfInstanceID}"
	localVarPath = strings.Replace(localVarPath, "{"+"nfInstanceID"+"}", openapi.StringOfValue(*request.NfInstanceID), -1)

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

	if request.ContentEncoding != nil {
		localVarHeaderParams["Content-Encoding"] = openapi.ParameterToString(request.ContentEncoding, "csv")
	}

	if request.AcceptEncoding != nil {
		localVarHeaderParams["Accept-Encoding"] = openapi.ParameterToString(request.AcceptEncoding, "csv")
	}

	// body params
	localVarPostBody = request.NrfNfManagementNfProfile

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
		err = openapi.Deserialize(&localVarReturnValue.NrfNfManagementNfProfile, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		localVarReturnValue.AcceptEncoding = localVarHTTPResponse.Header.Get("Accept-Encoding")
		localVarReturnValue.ContentEncoding = localVarHTTPResponse.Header.Get("Content-Encoding")
		localVarReturnValue.ETag = localVarHTTPResponse.Header.Get("ETag")
		return &localVarReturnValue, nil
	case 201:
		err = openapi.Deserialize(&localVarReturnValue.NrfNfManagementNfProfile, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		localVarReturnValue.Location = localVarHTTPResponse.Header.Get("Location")
		localVarReturnValue.AcceptEncoding = localVarHTTPResponse.Header.Get("Accept-Encoding")
		localVarReturnValue.ContentEncoding = localVarHTTPResponse.Header.Get("Content-Encoding")
		localVarReturnValue.ETag = localVarHTTPResponse.Header.Get("ETag")
		return &localVarReturnValue, nil
	case 307:
		var v RegisterNFInstanceError
		err = openapi.Deserialize(&v.RedirectResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		v.Location = localVarHTTPResponse.Header.Get("Location")
		apiError.ErrorModel = v
		return nil, apiError
	case 308:
		var v RegisterNFInstanceError
		err = openapi.Deserialize(&v.RedirectResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		v.Location = localVarHTTPResponse.Header.Get("Location")
		apiError.ErrorModel = v
		return nil, apiError
	case 400:
		var v RegisterNFInstanceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 401:
		var v RegisterNFInstanceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 403:
		var v RegisterNFInstanceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 404:
		var v RegisterNFInstanceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 411:
		var v RegisterNFInstanceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 413:
		var v RegisterNFInstanceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 415:
		var v RegisterNFInstanceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 429:
		var v RegisterNFInstanceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 500:
		var v RegisterNFInstanceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 501:
		var v RegisterNFInstanceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 503:
		var v RegisterNFInstanceError
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
NFInstanceIDDocumentApiService Update NF Instance profile
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param NfInstanceID - Unique ID of the NF Instance to update
 * @param PatchItem -
 * @param ContentEncoding - Content-Encoding, described in IETF RFC 7231
 * @param AcceptEncoding - Accept-Encoding, described in IETF RFC 7231
 * @param IfMatch - Validator for conditional requests, as described in IETF RFC 7232, 3.2

@return UpdateNFInstanceResponse
*/

// UpdateNFInstanceRequest
type UpdateNFInstanceRequest struct {
	NfInstanceID    *string
	PatchItem       []models.PatchItem
	ContentEncoding *string
	AcceptEncoding  *string
	IfMatch         *string
}

func (r *UpdateNFInstanceRequest) SetNfInstanceID(NfInstanceID string) {
	r.NfInstanceID = &NfInstanceID
}
func (r *UpdateNFInstanceRequest) SetPatchItem(PatchItem []models.PatchItem) {
	r.PatchItem = PatchItem
}
func (r *UpdateNFInstanceRequest) SetContentEncoding(ContentEncoding string) {
	r.ContentEncoding = &ContentEncoding
}
func (r *UpdateNFInstanceRequest) SetAcceptEncoding(AcceptEncoding string) {
	r.AcceptEncoding = &AcceptEncoding
}
func (r *UpdateNFInstanceRequest) SetIfMatch(IfMatch string) {
	r.IfMatch = &IfMatch
}

type UpdateNFInstanceResponse struct {
	AcceptEncoding           string
	ContentEncoding          string
	ETag                     string
	NrfNfManagementNfProfile models.NrfNfManagementNfProfile
}

type UpdateNFInstanceError struct {
	Location         string
	ProblemDetails   models.ProblemDetails
	RedirectResponse models.RedirectResponse
}

func (a *NFInstanceIDDocumentApiService) UpdateNFInstance(ctx context.Context, request *UpdateNFInstanceRequest) (*UpdateNFInstanceResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Patch")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  UpdateNFInstanceResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/nf-instances/{nfInstanceID}"
	localVarPath = strings.Replace(localVarPath, "{"+"nfInstanceID"+"}", openapi.StringOfValue(*request.NfInstanceID), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHTTPContentTypes := []string{"application/json-patch+json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := strings.Join(localVarHTTPHeaderAccepts, ", ")
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	if request.ContentEncoding != nil {
		localVarHeaderParams["Content-Encoding"] = openapi.ParameterToString(request.ContentEncoding, "csv")
	}

	if request.AcceptEncoding != nil {
		localVarHeaderParams["Accept-Encoding"] = openapi.ParameterToString(request.AcceptEncoding, "csv")
	}

	if request.IfMatch != nil {
		localVarHeaderParams["If-Match"] = openapi.ParameterToString(request.IfMatch, "csv")
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
		err = openapi.Deserialize(&localVarReturnValue.NrfNfManagementNfProfile, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		localVarReturnValue.AcceptEncoding = localVarHTTPResponse.Header.Get("Accept-Encoding")
		localVarReturnValue.ETag = localVarHTTPResponse.Header.Get("ETag")
		localVarReturnValue.ContentEncoding = localVarHTTPResponse.Header.Get("Content-Encoding")
		return &localVarReturnValue, nil
	case 204:
		return &localVarReturnValue, nil
	case 307:
		var v UpdateNFInstanceError
		err = openapi.Deserialize(&v.RedirectResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		v.Location = localVarHTTPResponse.Header.Get("Location")
		apiError.ErrorModel = v
		return nil, apiError
	case 308:
		var v UpdateNFInstanceError
		err = openapi.Deserialize(&v.RedirectResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		v.Location = localVarHTTPResponse.Header.Get("Location")
		apiError.ErrorModel = v
		return nil, apiError
	case 400:
		var v UpdateNFInstanceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 403:
		var v UpdateNFInstanceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 404:
		var v UpdateNFInstanceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 409:
		var v UpdateNFInstanceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 411:
		var v UpdateNFInstanceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 412:
		var v UpdateNFInstanceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 413:
		var v UpdateNFInstanceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 415:
		var v UpdateNFInstanceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 429:
		var v UpdateNFInstanceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 500:
		var v UpdateNFInstanceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 501:
		var v UpdateNFInstanceError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 503:
		var v UpdateNFInstanceError
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
