package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// RateLimit is a policy than can be applied to limit traffic within a customer domain.
type RateLimit struct {
	ID          string                  `json:"id,omitempty"`
	Disabled    bool                    `json:"disabled,omitempty"`
	Description string                  `json:"description,omitempty"`
	Match       RateLimitTrafficMatcher `json:"match"`
	Bypass      []RateLimitKeyValue     `json:"bypass,omitempty"`
	Threshold   int                     `json:"threshold"`
	Period      int                     `json:"period"`
	Action      RateLimitAction         `json:"action"`
	Correlate   *RateLimitCorrelate     `json:"correlate,omitempty"`
}

// RateLimitTrafficMatcher contains the rules that will be used to apply a rate limit to traffic.
type RateLimitTrafficMatcher struct {
	Request  RateLimitRequestMatcher  `json:"request"`
	Response RateLimitResponseMatcher `json:"response"`
}

// RateLimitRequestMatcher contains the matching rules pertaining to requests.
type RateLimitRequestMatcher struct {
	Methods    []string `json:"methods,omitempty"`
	Schemes    []string `json:"schemes,omitempty"`
	URLPattern string   `json:"url,omitempty"`
}

// RateLimitResponseMatcher contains the matching rules pertaining to responses.
type RateLimitResponseMatcher struct {
	Statuses      []int                            `json:"status,omitempty"`
	OriginTraffic *bool                            `json:"origin_traffic,omitempty"` // api defaults to true so we need an explicit empty value
	Headers       []RateLimitResponseMatcherHeader `json:"headers,omitempty"`
}

// RateLimitResponseMatcherHeader contains the structure of the origin
// HTTP headers used in request matcher checks.
type RateLimitResponseMatcherHeader struct {
	Name  string `json:"name"`
	Op    string `json:"op"`
	Value string `json:"value"`
}

// RateLimitKeyValue is k-v formatted as expected in the rate limit description.
type RateLimitKeyValue struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// RateLimitAction is the action that will be taken when the rate limit threshold is reached.
type RateLimitAction struct {
	Mode     string                   `json:"mode"`
	Timeout  int                      `json:"timeout"`
	Response *RateLimitActionResponse `json:"response"`
}

// RateLimitActionResponse is the response that will be returned when rate limit action is triggered.
type RateLimitActionResponse struct {
	ContentType string `json:"content_type"`
	Body        string `json:"body"`
}

// RateLimitCorrelate pertainings to NAT support.
type RateLimitCorrelate struct {
	By string `json:"by"`
}

type rateLimitResponse struct {
	Response
	Result RateLimit `json:"result"`
}

type rateLimitListResponse struct {
	Response
	Result     []RateLimit `json:"result"`
	ResultInfo ResultInfo  `json:"result_info"`
}

type RateLimitListParams struct {
	ResultInfo
}

// CreateRateLimit creates a new rate limit for a zone.
//
// API reference: https://api.cloudflare.com/#rate-limits-for-a-zone-create-a-ratelimit
func (api *API) CreateRateLimit(ctx context.Context, zoneID string, limit RateLimit) (RateLimit, error) {
	uri := fmt.Sprintf("/zones/%s/rate_limits", zoneID)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, limit)
	if err != nil {
		return RateLimit{}, err
	}
	var r rateLimitResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return RateLimit{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// ListRateLimits returns Rate Limits for a zone, paginated according to the provided params
//
// API reference: https://api.cloudflare.com/#rate-limits-for-a-zone-list-rate-limits
func (api *API) ListRateLimits(ctx context.Context, rc *ResourceContainer, params RateLimitListParams) ([]RateLimit, *ResultInfo, error) {
	uri := buildURI(fmt.Sprintf("/zones/%s/rate_limits", rc.Identifier), params)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []RateLimit{}, &ResultInfo{}, err
	}

	var r rateLimitListResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return []RateLimit{}, &ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	if params.PerPage < 1 && params.Page < 1 {
		var rateLimits []RateLimit
		//params.PerPage = 50
		//params.Page = 1

		for !params.ResultInfo.Done() {
			uri := buildURI(fmt.Sprintf("/zones/%s/rate_limits", rc.Identifier), params)

			res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
			if err != nil {
				return []RateLimit{}, &ResultInfo{}, err
			}

			var rResponse rateLimitListResponse
			err = json.Unmarshal(res, &rResponse)
			if err != nil {
				return []RateLimit{}, &ResultInfo{}, fmt.Errorf("failed to unmarshal filters JSON data: %w", err)
			}

			rateLimits = append(rateLimits, rResponse.Result...)
			params.ResultInfo = rResponse.ResultInfo.Next()
		}
		r.Result = rateLimits
	}

	return r.Result, &r.ResultInfo, nil
}

// RateLimit fetches detail about one Rate Limit for a zone.
//
// API reference: https://api.cloudflare.com/#rate-limits-for-a-zone-rate-limit-details
func (api *API) RateLimit(ctx context.Context, rc *ResourceContainer, limitID string) (RateLimit, error) {
	uri := fmt.Sprintf("/zones/%s/rate_limits/%s", rc.Identifier, limitID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return RateLimit{}, err
	}
	var r rateLimitResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return RateLimit{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// UpdateRateLimit lets you replace a Rate Limit for a zone.
//
// API reference: https://api.cloudflare.com/#rate-limits-for-a-zone-update-rate-limit
func (api *API) UpdateRateLimit(ctx context.Context, zoneID, limitID string, limit RateLimit) (RateLimit, error) {
	uri := fmt.Sprintf("/zones/%s/rate_limits/%s", zoneID, limitID)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, limit)
	if err != nil {
		return RateLimit{}, err
	}
	var r rateLimitResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return RateLimit{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// DeleteRateLimit deletes a Rate Limit for a zone.
//
// API reference: https://api.cloudflare.com/#rate-limits-for-a-zone-delete-rate-limit
func (api *API) DeleteRateLimit(ctx context.Context, zoneID, limitID string) error {
	uri := fmt.Sprintf("/zones/%s/rate_limits/%s", zoneID, limitID)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}
	var r rateLimitResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return nil
}
