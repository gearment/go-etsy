/*
Etsy Open API v3

<div class=\"wt-text-body-01\"><p class=\"wt-pt-xs-2 wt-pb-xs-2\">Etsy's Open API provides a simple RESTful interface for various Etsy.com features. The API endpoints are meant to replace Etsy's Open API v2, which is scheduled to end service in 2022.</p><p class=\"wt-pb-xs-2\">All of the endpoints are callable and the majority of the API endpoints are now in a beta phase. This means we do not expect to make any breaking changes before our general release. A handful of endpoints are currently interface stubs (labeled “Feedback Only”) and returns a \"501 Not Implemented\" response code when called.</p><p class=\"wt-pb-xs-2\">If you'd like to report an issue or provide feedback on the API design, <a target=\"_blank\" class=\"wt-text-link wt-p-xs-0\" href=\"https://github.com/etsy/open-api/discussions\">please add an issue in Github</a>.</p></div>&copy; 2021-2024 Etsy, Inc. All Rights Reserved. Use of this code is subject to Etsy's <a class='wt-text-link wt-p-xs-0' target='_blank' href='https://www.etsy.com/legal/api'>API Developer Terms of Use</a>.

API version: 3.0.0
Contact: developers@etsy.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package goEtsy

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type ReviewAPI interface {

	/*
			GetReviewsByListing Method for GetReviewsByListing

			<div class="wt-display-flex-xs wt-align-items-center wt-mt-xs-2 wt-mb-xs-3"><span class="wt-badge wt-badge--notificationPrimary wt-bg-slime-tint wt-mr-xs-2">General Release</span><a class="wt-text-link" href="https://github.com/etsy/open-api/discussions" target="_blank" rel="noopener noreferrer">Report bug</a></div><div class="wt-display-flex-xs wt-align-items-center wt-mt-xs-2 wt-mb-xs-3"><p class="wt-text-body-01 banner-text">This endpoint is ready for production use.</p></div>

		Open API V3 to retrieve the reviews for a listing given its ID.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param listingId The numeric ID for the [listing](/documentation/reference#tag/ShopListing) associated to this transaction.
			@return ReviewAPIGetReviewsByListingRequest
	*/
	GetReviewsByListing(ctx context.Context, listingId int64) ReviewAPIGetReviewsByListingRequest

	// GetReviewsByListingExecute executes the request
	//  @return ListingReviews
	GetReviewsByListingExecute(r ReviewAPIGetReviewsByListingRequest) (*ListingReviews, *http.Response, error)

	/*
			GetReviewsByShop Method for GetReviewsByShop

			<div class="wt-display-flex-xs wt-align-items-center wt-mt-xs-2 wt-mb-xs-3"><span class="wt-badge wt-badge--notificationPrimary wt-bg-slime-tint wt-mr-xs-2">General Release</span><a class="wt-text-link" href="https://github.com/etsy/open-api/discussions" target="_blank" rel="noopener noreferrer">Report bug</a></div><div class="wt-display-flex-xs wt-align-items-center wt-mt-xs-2 wt-mb-xs-3"><p class="wt-text-body-01 banner-text">This endpoint is ready for production use.</p></div>

		Open API V3 to retrieve the reviews from a shop given its ID.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param shopId The unique positive non-zero numeric ID for an Etsy Shop.
			@return ReviewAPIGetReviewsByShopRequest
	*/
	GetReviewsByShop(ctx context.Context, shopId int64) ReviewAPIGetReviewsByShopRequest

	// GetReviewsByShopExecute executes the request
	//  @return TransactionReviews
	GetReviewsByShopExecute(r ReviewAPIGetReviewsByShopRequest) (*TransactionReviews, *http.Response, error)
}

// ReviewAPIService ReviewAPI service
type ReviewAPIService service

type ReviewAPIGetReviewsByListingRequest struct {
	ctx        context.Context
	ApiService ReviewAPI
	listingId  int64
	limit      *int64
	offset     *int64
	minCreated *int64
	maxCreated *int64
}

// The maximum number of results to return.
func (r ReviewAPIGetReviewsByListingRequest) Limit(limit int64) ReviewAPIGetReviewsByListingRequest {
	r.limit = &limit
	return r
}

// The number of records to skip before selecting the first result.
func (r ReviewAPIGetReviewsByListingRequest) Offset(offset int64) ReviewAPIGetReviewsByListingRequest {
	r.offset = &offset
	return r
}

// The earliest unix timestamp for when a record was created.
func (r ReviewAPIGetReviewsByListingRequest) MinCreated(minCreated int64) ReviewAPIGetReviewsByListingRequest {
	r.minCreated = &minCreated
	return r
}

// The latest unix timestamp for when a record was created.
func (r ReviewAPIGetReviewsByListingRequest) MaxCreated(maxCreated int64) ReviewAPIGetReviewsByListingRequest {
	r.maxCreated = &maxCreated
	return r
}

func (r ReviewAPIGetReviewsByListingRequest) Execute() (*ListingReviews, *http.Response, error) {
	return r.ApiService.GetReviewsByListingExecute(r)
}

