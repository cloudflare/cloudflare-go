package cloudflare

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var ErrMissingScriptID = errors.New("required script id missing")

type PageShieldScript struct {
	ScriptID                string                     `json:"script_id,omitempty"`
	ScriptURL               string                     `json:"script_url,omitempty"`
	AddedAt                 *time.Time                 `json:"added_at,omitempty"`
	FirstSeenAt             *time.Time                 `json:"first_seen_at,omitempty"`
	LastSeenAt              *time.Time                 `json:"last_seen_at,omitempty"`
	Host                    string                     `json:"host,omitempty"`
	DomainReportedMalicious bool                       `json:"domain_reported_malicious,omitempty"`
	SeenOnFirst             string                     `json:"seen_on_first,omitempty"`
	SeenOn                  []string                   `json:"seen_on,omitempty"`
	AppearsInCDNCGIPath     bool                       `json:"appears_in_cdn_cgi_path,omitempty"`
	Hash                    string                     `json:"hash,omitempty"`
	JsIntegrityScore        int                        `json:"js_integrity_score,omitempty"`
	FetchedAt               *time.Time                 `json:"fetched_at,omitempty"`
	Versions                []PageShieldScriptVersions `json:"versions,omitempty"`
}

type PageShieldScriptVersions struct {
	Hash             string     `json:"hash,omitempty"`
	JsIntegrityScore int        `json:"js_integrity_score,omitempty"`
	FetchedAt        *time.Time `json:"fetched_at,omitempty"`
}

type ListPageShieldParameters struct {
	ScriptPage          string `url:"script_page,omitempty"`
	OrderBy             string `url:"order_by,omitempty"`
	Hosts               string `url:"hosts,omitempty"`
	ExcludeCDNCGI       bool   `url:"exclude_cdn_cgi,omitempty"`
	Status              string `url:"status,omitempty"`
	ExcludeScriptURLS   string `url:"exclude_script_urls,omitempty"`
	PrioritizeMalicious bool   `url:"prioritize_malicious,omitempty"`
	ScriptURLs          string `url:"script_urls,omitempty"`
	Direction           string `url:"direction,omitempty"`
	ResultInfo
}

type ListPageShieldResponse struct {
	Result     []PageShieldScript `json:"result"`
	ResultInfo `json:"result_info,omitempty"`
	Response
}

// GetPageShieldStatus Fetches the Page Shield status.
//
// API Reference: https://api.cloudflare.com/#page-shield-get-page-shield-status
func (api *API) GetPageShieldStatus(ctx context.Context, rc *ResourceContainer) error {
	if rc.Identifier == "" {
		return ErrMissingZoneID
	}
	uri := fmt.Sprintf("/zones/%s/page_shield", rc.Identifier)

	_, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	return err
}

// UpdatePageShieldStatus Fetches the Page Shield status.
//
// API Reference: https://api.cloudflare.com/#page-shield-update-page-shield-status
func (api *API) UpdatePageShieldStatus(ctx context.Context, rc *ResourceContainer) error {
	if rc.Identifier == "" {
		return ErrMissingZoneID
	}
	uri := fmt.Sprintf("/zones/%s/page_shield", rc.Identifier)

	_, err := api.makeRequestContext(ctx, http.MethodPut, uri, nil)
	return err
}

// ListPageShieldScripts Fetches the Page Shield status.
//
// API Reference: https://api.cloudflare.com/#page-shield-list-page-shield-scripts
func (api *API) ListPageShieldScripts(ctx context.Context, rc *ResourceContainer, params ListPageShieldParameters) ([]PageShieldScript, *ResultInfo, error) {
	if rc.Identifier == "" {
		return []PageShieldScript{}, &ResultInfo{}, ErrMissingZoneID
	}

	autoPaginate := true
	if params.PerPage >= 1 || params.Page >= 1 {
		autoPaginate = false
	}
	if params.PerPage < 1 {
		params.PerPage = 50
	}
	if params.Page < 1 {
		params.Page = 1
	}
	var scripts []PageShieldScript
	var rResponse ListPageShieldResponse
	for {
		uri := buildURI(fmt.Sprintf("/zones/%s/page_shield/scripts", rc.Identifier), params)
		res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
		if err != nil {
			return []PageShieldScript{}, &ResultInfo{}, err
		}
		if err := json.Unmarshal(res, &rResponse); err != nil {
			return []PageShieldScript{}, &ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
		}
		scripts = append(scripts, rResponse.Result...)
		params.ResultInfo = rResponse.ResultInfo.Next()
		if params.ResultInfo.Done() || !autoPaginate {
			break
		}
	}
	return scripts, &rResponse.ResultInfo, nil
}

// GetPageShieldScript Fetches the Page Shield status.
//
// API Reference: https://api.cloudflare.com/#page-shield-get-a-page-shield-script
func (api *API) GetPageShieldScript(ctx context.Context, rc *ResourceContainer, ScriptID string) (PageShieldScript, error) {
	if rc.Identifier == "" {
		return PageShieldScript{}, ErrMissingZoneID
	}
	if ScriptID == "" {
		return PageShieldScript{}, ErrMissingScriptID
	}
	uri := fmt.Sprintf("/zones/%s/page_shield/scripts/%s", rc.Identifier, ScriptID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	var r PageShieldScript
	if err := json.Unmarshal(res, &r); err != nil {
		return PageShieldScript{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r, err
}
