// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZarazHistoryConfigService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZarazHistoryConfigService] method
// instead.
type ZarazHistoryConfigService struct {
	Options []option.RequestOption
}

// NewZarazHistoryConfigService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZarazHistoryConfigService(opts ...option.RequestOption) (r *ZarazHistoryConfigService) {
	r = &ZarazHistoryConfigService{}
	r.Options = opts
	return
}

// Gets a history of published Zaraz configurations by ID(s) for a zone.
func (r *ZarazHistoryConfigService) Get(ctx context.Context, zoneID string, query ZarazHistoryConfigGetParams, opts ...option.RequestOption) (res *ZarazHistoryConfigGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/zaraz/history/configs", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ZarazHistoryConfigGetResponse struct {
	Errors   []ZarazHistoryConfigGetResponseError   `json:"errors"`
	Messages []ZarazHistoryConfigGetResponseMessage `json:"messages"`
	// Object where keys are numericc onfiguration IDs
	Result interface{} `json:"result"`
	// Whether the API call was successful
	Success bool                              `json:"success"`
	JSON    zarazHistoryConfigGetResponseJSON `json:"-"`
}

// zarazHistoryConfigGetResponseJSON contains the JSON metadata for the struct
// [ZarazHistoryConfigGetResponse]
type zarazHistoryConfigGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryConfigGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryConfigGetResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    zarazHistoryConfigGetResponseErrorJSON `json:"-"`
}

// zarazHistoryConfigGetResponseErrorJSON contains the JSON metadata for the struct
// [ZarazHistoryConfigGetResponseError]
type zarazHistoryConfigGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryConfigGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryConfigGetResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    zarazHistoryConfigGetResponseMessageJSON `json:"-"`
}

// zarazHistoryConfigGetResponseMessageJSON contains the JSON metadata for the
// struct [ZarazHistoryConfigGetResponseMessage]
type zarazHistoryConfigGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryConfigGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryConfigGetParams struct {
	// Comma separated list of Zaraz configuration IDs
	IDs param.Field[[]int64] `query:"ids,required"`
}

// URLQuery serializes [ZarazHistoryConfigGetParams]'s query parameters as
// `url.Values`.
func (r ZarazHistoryConfigGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
