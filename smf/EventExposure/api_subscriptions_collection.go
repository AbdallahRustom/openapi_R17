/*
 * Nsmf_EventExposure
 *
 * Session Management Event Exposure Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.508 V17.10.0; 5G System; Session Management Event Exposure Service.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.508/
 *
 * API version: 1.2.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package EventExposure

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

type SubscriptionsCollectionApiService service

/*
SubscriptionsCollectionApiService Create an individual subscription for event notifications from the SMF
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param NsmfEventExposure -

@return CreateIndividualSubcriptionResponse
*/

// CreateIndividualSubcriptionRequest
type CreateIndividualSubcriptionRequest struct {
	NsmfEventExposure *models.NsmfEventExposure
}

func (r *CreateIndividualSubcriptionRequest) SetNsmfEventExposure(NsmfEventExposure models.NsmfEventExposure) {
	r.NsmfEventExposure = &NsmfEventExposure
}

type CreateIndividualSubcriptionResponse struct {
	Location          string
	NsmfEventExposure models.NsmfEventExposure
}

type CreateIndividualSubcriptionError struct {
	ProblemDetails models.ProblemDetails
}

func (a *SubscriptionsCollectionApiService) CreateIndividualSubcription(ctx context.Context, request *CreateIndividualSubcriptionRequest) (*CreateIndividualSubcriptionResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Post")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CreateIndividualSubcriptionResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscriptions"

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
	localVarPostBody = request.NsmfEventExposure

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
		err = openapi.Deserialize(&localVarReturnValue.NsmfEventExposure, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		localVarReturnValue.Location = localVarHTTPResponse.Header.Get("Location")
		return &localVarReturnValue, nil
	case 400:
		var v CreateIndividualSubcriptionError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 401:
		var v CreateIndividualSubcriptionError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 403:
		var v CreateIndividualSubcriptionError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 404:
		var v CreateIndividualSubcriptionError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 411:
		var v CreateIndividualSubcriptionError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 413:
		var v CreateIndividualSubcriptionError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 415:
		var v CreateIndividualSubcriptionError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 429:
		var v CreateIndividualSubcriptionError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 500:
		var v CreateIndividualSubcriptionError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 503:
		var v CreateIndividualSubcriptionError
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

// CreateIndividualSubcriptionMyNotificationPostRequest
type CreateIndividualSubcriptionMyNotificationPostRequest struct {
	NsmfEventExposureNotification *models.NsmfEventExposureNotification
}

func (r *CreateIndividualSubcriptionMyNotificationPostRequest) SetNsmfEventExposureNotification(NsmfEventExposureNotification models.NsmfEventExposureNotification) {
	r.NsmfEventExposureNotification = &NsmfEventExposureNotification
}

type CreateIndividualSubcriptionMyNotificationPostResponse struct {
}

type CreateIndividualSubcriptionMyNotificationPostError struct {
	Location             string
	Var3gppSbiTargetNfId string
	ProblemDetails       models.ProblemDetails
	RedirectResponse     models.RedirectResponse
}

func (a *SubscriptionsCollectionApiService) CreateIndividualSubcriptionMyNotificationPost(ctx context.Context, uri string, request *CreateIndividualSubcriptionMyNotificationPostRequest) (*CreateIndividualSubcriptionMyNotificationPostResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("POST")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CreateIndividualSubcriptionMyNotificationPostResponse
	)

	// create path and map variables
	localVarPath := uri

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
	localVarPostBody = request.NsmfEventExposureNotification

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
		var v CreateIndividualSubcriptionMyNotificationPostError
		err = openapi.Deserialize(&v.RedirectResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		v.Location = localVarHTTPResponse.Header.Get("Location")
		v.Var3gppSbiTargetNfId = localVarHTTPResponse.Header.Get("3gpp-Sbi-Target-Nf-Id")
		apiError.ErrorModel = v
		return nil, apiError
	case 308:
		var v CreateIndividualSubcriptionMyNotificationPostError
		err = openapi.Deserialize(&v.RedirectResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		v.Location = localVarHTTPResponse.Header.Get("Location")
		v.Var3gppSbiTargetNfId = localVarHTTPResponse.Header.Get("3gpp-Sbi-Target-Nf-Id")
		apiError.ErrorModel = v
		return nil, apiError
	case 400:
		var v CreateIndividualSubcriptionMyNotificationPostError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 401:
		var v CreateIndividualSubcriptionMyNotificationPostError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 403:
		var v CreateIndividualSubcriptionMyNotificationPostError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 404:
		var v CreateIndividualSubcriptionMyNotificationPostError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 411:
		var v CreateIndividualSubcriptionMyNotificationPostError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 413:
		var v CreateIndividualSubcriptionMyNotificationPostError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 415:
		var v CreateIndividualSubcriptionMyNotificationPostError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 429:
		var v CreateIndividualSubcriptionMyNotificationPostError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 500:
		var v CreateIndividualSubcriptionMyNotificationPostError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 503:
		var v CreateIndividualSubcriptionMyNotificationPostError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	default:
		return &localVarReturnValue, nil
	}
}
