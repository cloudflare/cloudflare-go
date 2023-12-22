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

// AccountLoadBalancerPreviewService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountLoadBalancerPreviewService] method instead.
type AccountLoadBalancerPreviewService struct {
	Options []option.RequestOption
}

// NewAccountLoadBalancerPreviewService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountLoadBalancerPreviewService(opts ...option.RequestOption) (r *AccountLoadBalancerPreviewService) {
	r = &AccountLoadBalancerPreviewService{}
	r.Options = opts
	return
}

// Get the result of a previous preview operation using the provided preview_id.
func (r *AccountLoadBalancerPreviewService) Get(ctx context.Context, accountIdentifier string, previewID interface{}, opts ...option.RequestOption) (res *PreviewResultResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/preview/%v", accountIdentifier, previewID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PreviewResultResponse struct {
	Errors   []PreviewResultResponseError   `json:"errors"`
	Messages []PreviewResultResponseMessage `json:"messages"`
	Result   PreviewResultResponseResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success PreviewResultResponseSuccess `json:"success"`
	JSON    previewResultResponseJSON    `json:"-"`
}

// previewResultResponseJSON contains the JSON metadata for the struct
// [PreviewResultResponse]
type previewResultResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreviewResultResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PreviewResultResponseError struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    previewResultResponseErrorJSON `json:"-"`
}

// previewResultResponseErrorJSON contains the JSON metadata for the struct
// [PreviewResultResponseError]
type previewResultResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreviewResultResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PreviewResultResponseMessage struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    previewResultResponseMessageJSON `json:"-"`
}

// previewResultResponseMessageJSON contains the JSON metadata for the struct
// [PreviewResultResponseMessage]
type previewResultResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreviewResultResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [PreviewResultResponseResultObject] or [shared.UnionString].
type PreviewResultResponseResult interface {
	ImplementsPreviewResultResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PreviewResultResponseResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Whether the API call was successful
type PreviewResultResponseSuccess bool

const (
	PreviewResultResponseSuccessTrue PreviewResultResponseSuccess = true
)
