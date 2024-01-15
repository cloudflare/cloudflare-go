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
func (r *AccountWorkerScriptTailService) Delete(ctx context.Context, accountIdentifier string, scriptName string, id string, opts ...option.RequestOption) (res *AccountWorkerScriptTailDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/tails/%s", accountIdentifier, scriptName, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get list of tails currently deployed on a Worker.
func (r *AccountWorkerScriptTailService) WorkerTailLogsListTails(ctx context.Context, accountIdentifier string, scriptName string, opts ...option.RequestOption) (res *AccountWorkerScriptTailWorkerTailLogsListTailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/tails", accountIdentifier, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Starts a tail that receives logs and exception from a Worker.
func (r *AccountWorkerScriptTailService) WorkerTailLogsStartTail(ctx context.Context, accountIdentifier string, scriptName string, opts ...option.RequestOption) (res *AccountWorkerScriptTailWorkerTailLogsStartTailResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/tails", accountIdentifier, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type AccountWorkerScriptTailDeleteResponse struct {
	Errors   []AccountWorkerScriptTailDeleteResponseError   `json:"errors,required"`
	Messages []AccountWorkerScriptTailDeleteResponseMessage `json:"messages,required"`
	Result   AccountWorkerScriptTailDeleteResponseResult    `json:"result,required"`
	// Whether the API call was successful
	Success AccountWorkerScriptTailDeleteResponseSuccess `json:"success,required"`
	JSON    accountWorkerScriptTailDeleteResponseJSON    `json:"-"`
}

// accountWorkerScriptTailDeleteResponseJSON contains the JSON metadata for the
// struct [AccountWorkerScriptTailDeleteResponse]
type accountWorkerScriptTailDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptTailDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptTailDeleteResponseError struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountWorkerScriptTailDeleteResponseErrorJSON `json:"-"`
}

// accountWorkerScriptTailDeleteResponseErrorJSON contains the JSON metadata for
// the struct [AccountWorkerScriptTailDeleteResponseError]
type accountWorkerScriptTailDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptTailDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptTailDeleteResponseMessage struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accountWorkerScriptTailDeleteResponseMessageJSON `json:"-"`
}

// accountWorkerScriptTailDeleteResponseMessageJSON contains the JSON metadata for
// the struct [AccountWorkerScriptTailDeleteResponseMessage]
type accountWorkerScriptTailDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptTailDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [AccountWorkerScriptTailDeleteResponseResultUnknown],
// [AccountWorkerScriptTailDeleteResponseResultArray] or [shared.UnionString].
type AccountWorkerScriptTailDeleteResponseResult interface {
	ImplementsAccountWorkerScriptTailDeleteResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountWorkerScriptTailDeleteResponseResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccountWorkerScriptTailDeleteResponseResultArray []interface{}

func (r AccountWorkerScriptTailDeleteResponseResultArray) ImplementsAccountWorkerScriptTailDeleteResponseResult() {
}

// Whether the API call was successful
type AccountWorkerScriptTailDeleteResponseSuccess bool

const (
	AccountWorkerScriptTailDeleteResponseSuccessTrue AccountWorkerScriptTailDeleteResponseSuccess = true
)

type AccountWorkerScriptTailWorkerTailLogsListTailsResponse struct {
	Errors   []AccountWorkerScriptTailWorkerTailLogsListTailsResponseError   `json:"errors"`
	Messages []AccountWorkerScriptTailWorkerTailLogsListTailsResponseMessage `json:"messages"`
	Result   AccountWorkerScriptTailWorkerTailLogsListTailsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountWorkerScriptTailWorkerTailLogsListTailsResponseSuccess `json:"success"`
	JSON    accountWorkerScriptTailWorkerTailLogsListTailsResponseJSON    `json:"-"`
}

// accountWorkerScriptTailWorkerTailLogsListTailsResponseJSON contains the JSON
// metadata for the struct [AccountWorkerScriptTailWorkerTailLogsListTailsResponse]
type accountWorkerScriptTailWorkerTailLogsListTailsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptTailWorkerTailLogsListTailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptTailWorkerTailLogsListTailsResponseError struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    accountWorkerScriptTailWorkerTailLogsListTailsResponseErrorJSON `json:"-"`
}

// accountWorkerScriptTailWorkerTailLogsListTailsResponseErrorJSON contains the
// JSON metadata for the struct
// [AccountWorkerScriptTailWorkerTailLogsListTailsResponseError]
type accountWorkerScriptTailWorkerTailLogsListTailsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptTailWorkerTailLogsListTailsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptTailWorkerTailLogsListTailsResponseMessage struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    accountWorkerScriptTailWorkerTailLogsListTailsResponseMessageJSON `json:"-"`
}

