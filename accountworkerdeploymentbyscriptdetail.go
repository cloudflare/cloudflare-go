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

// AccountWorkerDeploymentByScriptDetailService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountWorkerDeploymentByScriptDetailService] method instead.
type AccountWorkerDeploymentByScriptDetailService struct {
	Options []option.RequestOption
}

// NewAccountWorkerDeploymentByScriptDetailService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountWorkerDeploymentByScriptDetailService(opts ...option.RequestOption) (r *AccountWorkerDeploymentByScriptDetailService) {
	r = &AccountWorkerDeploymentByScriptDetailService{}
	r.Options = opts
	return
}

// Get Deployment Detail
func (r *AccountWorkerDeploymentByScriptDetailService) Get(ctx context.Context, accountIdentifier string, scriptIdentifier string, deploymentIdentifier string, opts ...option.RequestOption) (res *AccountWorkerDeploymentByScriptDetailGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/deployments/by-script/%s/detail/%s", accountIdentifier, scriptIdentifier, deploymentIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountWorkerDeploymentByScriptDetailGetResponse struct {
	ID        string                                                    `json:"id"`
	Errors    []AccountWorkerDeploymentByScriptDetailGetResponseError   `json:"errors"`
	Messages  []AccountWorkerDeploymentByScriptDetailGetResponseMessage `json:"messages"`
	Metadata  interface{}                                               `json:"metadata"`
	Number    float64                                                   `json:"number"`
	Resources interface{}                                               `json:"resources"`
	Result    AccountWorkerDeploymentByScriptDetailGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountWorkerDeploymentByScriptDetailGetResponseSuccess `json:"success"`
	JSON    accountWorkerDeploymentByScriptDetailGetResponseJSON    `json:"-"`
}

// accountWorkerDeploymentByScriptDetailGetResponseJSON contains the JSON metadata
// for the struct [AccountWorkerDeploymentByScriptDetailGetResponse]
type accountWorkerDeploymentByScriptDetailGetResponseJSON struct {
	ID          apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
	Metadata    apijson.Field
	Number      apijson.Field
	Resources   apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDeploymentByScriptDetailGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerDeploymentByScriptDetailGetResponseError struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    accountWorkerDeploymentByScriptDetailGetResponseErrorJSON `json:"-"`
}

// accountWorkerDeploymentByScriptDetailGetResponseErrorJSON contains the JSON
// metadata for the struct [AccountWorkerDeploymentByScriptDetailGetResponseError]
type accountWorkerDeploymentByScriptDetailGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDeploymentByScriptDetailGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerDeploymentByScriptDetailGetResponseMessage struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    accountWorkerDeploymentByScriptDetailGetResponseMessageJSON `json:"-"`
}

// accountWorkerDeploymentByScriptDetailGetResponseMessageJSON contains the JSON
// metadata for the struct
// [AccountWorkerDeploymentByScriptDetailGetResponseMessage]
type accountWorkerDeploymentByScriptDetailGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDeploymentByScriptDetailGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [AccountWorkerDeploymentByScriptDetailGetResponseResultUnknown],
// [AccountWorkerDeploymentByScriptDetailGetResponseResultArray] or
// [shared.UnionString].
type AccountWorkerDeploymentByScriptDetailGetResponseResult interface {
	ImplementsAccountWorkerDeploymentByScriptDetailGetResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountWorkerDeploymentByScriptDetailGetResponseResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccountWorkerDeploymentByScriptDetailGetResponseResultArray []interface{}

func (r AccountWorkerDeploymentByScriptDetailGetResponseResultArray) ImplementsAccountWorkerDeploymentByScriptDetailGetResponseResult() {
}

// Whether the API call was successful
type AccountWorkerDeploymentByScriptDetailGetResponseSuccess bool

const (
	AccountWorkerDeploymentByScriptDetailGetResponseSuccessTrue AccountWorkerDeploymentByScriptDetailGetResponseSuccess = true
)
