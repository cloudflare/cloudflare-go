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

// AccountWorkerDeploymentByScriptService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountWorkerDeploymentByScriptService] method instead.
type AccountWorkerDeploymentByScriptService struct {
	Options []option.RequestOption
	Details *AccountWorkerDeploymentByScriptDetailService
}

// NewAccountWorkerDeploymentByScriptService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountWorkerDeploymentByScriptService(opts ...option.RequestOption) (r *AccountWorkerDeploymentByScriptService) {
	r = &AccountWorkerDeploymentByScriptService{}
	r.Options = opts
	r.Details = NewAccountWorkerDeploymentByScriptDetailService(opts...)
	return
}

// List Deployments
func (r *AccountWorkerDeploymentByScriptService) WorkerDeploymentsListDeployments(ctx context.Context, accountIdentifier string, scriptIdentifier string, opts ...option.RequestOption) (res *AccountWorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/deployments/by-script/%s", accountIdentifier, scriptIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountWorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponse struct {
	Errors   []AccountWorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseError   `json:"errors"`
	Items    []interface{}                                                                    `json:"items"`
	Latest   interface{}                                                                      `json:"latest"`
	Messages []AccountWorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseMessage `json:"messages"`
	Result   AccountWorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountWorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseSuccess `json:"success"`
	JSON    accountWorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseJSON    `json:"-"`
}

// accountWorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseJSON
// contains the JSON metadata for the struct
// [AccountWorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponse]
type accountWorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseJSON struct {
	Errors      apijson.Field
	Items       apijson.Field
	Latest      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseError struct {
	Code    int64                                                                            `json:"code,required"`
	Message string                                                                           `json:"message,required"`
	JSON    accountWorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseErrorJSON `json:"-"`
}

// accountWorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountWorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseError]
type accountWorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseMessage struct {
	Code    int64                                                                              `json:"code,required"`
	Message string                                                                             `json:"message,required"`
	JSON    accountWorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseMessageJSON `json:"-"`
}

// accountWorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountWorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseMessage]
type accountWorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [AccountWorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseResultUnknown],
// [AccountWorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseResultArray]
// or [shared.UnionString].
type AccountWorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseResult interface {
	ImplementsAccountWorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountWorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccountWorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseResultArray []interface{}

func (r AccountWorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseResultArray) ImplementsAccountWorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseResult() {
}

// Whether the API call was successful
type AccountWorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseSuccess bool

const (
	AccountWorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseSuccessTrue AccountWorkerDeploymentByScriptWorkerDeploymentsListDeploymentsResponseSuccess = true
)