// accountWorkerScriptTailWorkerTailLogsListTailsResponseMessageJSON contains the
// JSON metadata for the struct
// [AccountWorkerScriptTailWorkerTailLogsListTailsResponseMessage]
type accountWorkerScriptTailWorkerTailLogsListTailsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptTailWorkerTailLogsListTailsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptTailWorkerTailLogsListTailsResponseResult struct {
	ID        interface{}                                                      `json:"id"`
	ExpiresAt interface{}                                                      `json:"expires_at"`
	URL       interface{}                                                      `json:"url"`
	JSON      accountWorkerScriptTailWorkerTailLogsListTailsResponseResultJSON `json:"-"`
}

// accountWorkerScriptTailWorkerTailLogsListTailsResponseResultJSON contains the
// JSON metadata for the struct
// [AccountWorkerScriptTailWorkerTailLogsListTailsResponseResult]
type accountWorkerScriptTailWorkerTailLogsListTailsResponseResultJSON struct {
	ID          apijson.Field
	ExpiresAt   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptTailWorkerTailLogsListTailsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountWorkerScriptTailWorkerTailLogsListTailsResponseSuccess bool

const (
	AccountWorkerScriptTailWorkerTailLogsListTailsResponseSuccessTrue AccountWorkerScriptTailWorkerTailLogsListTailsResponseSuccess = true
)

type AccountWorkerScriptTailWorkerTailLogsStartTailResponse struct {
	Errors   []AccountWorkerScriptTailWorkerTailLogsStartTailResponseError   `json:"errors"`
	Messages []AccountWorkerScriptTailWorkerTailLogsStartTailResponseMessage `json:"messages"`
	Result   AccountWorkerScriptTailWorkerTailLogsStartTailResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountWorkerScriptTailWorkerTailLogsStartTailResponseSuccess `json:"success"`
	JSON    accountWorkerScriptTailWorkerTailLogsStartTailResponseJSON    `json:"-"`
}

// accountWorkerScriptTailWorkerTailLogsStartTailResponseJSON contains the JSON
// metadata for the struct [AccountWorkerScriptTailWorkerTailLogsStartTailResponse]
type accountWorkerScriptTailWorkerTailLogsStartTailResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptTailWorkerTailLogsStartTailResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptTailWorkerTailLogsStartTailResponseError struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    accountWorkerScriptTailWorkerTailLogsStartTailResponseErrorJSON `json:"-"`
}

// accountWorkerScriptTailWorkerTailLogsStartTailResponseErrorJSON contains the
// JSON metadata for the struct
// [AccountWorkerScriptTailWorkerTailLogsStartTailResponseError]
type accountWorkerScriptTailWorkerTailLogsStartTailResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptTailWorkerTailLogsStartTailResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptTailWorkerTailLogsStartTailResponseMessage struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    accountWorkerScriptTailWorkerTailLogsStartTailResponseMessageJSON `json:"-"`
}

// accountWorkerScriptTailWorkerTailLogsStartTailResponseMessageJSON contains the
// JSON metadata for the struct
// [AccountWorkerScriptTailWorkerTailLogsStartTailResponseMessage]
type accountWorkerScriptTailWorkerTailLogsStartTailResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptTailWorkerTailLogsStartTailResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptTailWorkerTailLogsStartTailResponseResult struct {
	ID        interface{}                                                      `json:"id"`
	ExpiresAt interface{}                                                      `json:"expires_at"`
	URL       interface{}                                                      `json:"url"`
	JSON      accountWorkerScriptTailWorkerTailLogsStartTailResponseResultJSON `json:"-"`
}

// accountWorkerScriptTailWorkerTailLogsStartTailResponseResultJSON contains the
// JSON metadata for the struct
// [AccountWorkerScriptTailWorkerTailLogsStartTailResponseResult]
type accountWorkerScriptTailWorkerTailLogsStartTailResponseResultJSON struct {
	ID          apijson.Field
	ExpiresAt   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptTailWorkerTailLogsStartTailResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountWorkerScriptTailWorkerTailLogsStartTailResponseSuccess bool

const (
	AccountWorkerScriptTailWorkerTailLogsStartTailResponseSuccessTrue AccountWorkerScriptTailWorkerTailLogsStartTailResponseSuccess = true
)
