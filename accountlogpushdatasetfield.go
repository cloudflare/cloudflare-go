// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountLogpushDatasetFieldService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountLogpushDatasetFieldService] method instead.
type AccountLogpushDatasetFieldService struct {
	Options []option.RequestOption
}

// NewAccountLogpushDatasetFieldService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountLogpushDatasetFieldService(opts ...option.RequestOption) (r *AccountLogpushDatasetFieldService) {
	r = &AccountLogpushDatasetFieldService{}
	r.Options = opts
	return
}

// Lists all fields available for a dataset. The response result is an object with
// key-value pairs, where keys are field names, and values are descriptions.
func (r *AccountLogpushDatasetFieldService) GetAccountsAccountIdentifierLogpushDatasetsDatasetFields(ctx context.Context, accountIdentifier string, dataset string, opts ...option.RequestOption) (res *LogpushFieldResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/logpush/datasets/%s/fields", accountIdentifier, dataset)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type LogpushFieldResponseCollection struct {
	Errors   []LogpushFieldResponseCollectionError   `json:"errors"`
	Messages []LogpushFieldResponseCollectionMessage `json:"messages"`
	Result   interface{}                             `json:"result"`
	// Whether the API call was successful
	Success LogpushFieldResponseCollectionSuccess `json:"success"`
	JSON    logpushFieldResponseCollectionJSON    `json:"-"`
}

// logpushFieldResponseCollectionJSON contains the JSON metadata for the struct
// [LogpushFieldResponseCollection]
type logpushFieldResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushFieldResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushFieldResponseCollectionError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    logpushFieldResponseCollectionErrorJSON `json:"-"`
}

// logpushFieldResponseCollectionErrorJSON contains the JSON metadata for the
// struct [LogpushFieldResponseCollectionError]
type logpushFieldResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushFieldResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushFieldResponseCollectionMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    logpushFieldResponseCollectionMessageJSON `json:"-"`
}

// logpushFieldResponseCollectionMessageJSON contains the JSON metadata for the
// struct [LogpushFieldResponseCollectionMessage]
type logpushFieldResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushFieldResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LogpushFieldResponseCollectionSuccess bool

const (
	LogpushFieldResponseCollectionSuccessTrue LogpushFieldResponseCollectionSuccess = true
)
