// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oauth_scopes

import (
	"context"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// OAuthScopeService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOAuthScopeService] method instead.
type OAuthScopeService struct {
	Options []option.RequestOption
}

// NewOAuthScopeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOAuthScopeService(opts ...option.RequestOption) (r *OAuthScopeService) {
	r = &OAuthScopeService{}
	r.Options = opts
	return
}

// List all available OAuth scopes. This endpoint requires authentication but has
// no authorization role requirements.
func (r *OAuthScopeService) List(ctx context.Context, opts ...option.RequestOption) (res *pagination.SinglePage[OAuthScopeListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "oauth/scopes"
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

// List all available OAuth scopes. This endpoint requires authentication but has
// no authorization role requirements.
func (r *OAuthScopeService) ListAutoPaging(ctx context.Context, opts ...option.RequestOption) *pagination.SinglePageAutoPager[OAuthScopeListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, opts...))
}

// An available OAuth scope that can be assigned to an OAuth client.
type OAuthScopeListResponse struct {
	// The scope label to use in the scopes array when creating or updating an OAuth
	// client.
	ID string `json:"id" api:"required"`
	// Human-readable name of the OAuth scope.
	Name string `json:"name" api:"required"`
	// Category for grouping scopes in the UI.
	Category string `json:"category"`
	// The underlying resource scopes (Bach scopes) that define which resources this
	// OAuth scope can act upon.
	Scopes []string                   `json:"scopes"`
	JSON   oauthScopeListResponseJSON `json:"-"`
}

// oauthScopeListResponseJSON contains the JSON metadata for the struct
// [OAuthScopeListResponse]
type oauthScopeListResponseJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Category    apijson.Field
	Scopes      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OAuthScopeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthScopeListResponseJSON) RawJSON() string {
	return r.raw
}
