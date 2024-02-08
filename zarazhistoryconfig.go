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
	var env ZarazHistoryConfigGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/zaraz/history/configs", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZarazHistoryConfigGetResponse map[string]ZarazHistoryConfigGetResponse

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

type ZarazHistoryConfigGetResponseEnvelope struct {
	Errors   []ZarazHistoryConfigGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZarazHistoryConfigGetResponseEnvelopeMessages `json:"messages,required"`
	// Object where keys are numericc onfiguration IDs
	Result ZarazHistoryConfigGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success bool                                      `json:"success,required"`
	JSON    zarazHistoryConfigGetResponseEnvelopeJSON `json:"-"`
}

// zarazHistoryConfigGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZarazHistoryConfigGetResponseEnvelope]
type zarazHistoryConfigGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryConfigGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryConfigGetResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zarazHistoryConfigGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zarazHistoryConfigGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZarazHistoryConfigGetResponseEnvelopeErrors]
type zarazHistoryConfigGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryConfigGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryConfigGetResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zarazHistoryConfigGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zarazHistoryConfigGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZarazHistoryConfigGetResponseEnvelopeMessages]
type zarazHistoryConfigGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryConfigGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
