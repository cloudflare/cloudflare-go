// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package snippets

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
	"github.com/tidwall/gjson"
)

// SnippetService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewSnippetService] method instead.
type SnippetService struct {
	Options []option.RequestOption
	Content *ContentService
	Rules   *RuleService
}

// NewSnippetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSnippetService(opts ...option.RequestOption) (r *SnippetService) {
	r = &SnippetService{}
	r.Options = opts
	r.Content = NewContentService(opts...)
	r.Rules = NewRuleService(opts...)
	return
}

// Put Snippet
func (r *SnippetService) Update(ctx context.Context, zoneIdentifier string, snippetName string, body SnippetUpdateParams, opts ...option.RequestOption) (res *Snippet, err error) {
	opts = append(r.Options[:], opts...)
	var env SnippetUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/snippets/%s", zoneIdentifier, snippetName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// All Snippets
func (r *SnippetService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *pagination.SinglePage[Snippet], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/snippets", zoneIdentifier)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// All Snippets
func (r *SnippetService) ListAutoPaging(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Snippet] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, zoneIdentifier, opts...))
}

// Delete Snippet
func (r *SnippetService) Delete(ctx context.Context, zoneIdentifier string, snippetName string, opts ...option.RequestOption) (res *SnippetDeleteResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env SnippetDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/snippets/%s", zoneIdentifier, snippetName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Snippet
func (r *SnippetService) Get(ctx context.Context, zoneIdentifier string, snippetName string, opts ...option.RequestOption) (res *Snippet, err error) {
	opts = append(r.Options[:], opts...)
	var env SnippetGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/snippets/%s", zoneIdentifier, snippetName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Snippet Information
type Snippet struct {
	// Creation time of the snippet
	CreatedOn string `json:"created_on"`
	// Modification time of the snippet
	ModifiedOn string `json:"modified_on"`
	// Snippet identifying name
	SnippetName string      `json:"snippet_name"`
	JSON        snippetJSON `json:"-"`
}

// snippetJSON contains the JSON metadata for the struct [Snippet]
type snippetJSON struct {
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	SnippetName apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Snippet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r snippetJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [snippets.SnippetDeleteResponseUnknown],
// [snippets.SnippetDeleteResponseArray] or [shared.UnionString].
type SnippetDeleteResponseUnion interface {
	ImplementsSnippetsSnippetDeleteResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SnippetDeleteResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SnippetDeleteResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type SnippetDeleteResponseArray []interface{}

func (r SnippetDeleteResponseArray) ImplementsSnippetsSnippetDeleteResponseUnion() {}

type SnippetUpdateParams struct {
	// Content files of uploaded snippet
	Files    param.Field[string]                      `json:"files"`
	Metadata param.Field[SnippetUpdateParamsMetadata] `json:"metadata"`
}

func (r SnippetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SnippetUpdateParamsMetadata struct {
	// Main module name of uploaded snippet
	MainModule param.Field[string] `json:"main_module"`
}

func (r SnippetUpdateParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SnippetUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Snippet Information
	Result Snippet `json:"result,required"`
	// Whether the API call was successful
	Success SnippetUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    snippetUpdateResponseEnvelopeJSON    `json:"-"`
}

// snippetUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SnippetUpdateResponseEnvelope]
type snippetUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r snippetUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SnippetUpdateResponseEnvelopeSuccess bool

const (
	SnippetUpdateResponseEnvelopeSuccessTrue SnippetUpdateResponseEnvelopeSuccess = true
)

func (r SnippetUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SnippetUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SnippetDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo      `json:"errors,required"`
	Messages []shared.ResponseInfo      `json:"messages,required"`
	Result   SnippetDeleteResponseUnion `json:"result,required"`
	// Whether the API call was successful
	Success SnippetDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    snippetDeleteResponseEnvelopeJSON    `json:"-"`
}

// snippetDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [SnippetDeleteResponseEnvelope]
type snippetDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r snippetDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SnippetDeleteResponseEnvelopeSuccess bool

const (
	SnippetDeleteResponseEnvelopeSuccessTrue SnippetDeleteResponseEnvelopeSuccess = true
)

func (r SnippetDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SnippetDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SnippetGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Snippet Information
	Result Snippet `json:"result,required"`
	// Whether the API call was successful
	Success SnippetGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    snippetGetResponseEnvelopeJSON    `json:"-"`
}

// snippetGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SnippetGetResponseEnvelope]
type snippetGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r snippetGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SnippetGetResponseEnvelopeSuccess bool

const (
	SnippetGetResponseEnvelopeSuccessTrue SnippetGetResponseEnvelopeSuccess = true
)

func (r SnippetGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SnippetGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
