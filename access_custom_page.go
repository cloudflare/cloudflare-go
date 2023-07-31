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

func (api *API) AccessCustomPages(ctx context.Context, id string, pageOpts PaginationOptions) ([]AccessCustomPage, error) {
	return api.accessCustomPages(ctx, id, pageOpts, AccountRouteRoot)
}

func (api *API) accessCustomPages(ctx context.Context, id string, pageOpts PaginationOptions, routeRoot RouteRoot) ([]AccessCustomPage, error) {
	uri := buildURI(fmt.Sprintf("/%s/%s/access/apps/custom_pages", routeRoot, id), pageOpts)
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

func (api *API) AccessCustomPage(ctx context.Context, id string, customPageID string) (AccessCustomPage, error) {
	return api.accessCustomPage(ctx, id, customPageID, AccountRouteRoot)
}

func (api *API) accessCustomPage(ctx context.Context, id string, customPageID string, routeRoot RouteRoot) (AccessCustomPage, error) {
	uri := fmt.Sprintf("/%s/%s/access/apps/custom_pages/%s", routeRoot, id, customPageID)
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

func (api *API) CreateAccessCustomPage(ctx context.Context, id string, customPage AccessCustomPage) (AccessCustomPage, error) {
	return api.createCustomPage(ctx, id, customPage, AccountRouteRoot)
}

func (api *API) createCustomPage(ctx context.Context, id string, customPage AccessCustomPage, routeRoot RouteRoot) (AccessCustomPage, error) {
	uri := fmt.Sprintf("/%s/%s/access/apps/custom_pages", routeRoot, id)
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

func (api *API) DeleteAccessCustomPage(ctx context.Context, id string, customPageID string) error {
	return api.deleteCustomPage(ctx, id, customPageID, AccountRouteRoot)
}

func (api *API) deleteCustomPage(ctx context.Context, id string, customPageID string, routeRoot RouteRoot) error {
	uri := fmt.Sprintf("/%s/%s/access/apps/custom_pages/%s", routeRoot, id, customPageID)
	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}
	return nil
}

func (api *API) UpdateAccessCustomPage(ctx context.Context, id string, customPageID string, customPage AccessCustomPage) (AccessCustomPage, error) {
	return api.updateCustomPage(ctx, id, customPageID, customPage, AccountRouteRoot)
}

func (api *API) updateCustomPage(ctx context.Context, id string, customPageID string, customPage AccessCustomPage, routeRoot RouteRoot) (AccessCustomPage, error) {
	uri := fmt.Sprintf("/%s/%s/access/apps/custom_pages/%s", routeRoot, id, customPageID)
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
