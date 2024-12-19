// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package snippets

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// SnippetService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSnippetService] method instead.
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
func (r *SnippetService) Update(ctx context.Context, snippetName string, params SnippetUpdateParams, opts ...option.RequestOption) (res *Snippet, err error) {
	var env SnippetUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if snippetName == "" {
		err = errors.New("missing required snippet_name parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/snippets/%s", params.ZoneID, snippetName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// All Snippets
func (r *SnippetService) List(ctx context.Context, query SnippetListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Snippet], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/snippets", query.ZoneID)
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
func (r *SnippetService) ListAutoPaging(ctx context.Context, query SnippetListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Snippet] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete Snippet
func (r *SnippetService) Delete(ctx context.Context, snippetName string, body SnippetDeleteParams, opts ...option.RequestOption) (res *SnippetDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if snippetName == "" {
		err = errors.New("missing required snippet_name parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/snippets/%s", body.ZoneID, snippetName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Snippet
func (r *SnippetService) Get(ctx context.Context, snippetName string, query SnippetGetParams, opts ...option.RequestOption) (res *Snippet, err error) {
	var env SnippetGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if snippetName == "" {
		err = errors.New("missing required snippet_name parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/snippets/%s", query.ZoneID, snippetName)
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

type SnippetDeleteResponse struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success SnippetDeleteResponseSuccess `json:"success,required"`
	JSON    snippetDeleteResponseJSON    `json:"-"`
}

// snippetDeleteResponseJSON contains the JSON metadata for the struct
// [SnippetDeleteResponse]
type snippetDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r snippetDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SnippetDeleteResponseSuccess bool

const (
	SnippetDeleteResponseSuccessTrue SnippetDeleteResponseSuccess = true
)

func (r SnippetDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case SnippetDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type SnippetUpdateParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Content files of uploaded snippet
	Files    param.Field[string]                      `json:"files"`
	Metadata param.Field[SnippetUpdateParamsMetadata] `json:"metadata"`
}

func (r SnippetUpdateParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
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
	// Whether the API call was successful
	Success SnippetUpdateResponseEnvelopeSuccess `json:"success,required"`
	// Snippet Information
	Result Snippet                           `json:"result"`
	JSON   snippetUpdateResponseEnvelopeJSON `json:"-"`
}

// snippetUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SnippetUpdateResponseEnvelope]
type snippetUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
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

type SnippetListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SnippetDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SnippetGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SnippetGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success SnippetGetResponseEnvelopeSuccess `json:"success,required"`
	// Snippet Information
	Result Snippet                        `json:"result"`
	JSON   snippetGetResponseEnvelopeJSON `json:"-"`
}

// snippetGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SnippetGetResponseEnvelope]
type snippetGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
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
