// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// AccountWorkerScriptTailService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountWorkerScriptTailService] method instead.
type AccountWorkerScriptTailService struct {
	Options []option.RequestOption
}

// NewAccountWorkerScriptTailService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountWorkerScriptTailService(opts ...option.RequestOption) (r *AccountWorkerScriptTailService) {
	r = &AccountWorkerScriptTailService{}
	r.Options = opts
	return
}

// Deletes a tail from a Worker.
func (r *AccountWorkerScriptTailService) Delete(ctx context.Context, accountIdentifier string, scriptName string, id string, opts ...option.RequestOption) (res *APIResponseCommon, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/tails/%s", accountIdentifier, scriptName, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get list of tails currently deployed on a worker.
func (r *AccountWorkerScriptTailService) WorkerTailLogsListTails(ctx context.Context, accountIdentifier string, scriptName string, opts ...option.RequestOption) (res *TailResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/tails", accountIdentifier, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Starts a tail that receives logs and exception from a Worker.
func (r *AccountWorkerScriptTailService) WorkerTailLogsStartTail(ctx context.Context, accountIdentifier string, scriptName string, opts ...option.RequestOption) (res *TailResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/tails", accountIdentifier, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type APIResponseCommon struct {
	Errors   []APIResponseCommonError   `json:"errors,required"`
	Messages []APIResponseCommonMessage `json:"messages,required"`
	Result   APIResponseCommonResult    `json:"result,required"`
	// Whether the API call was successful
	Success APIResponseCommonSuccess `json:"success,required"`
	JSON    apiResponseCommonJSON    `json:"-"`
}

// apiResponseCommonJSON contains the JSON metadata for the struct
// [APIResponseCommon]
type apiResponseCommonJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseCommon) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type APIResponseCommonError struct {
	Code    int64                      `json:"code,required"`
	Message string                     `json:"message,required"`
	JSON    apiResponseCommonErrorJSON `json:"-"`
}

// apiResponseCommonErrorJSON contains the JSON metadata for the struct
// [APIResponseCommonError]
type apiResponseCommonErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseCommonError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type APIResponseCommonMessage struct {
	Code    int64                        `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    apiResponseCommonMessageJSON `json:"-"`
}

// apiResponseCommonMessageJSON contains the JSON metadata for the struct
// [APIResponseCommonMessage]
type apiResponseCommonMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseCommonMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [APIResponseCommonResultObject],
// [APIResponseCommonResultObject] or [shared.UnionString].
type APIResponseCommonResult interface {
	ImplementsAPIResponseCommonResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*APIResponseCommonResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Whether the API call was successful
type APIResponseCommonSuccess bool

const (
	APIResponseCommonSuccessTrue APIResponseCommonSuccess = true
)

type TailResponse struct {
	Errors   []TailResponseError   `json:"errors"`
	Messages []TailResponseMessage `json:"messages"`
	Result   TailResponseResult    `json:"result"`
	// Whether the API call was successful
	Success TailResponseSuccess `json:"success"`
	JSON    tailResponseJSON    `json:"-"`
}

// tailResponseJSON contains the JSON metadata for the struct [TailResponse]
type tailResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TailResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TailResponseError struct {
	Code    int64                 `json:"code,required"`
	Message string                `json:"message,required"`
	JSON    tailResponseErrorJSON `json:"-"`
}

// tailResponseErrorJSON contains the JSON metadata for the struct
// [TailResponseError]
type tailResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TailResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TailResponseMessage struct {
	Code    int64                   `json:"code,required"`
	Message string                  `json:"message,required"`
	JSON    tailResponseMessageJSON `json:"-"`
}

// tailResponseMessageJSON contains the JSON metadata for the struct
// [TailResponseMessage]
type tailResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TailResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TailResponseResult struct {
	ID        interface{}            `json:"id"`
	ExpiresAt interface{}            `json:"expires_at"`
	URL       interface{}            `json:"url"`
	JSON      tailResponseResultJSON `json:"-"`
}

// tailResponseResultJSON contains the JSON metadata for the struct
// [TailResponseResult]
type tailResponseResultJSON struct {
	ID          apijson.Field
	ExpiresAt   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TailResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TailResponseSuccess bool

const (
	TailResponseSuccessTrue TailResponseSuccess = true
)
