package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// PageShieldScript represents a Page Shield script.
type PageShieldScript struct {
	AddedAt                 string   `json:"added_at"`
	DomainReportedMalicious bool     `json:"domain_reported_malicious"`
	FetchedAt               string   `json:"fetched_at"`
	FirstPageURL            string   `json:"first_page_url"`
	FirstSeenAt             string   `json:"first_seen_at"`
	Hash                    string   `json:"hash"`
	Host                    string   `json:"host"`
	ID                      string   `json:"id"`
	JsIntegrityScore        int      `json:"js_integrity_score"`
	LastSeenAt              string   `json:"last_seen_at"`
	PageURLs                []string `json:"page_urls"`
	URL                     string   `json:"url"`
	URLContainsCdnCgiPath   bool     `json:"url_contains_cdn_cgi_path"`
}

// ListPageShieldScriptsParams represents a PageShield Script request parameters
type ListPageShieldScriptsParams struct {
	Direction           string `json:"direction"`
	ExcludeCdnCgi       bool   `json:"exclude_cdn_cgi"`
	ExcludeDuplicates   bool   `json:"exclude_duplicates"`
	ExcludeUrls         string `json:"exclude_urls"`
	Export              string `json:"export"`
	Hosts               string `json:"hosts"`
	OrderBy             string `json:"order_by"`
	Page                string `json:"page"`
	PageURL             string `json:"page_url"`
	PerPage             int    `json:"per_page"`
	PrioritizeMalicious bool   `json:"prioritize_malicious"`
	Status              string `json:"status"`
	URLs                string `json:"urls"`
}

type PageShieldScriptsResponse struct {
	Results []PageShieldScript `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

type PageShieldScriptResponse struct {
	Result   PageShieldScript          `json:"result"`
	Versions []PageShieldScriptVersion `json:"versions"`
}

type PageShieldScriptVersion struct {
	FetchedAt        string `json:"fetched_at"`
	Hash             string `json:"hash"`
	JsIntegrityScore int    `json:"js_integrity_score"`
}

func (api *API) ListPageShieldScripts(ctx context.Context, rc *ResourceContainer, params ListPageShieldScriptsParams) ([]PageShieldScript, ResultInfo, error) {
	path := fmt.Sprintf("/zones/%s/page_shield/scripts", rc.Identifier)

	uri := buildURI(path, params)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, ResultInfo{}, err
	}

	var psResponse PageShieldScriptsResponse
	err = json.Unmarshal(res, &psResponse)
	if err != nil {
		return nil, ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return psResponse.Results, psResponse.ResultInfo, nil
}

func (api *API) GetPageShieldScript(ctx context.Context, rc *ResourceContainer, scriptID string) (*PageShieldScript, []PageShieldScriptVersion, error) {
	path := fmt.Sprintf("/zones/%s/page_shield/scripts/%s", rc.Identifier, scriptID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var psResponse PageShieldScriptResponse
	err = json.Unmarshal(res, &psResponse)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return &psResponse.Result, psResponse.Versions, nil
}