/*
GetReviewsByListing Method for GetReviewsByListing

<div class="wt-display-flex-xs wt-align-items-center wt-mt-xs-2 wt-mb-xs-3"><span class="wt-badge wt-badge--notificationPrimary wt-bg-slime-tint wt-mr-xs-2">General Release</span><a class="wt-text-link" href="https://github.com/etsy/open-api/discussions" target="_blank" rel="noopener noreferrer">Report bug</a></div><div class="wt-display-flex-xs wt-align-items-center wt-mt-xs-2 wt-mb-xs-3"><p class="wt-text-body-01 banner-text">This endpoint is ready for production use.</p></div>

Open API V3 to retrieve the reviews for a listing given its ID.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param listingId The numeric ID for the [listing](/documentation/reference#tag/ShopListing) associated to this transaction.
	@return ReviewAPIGetReviewsByListingRequest
*/
func (a *ReviewAPIService) GetReviewsByListing(ctx context.Context, listingId int64) ReviewAPIGetReviewsByListingRequest {
	return ReviewAPIGetReviewsByListingRequest{
		ApiService: a,
		ctx:        ctx,
		listingId:  listingId,
	}
}

// Execute executes the request
//
//	@return ListingReviews
func (a *ReviewAPIService) GetReviewsByListingExecute(r ReviewAPIGetReviewsByListingRequest) (*ListingReviews, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListingReviews
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReviewAPIService.GetReviewsByListing")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/application/listings/{listing_id}/reviews"
	localVarPath = strings.Replace(localVarPath, "{"+"listing_id"+"}", url.PathEscape(parameterValueToString(r.listingId, "listingId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.listingId < 1 {
		return localVarReturnValue, nil, reportError("listingId must be greater than 1")
	}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	} else {
		var defaultValue int64 = 25
		r.limit = &defaultValue
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	} else {
		var defaultValue int64 = 0
		r.offset = &defaultValue
	}
	if r.minCreated != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "min_created", r.minCreated, "form", "")
	}
	if r.maxCreated != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "max_created", r.maxCreated, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["x-api-key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorSchema
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorSchema
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorSchema
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ReviewAPIGetReviewsByShopRequest struct {
	ctx        context.Context
	ApiService ReviewAPI
	shopId     int64
	limit      *int64
	offset     *int64
	minCreated *int64
	maxCreated *int64
}

// The maximum number of results to return.
func (r ReviewAPIGetReviewsByShopRequest) Limit(limit int64) ReviewAPIGetReviewsByShopRequest {
	r.limit = &limit
	return r
}

// The number of records to skip before selecting the first result.
func (r ReviewAPIGetReviewsByShopRequest) Offset(offset int64) ReviewAPIGetReviewsByShopRequest {
	r.offset = &offset
	return r
}

// The earliest unix timestamp for when a record was created.
func (r ReviewAPIGetReviewsByShopRequest) MinCreated(minCreated int64) ReviewAPIGetReviewsByShopRequest {
	r.minCreated = &minCreated
	return r
}

// The latest unix timestamp for when a record was created.
func (r ReviewAPIGetReviewsByShopRequest) MaxCreated(maxCreated int64) ReviewAPIGetReviewsByShopRequest {
	r.maxCreated = &maxCreated
	return r
}

func (r ReviewAPIGetReviewsByShopRequest) Execute() (*TransactionReviews, *http.Response, error) {
	return r.ApiService.GetReviewsByShopExecute(r)
}

/*
GetReviewsByShop Method for GetReviewsByShop

<div class="wt-display-flex-xs wt-align-items-center wt-mt-xs-2 wt-mb-xs-3"><span class="wt-badge wt-badge--notificationPrimary wt-bg-slime-tint wt-mr-xs-2">General Release</span><a class="wt-text-link" href="https://github.com/etsy/open-api/discussions" target="_blank" rel="noopener noreferrer">Report bug</a></div><div class="wt-display-flex-xs wt-align-items-center wt-mt-xs-2 wt-mb-xs-3"><p class="wt-text-body-01 banner-text">This endpoint is ready for production use.</p></div>

Open API V3 to retrieve the reviews from a shop given its ID.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param shopId The unique positive non-zero numeric ID for an Etsy Shop.
	@return ReviewAPIGetReviewsByShopRequest
*/
func (a *ReviewAPIService) GetReviewsByShop(ctx context.Context, shopId int64) ReviewAPIGetReviewsByShopRequest {
	return ReviewAPIGetReviewsByShopRequest{
		ApiService: a,
		ctx:        ctx,
		shopId:     shopId,
	}
}

// Execute executes the request
//
//	@return TransactionReviews
func (a *ReviewAPIService) GetReviewsByShopExecute(r ReviewAPIGetReviewsByShopRequest) (*TransactionReviews, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TransactionReviews
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReviewAPIService.GetReviewsByShop")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/application/shops/{shop_id}/reviews"
	localVarPath = strings.Replace(localVarPath, "{"+"shop_id"+"}", url.PathEscape(parameterValueToString(r.shopId, "shopId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.shopId < 1 {
		return localVarReturnValue, nil, reportError("shopId must be greater than 1")
	}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	} else {
		var defaultValue int64 = 25
		r.limit = &defaultValue
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	} else {
		var defaultValue int64 = 0
		r.offset = &defaultValue
	}
	if r.minCreated != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "min_created", r.minCreated, "form", "")
	}
	if r.maxCreated != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "max_created", r.maxCreated, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["x-api-key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorSchema
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorSchema
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorSchema
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
