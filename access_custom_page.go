package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type AccessCustomPageType string

const (
	Forbidden      AccessCustomPageType = "forbidden"
	IdentityDenied AccessCustomPageType = "identity_denied"
)

type AccessCustomPage struct {
	// The HTML content of the custom page.
	CustomHTML string               `json:"custom_html,omitempty"`
	Name       string               `json:"name,omitempty"`
	AppCount   int                  `json:"app_count,omitempty"`
	Type       AccessCustomPageType `json:"type,omitempty"`
	UID        string               `json:"uid,omitempty"`
}

type AccessCustomPageListResponse struct {
	Response
	Result     []AccessCustomPage `json:"result"`
	ResultInfo `json:"result_info"`
}

type AccessCustomPageResponse struct {
	Response
	Result AccessCustomPage `json:"result"`
}

func (api *API) AccessCustomPages(ctx context.Context, rc *ResourceContainer, pageOpts PaginationOptions) ([]AccessCustomPage, error) {
	uri := buildURI(fmt.Sprintf("/%s/%s/access/custom_pages", rc.Level, rc.Identifier), pageOpts)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []AccessCustomPage{}, err
	}

	var customPagesResponse AccessCustomPageListResponse
	err = json.Unmarshal(res, &customPagesResponse)
	if err != nil {
		return []AccessCustomPage{}, err
	}
	return customPagesResponse.Result, nil
}

func (api *API) AccessCustomPage(ctx context.Context, rc *ResourceContainer, customPageID string) (AccessCustomPage, error) {
	uri := fmt.Sprintf("/%s/%s/access/custom_pages/%s", rc.Level, rc.Identifier, customPageID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return AccessCustomPage{}, err
	}

	var customPageResponse AccessCustomPageResponse
	err = json.Unmarshal(res, &customPageResponse)
	if err != nil {
		return AccessCustomPage{}, err
	}
	return customPageResponse.Result, nil
}

func (api *API) CreateAccessCustomPage(ctx context.Context, rc *ResourceContainer, customPage AccessCustomPage) (AccessCustomPage, error) {
	uri := fmt.Sprintf("/%s/%s/access/custom_pages", rc.Level, rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, customPage)
	if err != nil {
		return AccessCustomPage{}, err
	}

	var customPageResponse AccessCustomPageResponse
	err = json.Unmarshal(res, &customPageResponse)
	if err != nil {
		return AccessCustomPage{}, err
	}
	return customPageResponse.Result, nil
}

func (api *API) DeleteAccessCustomPage(ctx context.Context, rc *ResourceContainer, customPageID string) error {
	uri := fmt.Sprintf("/%s/%s/access/custom_pages/%s", rc.Level, rc.Identifier, customPageID)
	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}
	return nil
}

func (api *API) UpdateAccessCustomPage(ctx context.Context, rc *ResourceContainer, customPageID string, customPage AccessCustomPage) (AccessCustomPage, error) {
	uri := fmt.Sprintf("/%s/%s/access/custom_pages/%s", rc.Level, rc.Identifier, customPageID)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, customPage)
	if err != nil {
		return AccessCustomPage{}, err
	}

	var customPageResponse AccessCustomPageResponse
	err = json.Unmarshal(res, &customPageResponse)
	if err != nil {
		return AccessCustomPage{}, err
	}
	return customPageResponse.Result, nil
}
