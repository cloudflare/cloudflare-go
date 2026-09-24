// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_rooms

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// SettingService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingService] method instead.
type SettingService struct {
	Options []option.RequestOption
}

// NewSettingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSettingService(opts ...option.RequestOption) (r *SettingService) {
	r = &SettingService{}
	r.Options = opts
	return
}

// Replace zone-level Waiting Room settings.
func (r *SettingService) Update(ctx context.Context, params SettingUpdateParams, opts ...option.RequestOption) (res *SettingUpdateResponse, err error) {
	var env SettingUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/waiting_rooms/settings", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Patch zone-level Waiting Room settings.
func (r *SettingService) Edit(ctx context.Context, params SettingEditParams, opts ...option.RequestOption) (res *SettingEditResponse, err error) {
	var env SettingEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/waiting_rooms/settings", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Get zone-level Waiting Room settings.
func (r *SettingService) Get(ctx context.Context, query SettingGetParams, opts ...option.RequestOption) (res *SettingGetResponse, err error) {
	var env SettingGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/waiting_rooms/settings", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type SettingUpdateResponse struct {
	// Whether to allow verified search engine crawlers to bypass all waiting rooms on
	// this zone. Verified search engine crawlers will not be tracked or counted by the
	// waiting room system, and will not appear in waiting room analytics.
	SearchEngineCrawlerBypass bool                      `json:"search_engine_crawler_bypass" api:"required"`
	JSON                      settingUpdateResponseJSON `json:"-"`
}

// settingUpdateResponseJSON contains the JSON metadata for the struct
// [SettingUpdateResponse]
type settingUpdateResponseJSON struct {
	SearchEngineCrawlerBypass apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *SettingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type SettingEditResponse struct {
	// Whether to allow verified search engine crawlers to bypass all waiting rooms on
	// this zone. Verified search engine crawlers will not be tracked or counted by the
	// waiting room system, and will not appear in waiting room analytics.
	SearchEngineCrawlerBypass bool                    `json:"search_engine_crawler_bypass" api:"required"`
	JSON                      settingEditResponseJSON `json:"-"`
}

// settingEditResponseJSON contains the JSON metadata for the struct
// [SettingEditResponse]
type settingEditResponseJSON struct {
	SearchEngineCrawlerBypass apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *SettingEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseJSON) RawJSON() string {
	return r.raw
}

type SettingGetResponse struct {
	// Whether to allow verified search engine crawlers to bypass all waiting rooms on
	// this zone. Verified search engine crawlers will not be tracked or counted by the
	// waiting room system, and will not appear in waiting room analytics.
	SearchEngineCrawlerBypass bool                   `json:"search_engine_crawler_bypass" api:"required"`
	JSON                      settingGetResponseJSON `json:"-"`
}

// settingGetResponseJSON contains the JSON metadata for the struct
// [SettingGetResponse]
type settingGetResponseJSON struct {
	SearchEngineCrawlerBypass apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *SettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseJSON) RawJSON() string {
	return r.raw
}

type SettingUpdateParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// Whether to allow verified search engine crawlers to bypass all waiting rooms on
	// this zone. Verified search engine crawlers will not be tracked or counted by the
	// waiting room system, and will not appear in waiting room analytics.
	SearchEngineCrawlerBypass param.Field[bool] `json:"search_engine_crawler_bypass"`
}

func (r SettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingUpdateResponseEnvelope struct {
	Errors   []SettingUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   SettingUpdateResponse                   `json:"result" api:"required,nullable"`
	Success  bool                                    `json:"success" api:"required"`
	JSON     settingUpdateResponseEnvelopeJSON       `json:"-"`
}

// settingUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingUpdateResponseEnvelope]
type settingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingUpdateResponseEnvelopeErrors struct {
	Code             int64                                     `json:"code" api:"required"`
	Message          string                                    `json:"message" api:"required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           SettingUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingUpdateResponseEnvelopeErrors]
type settingUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    settingUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [SettingUpdateResponseEnvelopeErrorsSource]
type settingUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingUpdateResponseEnvelopeMessages struct {
	Code             int64                                       `json:"code" api:"required"`
	Message          string                                      `json:"message" api:"required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           SettingUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingUpdateResponseEnvelopeMessages]
type settingUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    settingUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [SettingUpdateResponseEnvelopeMessagesSource]
type settingUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type SettingEditParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// Whether to allow verified search engine crawlers to bypass all waiting rooms on
	// this zone. Verified search engine crawlers will not be tracked or counted by the
	// waiting room system, and will not appear in waiting room analytics.
	SearchEngineCrawlerBypass param.Field[bool] `json:"search_engine_crawler_bypass"`
}

func (r SettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingEditResponseEnvelope struct {
	Errors   []SettingEditResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingEditResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   SettingEditResponse                   `json:"result" api:"required,nullable"`
	Success  bool                                  `json:"success" api:"required"`
	JSON     settingEditResponseEnvelopeJSON       `json:"-"`
}

// settingEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingEditResponseEnvelope]
type settingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingEditResponseEnvelopeErrors struct {
	Code             int64                                   `json:"code" api:"required"`
	Message          string                                  `json:"message" api:"required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           SettingEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingEditResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [SettingEditResponseEnvelopeErrors]
type settingEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingEditResponseEnvelopeErrorsSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    settingEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [SettingEditResponseEnvelopeErrorsSource]
type settingEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingEditResponseEnvelopeMessages struct {
	Code             int64                                     `json:"code" api:"required"`
	Message          string                                    `json:"message" api:"required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           SettingEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingEditResponseEnvelopeMessages]
type settingEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingEditResponseEnvelopeMessagesSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    settingEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingEditResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [SettingEditResponseEnvelopeMessagesSource]
type settingEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type SettingGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type SettingGetResponseEnvelope struct {
	Errors   []SettingGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingGetResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   SettingGetResponse                   `json:"result" api:"required,nullable"`
	Success  bool                                 `json:"success" api:"required"`
	JSON     settingGetResponseEnvelopeJSON       `json:"-"`
}

// settingGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingGetResponseEnvelope]
type settingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingGetResponseEnvelopeErrors struct {
	Code             int64                                  `json:"code" api:"required"`
	Message          string                                 `json:"message" api:"required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           SettingGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [SettingGetResponseEnvelopeErrors]
type settingGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingGetResponseEnvelopeErrorsSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    settingGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [SettingGetResponseEnvelopeErrorsSource]
type settingGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingGetResponseEnvelopeMessages struct {
	Code             int64                                    `json:"code" api:"required"`
	Message          string                                   `json:"message" api:"required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           SettingGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [SettingGetResponseEnvelopeMessages]
type settingGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingGetResponseEnvelopeMessagesSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    settingGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [SettingGetResponseEnvelopeMessagesSource]
type settingGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}
