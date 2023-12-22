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
func (r *AccountWorkerDeploymentByScriptService) WorkerDeploymentsListDeployments(ctx context.Context, accountIdentifier string, scriptIdentifier string, opts ...option.RequestOption) (res *DeploymentsListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/deployments/by-script/%s", accountIdentifier, scriptIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type DeploymentsListResponse struct {
	Errors   []DeploymentsListResponseError   `json:"errors"`
	Items    []interface{}                    `json:"items"`
	Latest   interface{}                      `json:"latest"`
	Messages []DeploymentsListResponseMessage `json:"messages"`
	Result   DeploymentsListResponseResult    `json:"result"`
	// Whether the API call was successful
	Success DeploymentsListResponseSuccess `json:"success"`
	JSON    deploymentsListResponseJSON    `json:"-"`
}

// deploymentsListResponseJSON contains the JSON metadata for the struct
// [DeploymentsListResponse]
type deploymentsListResponseJSON struct {
	Errors      apijson.Field
	Items       apijson.Field
	Latest      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentsListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeploymentsListResponseError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    deploymentsListResponseErrorJSON `json:"-"`
}

// deploymentsListResponseErrorJSON contains the JSON metadata for the struct
// [DeploymentsListResponseError]
type deploymentsListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentsListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeploymentsListResponseMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    deploymentsListResponseMessageJSON `json:"-"`
}

// deploymentsListResponseMessageJSON contains the JSON metadata for the struct
// [DeploymentsListResponseMessage]
type deploymentsListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentsListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [DeploymentsListResponseResultObject],
// [DeploymentsListResponseResultObject] or [shared.UnionString].
type DeploymentsListResponseResult interface {
	ImplementsDeploymentsListResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DeploymentsListResponseResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Whether the API call was successful
type DeploymentsListResponseSuccess bool

const (
	DeploymentsListResponseSuccessTrue DeploymentsListResponseSuccess = true
)
