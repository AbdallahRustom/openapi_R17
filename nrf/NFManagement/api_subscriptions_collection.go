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

type SubscriptionsCollectionApiService service

/*
SubscriptionsCollectionApiService Create a new subscription
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param NrfNfManagementSubscriptionData -
 * @param ContentEncoding - Content-Encoding, described in IETF RFC 7231
 * @param AcceptEncoding - Accept-Encoding, described in IETF RFC 7231

@return CreateSubscriptionResponse
*/

// CreateSubscriptionRequest
type CreateSubscriptionRequest struct {
	NrfNfManagementSubscriptionData *models.NrfNfManagementSubscriptionData
	ContentEncoding                 *string
	AcceptEncoding                  *string
}

func (r *CreateSubscriptionRequest) SetNrfNfManagementSubscriptionData(NrfNfManagementSubscriptionData models.NrfNfManagementSubscriptionData) {
	r.NrfNfManagementSubscriptionData = &NrfNfManagementSubscriptionData
}
func (r *CreateSubscriptionRequest) SetContentEncoding(ContentEncoding string) {
	r.ContentEncoding = &ContentEncoding
}
func (r *CreateSubscriptionRequest) SetAcceptEncoding(AcceptEncoding string) {
	r.AcceptEncoding = &AcceptEncoding
}

type CreateSubscriptionResponse struct {
	AcceptEncoding                  string
	ContentEncoding                 string
	Location                        string
	NrfNfManagementSubscriptionData models.NrfNfManagementSubscriptionData
}

type CreateSubscriptionError struct {
	Location         string
	ProblemDetails   models.ProblemDetails
	RedirectResponse models.RedirectResponse
}

func (a *SubscriptionsCollectionApiService) CreateSubscription(ctx context.Context, request *CreateSubscriptionRequest) (*CreateSubscriptionResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Post")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CreateSubscriptionResponse
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

	if request.ContentEncoding != nil {
		localVarHeaderParams["Content-Encoding"] = openapi.ParameterToString(request.ContentEncoding, "csv")
	}

	if request.AcceptEncoding != nil {
		localVarHeaderParams["Accept-Encoding"] = openapi.ParameterToString(request.AcceptEncoding, "csv")
	}

	// body params
	localVarPostBody = request.NrfNfManagementSubscriptionData

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
		err = openapi.Deserialize(&localVarReturnValue.NrfNfManagementSubscriptionData, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		localVarReturnValue.Location = localVarHTTPResponse.Header.Get("Location")
		localVarReturnValue.AcceptEncoding = localVarHTTPResponse.Header.Get("Accept-Encoding")
		localVarReturnValue.ContentEncoding = localVarHTTPResponse.Header.Get("Content-Encoding")
		return &localVarReturnValue, nil
	case 307:
		var v CreateSubscriptionError
		err = openapi.Deserialize(&v.RedirectResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		v.Location = localVarHTTPResponse.Header.Get("Location")
		apiError.ErrorModel = v
		return nil, apiError
	case 308:
		var v CreateSubscriptionError
		err = openapi.Deserialize(&v.RedirectResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		v.Location = localVarHTTPResponse.Header.Get("Location")
		apiError.ErrorModel = v
		return nil, apiError
	case 400:
		var v CreateSubscriptionError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 401:
		var v CreateSubscriptionError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 403:
		var v CreateSubscriptionError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 404:
		var v CreateSubscriptionError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 411:
		var v CreateSubscriptionError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 413:
		var v CreateSubscriptionError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 415:
		var v CreateSubscriptionError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 429:
		var v CreateSubscriptionError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 500:
		var v CreateSubscriptionError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 501:
		var v CreateSubscriptionError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 503:
		var v CreateSubscriptionError
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

// CreateSubscriptionOnNFStatusEventPostRequest
type CreateSubscriptionOnNFStatusEventPostRequest struct {
	ContentEncoding                 *string
	NrfNfManagementNotificationData *models.NrfNfManagementNotificationData
}

func (r *CreateSubscriptionOnNFStatusEventPostRequest) SetContentEncoding(ContentEncoding string) {
	r.ContentEncoding = &ContentEncoding
}
func (r *CreateSubscriptionOnNFStatusEventPostRequest) SetNrfNfManagementNotificationData(NrfNfManagementNotificationData models.NrfNfManagementNotificationData) {
	r.NrfNfManagementNotificationData = &NrfNfManagementNotificationData
}

type CreateSubscriptionOnNFStatusEventPostResponse struct {
	AcceptEncoding string
}

type CreateSubscriptionOnNFStatusEventPostError struct {
	Location         string
	ProblemDetails   models.ProblemDetails
	RedirectResponse models.RedirectResponse
}

func (a *SubscriptionsCollectionApiService) CreateSubscriptionOnNFStatusEventPost(ctx context.Context, uri string, request *CreateSubscriptionOnNFStatusEventPostRequest) (*CreateSubscriptionOnNFStatusEventPostResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("POST")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CreateSubscriptionOnNFStatusEventPostResponse
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

	if request.ContentEncoding != nil {
		localVarHeaderParams["Content-Encoding"] = openapi.ParameterToString(request.ContentEncoding, "csv")
	}

	// body params

	if request.NrfNfManagementNotificationData != nil {
		localVarPostBody = request.NrfNfManagementNotificationData
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
		var v CreateSubscriptionOnNFStatusEventPostError
		err = openapi.Deserialize(&v.RedirectResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		v.Location = localVarHTTPResponse.Header.Get("Location")
		apiError.ErrorModel = v
		return nil, apiError
	case 308:
		var v CreateSubscriptionOnNFStatusEventPostError
		err = openapi.Deserialize(&v.RedirectResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		v.Location = localVarHTTPResponse.Header.Get("Location")
		apiError.ErrorModel = v
		return nil, apiError
	case 400:
		var v CreateSubscriptionOnNFStatusEventPostError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 401:
		var v CreateSubscriptionOnNFStatusEventPostError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 403:
		var v CreateSubscriptionOnNFStatusEventPostError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 404:
		var v CreateSubscriptionOnNFStatusEventPostError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 411:
		var v CreateSubscriptionOnNFStatusEventPostError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 413:
		var v CreateSubscriptionOnNFStatusEventPostError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 415:
		var v CreateSubscriptionOnNFStatusEventPostError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 429:
		var v CreateSubscriptionOnNFStatusEventPostError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 500:
		var v CreateSubscriptionOnNFStatusEventPostError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 501:
		var v CreateSubscriptionOnNFStatusEventPostError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 503:
		var v CreateSubscriptionOnNFStatusEventPostError
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