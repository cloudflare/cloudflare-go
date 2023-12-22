// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountSecondaryDNSTsigService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountSecondaryDNSTsigService] method instead.
type AccountSecondaryDNSTsigService struct {
	Options []option.RequestOption
}

// NewAccountSecondaryDNSTsigService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountSecondaryDNSTsigService(opts ...option.RequestOption) (r *AccountSecondaryDNSTsigService) {
	r = &AccountSecondaryDNSTsigService{}
	r.Options = opts
	return
}

// Get TSIG.
func (r *AccountSecondaryDNSTsigService) Get(ctx context.Context, accountIdentifier interface{}, identifier interface{}, opts ...option.RequestOption) (res *SingleResponseWdx0OxVv, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/secondary_dns/tsigs/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify TSIG.
func (r *AccountSecondaryDNSTsigService) Update(ctx context.Context, accountIdentifier interface{}, identifier interface{}, body AccountSecondaryDNSTsigUpdateParams, opts ...option.RequestOption) (res *SingleResponseWdx0OxVv, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/secondary_dns/tsigs/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete TSIG.
func (r *AccountSecondaryDNSTsigService) Delete(ctx context.Context, accountIdentifier interface{}, identifier interface{}, opts ...option.RequestOption) (res *SchemasIDResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/secondary_dns/tsigs/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create TSIG.
func (r *AccountSecondaryDNSTsigService) SecondaryDNSTsigNewTsig(ctx context.Context, accountIdentifier interface{}, body AccountSecondaryDNSTsigSecondaryDNSTsigNewTsigParams, opts ...option.RequestOption) (res *SingleResponseWdx0OxVv, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/secondary_dns/tsigs", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List TSIGs.
func (r *AccountSecondaryDNSTsigService) SecondaryDNSTsigListTsiGs(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *ResponseCollectionLevvNosb, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/secondary_dns/tsigs", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ResponseCollectionLevvNosb struct {
	Errors     []ResponseCollectionLevvNosbError    `json:"errors"`
	Messages   []ResponseCollectionLevvNosbMessage  `json:"messages"`
	Result     []ResponseCollectionLevvNosbResult   `json:"result"`
	ResultInfo ResponseCollectionLevvNosbResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ResponseCollectionLevvNosbSuccess `json:"success"`
	JSON    responseCollectionLevvNosbJSON    `json:"-"`
}

// responseCollectionLevvNosbJSON contains the JSON metadata for the struct
// [ResponseCollectionLevvNosb]
type responseCollectionLevvNosbJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionLevvNosb) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionLevvNosbError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    responseCollectionLevvNosbErrorJSON `json:"-"`
}

// responseCollectionLevvNosbErrorJSON contains the JSON metadata for the struct
// [ResponseCollectionLevvNosbError]
type responseCollectionLevvNosbErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionLevvNosbError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionLevvNosbMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    responseCollectionLevvNosbMessageJSON `json:"-"`
}

// responseCollectionLevvNosbMessageJSON contains the JSON metadata for the struct
// [ResponseCollectionLevvNosbMessage]
type responseCollectionLevvNosbMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionLevvNosbMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionLevvNosbResult struct {
	ID interface{} `json:"id,required"`
	// TSIG algorithm.
	Algo string `json:"algo,required"`
	// TSIG key name.
	Name string `json:"name,required"`
	// TSIG secret.
	Secret string                               `json:"secret,required"`
	JSON   responseCollectionLevvNosbResultJSON `json:"-"`
}

// responseCollectionLevvNosbResultJSON contains the JSON metadata for the struct
// [ResponseCollectionLevvNosbResult]
type responseCollectionLevvNosbResultJSON struct {
	ID          apijson.Field
	Algo        apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionLevvNosbResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionLevvNosbResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                  `json:"total_count"`
	JSON       responseCollectionLevvNosbResultInfoJSON `json:"-"`
}

// responseCollectionLevvNosbResultInfoJSON contains the JSON metadata for the
// struct [ResponseCollectionLevvNosbResultInfo]
type responseCollectionLevvNosbResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionLevvNosbResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ResponseCollectionLevvNosbSuccess bool

const (
	ResponseCollectionLevvNosbSuccessTrue ResponseCollectionLevvNosbSuccess = true
)

type SchemasIDResponse struct {
	Errors   []SchemasIDResponseError   `json:"errors"`
	Messages []SchemasIDResponseMessage `json:"messages"`
	Result   SchemasIDResponseResult    `json:"result"`
	// Whether the API call was successful
	Success SchemasIDResponseSuccess `json:"success"`
	JSON    schemasIDResponseJSON    `json:"-"`
}

// schemasIDResponseJSON contains the JSON metadata for the struct
// [SchemasIDResponse]
type schemasIDResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasIDResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasIDResponseError struct {
	Code    int64                      `json:"code,required"`
	Message string                     `json:"message,required"`
	JSON    schemasIDResponseErrorJSON `json:"-"`
}

// schemasIDResponseErrorJSON contains the JSON metadata for the struct
// [SchemasIDResponseError]
type schemasIDResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasIDResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasIDResponseMessage struct {
	Code    int64                        `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    schemasIDResponseMessageJSON `json:"-"`
}

// schemasIDResponseMessageJSON contains the JSON metadata for the struct
// [SchemasIDResponseMessage]
type schemasIDResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasIDResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasIDResponseResult struct {
	ID   interface{}                 `json:"id"`
	JSON schemasIDResponseResultJSON `json:"-"`
}

// schemasIDResponseResultJSON contains the JSON metadata for the struct
// [SchemasIDResponseResult]
type schemasIDResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasIDResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SchemasIDResponseSuccess bool

const (
	SchemasIDResponseSuccessTrue SchemasIDResponseSuccess = true
)

type SingleResponseWdx0OxVv struct {
	Errors   []SingleResponseWdx0OxVvError   `json:"errors"`
	Messages []SingleResponseWdx0OxVvMessage `json:"messages"`
	Result   SingleResponseWdx0OxVvResult    `json:"result"`
	// Whether the API call was successful
	Success SingleResponseWdx0OxVvSuccess `json:"success"`
	JSON    singleResponseWdx0OxVvJSON    `json:"-"`
}

// singleResponseWdx0OxVvJSON contains the JSON metadata for the struct
// [SingleResponseWdx0OxVv]
type singleResponseWdx0OxVvJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseWdx0OxVv) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponseWdx0OxVvError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    singleResponseWdx0OxVvErrorJSON `json:"-"`
}

// singleResponseWdx0OxVvErrorJSON contains the JSON metadata for the struct
// [SingleResponseWdx0OxVvError]
type singleResponseWdx0OxVvErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseWdx0OxVvError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponseWdx0OxVvMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    singleResponseWdx0OxVvMessageJSON `json:"-"`
}

// singleResponseWdx0OxVvMessageJSON contains the JSON metadata for the struct
// [SingleResponseWdx0OxVvMessage]
type singleResponseWdx0OxVvMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseWdx0OxVvMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponseWdx0OxVvResult struct {
	ID interface{} `json:"id,required"`
	// TSIG algorithm.
	Algo string `json:"algo,required"`
	// TSIG key name.
	Name string `json:"name,required"`
	// TSIG secret.
	Secret string                           `json:"secret,required"`
	JSON   singleResponseWdx0OxVvResultJSON `json:"-"`
}

// singleResponseWdx0OxVvResultJSON contains the JSON metadata for the struct
// [SingleResponseWdx0OxVvResult]
type singleResponseWdx0OxVvResultJSON struct {
	ID          apijson.Field
	Algo        apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseWdx0OxVvResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SingleResponseWdx0OxVvSuccess bool

const (
	SingleResponseWdx0OxVvSuccessTrue SingleResponseWdx0OxVvSuccess = true
)

type AccountSecondaryDNSTsigUpdateParams struct {
	// TSIG algorithm.
	Algo param.Field[string] `json:"algo,required"`
	// TSIG key name.
	Name param.Field[string] `json:"name,required"`
	// TSIG secret.
	Secret param.Field[string] `json:"secret,required"`
}

func (r AccountSecondaryDNSTsigUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountSecondaryDNSTsigSecondaryDNSTsigNewTsigParams struct {
	// TSIG algorithm.
	Algo param.Field[string] `json:"algo,required"`
	// TSIG key name.
	Name param.Field[string] `json:"name,required"`
	// TSIG secret.
	Secret param.Field[string] `json:"secret,required"`
}

func (r AccountSecondaryDNSTsigSecondaryDNSTsigNewTsigParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
