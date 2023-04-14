package cloudflare

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var ErrMissingSiteKey = errors.New("required site key missing")

type ChallengeWidget struct {
	SiteKey      string     `json:"sitekey,omitempty"`
	Secret       string     `json:"secret,omitempty"`
	CreatedOn    *time.Time `json:"created_on,omitempty"`
	ModifiedOn   *time.Time `json:"modified_on,omitempty"`
	Name         string     `json:"name,omitempty"`
	Domains      []string   `json:"domains,omitempty"`
	Type         string     `json:"type,omitempty"`
	BotFightMode bool       `json:"bot_fight_mode,omitempty"`
}

type ChallengeWidgetResponse struct {
	Response
	Result ChallengeWidget `json:"result"`
}

type ListChallengeWidgetRequest struct {
	ResultInfo
	Direction string `url:"direction,omitempty"`
	Order     string `url:"order,omitempty"`
}

type ListChallengeWidgetResponse struct {
	Response
	ResultInfo `json:"result_info"`
	Result     []ChallengeWidget `json:"result"`
}

type RotateChallengeWidgetRequest struct {
	SiteKey               string
	InvalidateImmediately bool `json:"invalidate_immediately,omitempty"`
}

// CreateChallengeWidget creates a new challenge widgets.
//
// API reference: https://api.cloudflare.com/#challenge-widgets-properties
func (api *API) CreateChallengeWidget(ctx context.Context, rc *ResourceContainer, params ChallengeWidget) (ChallengeWidget, error) {
	if rc.Identifier == "" {
		return ChallengeWidget{}, ErrMissingAccountID
	}
	uri := fmt.Sprintf("/accounts/%s/challenges/widgets", rc.Identifier)
	res, err := api.makeRequestContext(ctx, "POST", uri, params)
	if err != nil {
		return ChallengeWidget{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}
	var r ChallengeWidgetResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return ChallengeWidget{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// ListChallengeWidget lists challenge widgets.
//
// API reference: https://api.cloudflare.com/#challenge-widgets-list-challenge-widgets
func (api *API) ListChallengeWidget(ctx context.Context, rc *ResourceContainer, params ListChallengeWidgetRequest) ([]ChallengeWidget, *ResultInfo, error) {
	if rc.Identifier == "" {
		return []ChallengeWidget{}, &ResultInfo{}, ErrMissingAccountID
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

	var widgets []ChallengeWidget
	var r ListChallengeWidgetResponse
	for {
		uri := buildURI(fmt.Sprintf("/accounts/%s/challenges/widgets", rc.Identifier), params)

		res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
		if err != nil {
			return []ChallengeWidget{}, &ResultInfo{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
		}
		err = json.Unmarshal(res, &r)
		if err != nil {
			return []ChallengeWidget{}, &ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
		}
		widgets = append(widgets, r.Result...)
		params.ResultInfo = r.ResultInfo.Next()
		if params.ResultInfo.Done() || !autoPaginate {
			break
		}
	}

	return widgets, &r.ResultInfo, nil
}

// GetChallengeWidget shows a single challenge widget configuration..
//
// API reference: https://api.cloudflare.com/#challenge-widgets-challenge-widget-details
func (api *API) GetChallengeWidget(ctx context.Context, rc *ResourceContainer, SiteKey string) (ChallengeWidget, error) {
	if rc.Identifier == "" {
		return ChallengeWidget{}, ErrMissingAccountID
	}

	if SiteKey == "" {
		return ChallengeWidget{}, ErrMissingSiteKey
	}

	uri := fmt.Sprintf("/accounts/%s/challenges/widgets/%s", rc.Identifier, SiteKey)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)

	if err != nil {
		return ChallengeWidget{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}
	var r ChallengeWidgetResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return ChallengeWidget{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r.Result, nil
}

// UpdateChallengeWidget update the configuration of a widget.
//
// API reference: https://api.cloudflare.com/#challenge-widgets-update-a-challenge-widget
func (api *API) UpdateChallengeWidget(ctx context.Context, rc *ResourceContainer, widget ChallengeWidget) (ChallengeWidget, error) {
	if rc.Identifier == "" {
		return ChallengeWidget{}, ErrMissingAccountID
	}
	if widget.SiteKey == "" {
		return ChallengeWidget{}, ErrMissingSiteKey
	}
	uri := fmt.Sprintf("/accounts/%s/challenges/widgets/%s", rc.Identifier, widget.SiteKey)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, widget)
	if err != nil {
		return ChallengeWidget{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}
	var r ChallengeWidgetResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return ChallengeWidget{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// RotateChallengeWidget generates a new secret key for this widget. If invalidate_immediately is set to false, the previous secret remains valid for 2 hours.
//
// Note that secrets cannot be rotated again during the grace period.
//
// API reference: https://api.cloudflare.com/#challenge-widgets-rotate-secret-for-a-challenge-widget
func (api *API) RotateChallengeWidget(ctx context.Context, rc *ResourceContainer, param RotateChallengeWidgetRequest) (ChallengeWidget, error) {
	if rc.Identifier == "" {
		return ChallengeWidget{}, ErrMissingAccountID
	}
	if param.SiteKey == "" {
		return ChallengeWidget{}, ErrMissingSiteKey
	}

	uri := fmt.Sprintf("/accounts/%s/challenges/widgets/%s/rotate_secret", rc.Identifier, param.SiteKey)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, param)

	if err != nil {
		return ChallengeWidget{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}
	var r ChallengeWidgetResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return ChallengeWidget{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// DeleteChallengeWidget delete a challenge widget.
//
// API reference: https://api.cloudflare.com/#challenge-widgets-delete-a-challenge-widget
func (api *API) DeleteChallengeWidget(ctx context.Context, rc *ResourceContainer, SiteKey string) error {
	if rc.Identifier == "" {
		return ErrMissingAccountID
	}

	if SiteKey == "" {
		return ErrMissingSiteKey
	}
	uri := fmt.Sprintf("/accounts/%s/challenges/widgets/%s", rc.Identifier, SiteKey)

	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var r ChallengeWidgetResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return nil
}
