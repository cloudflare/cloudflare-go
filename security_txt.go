package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"net/http"
	"time"
)

type SecurityTXT struct {
	Acknowledgements   []string   `json:"acknowledgments"`
	Canonical          []string   `json:"canonical"`
	Contact            []string   `json:"contact"`
	Enabled            bool       `json:"enabled"`
	Encryption         []string   `json:"encryption"`
	Expires            *time.Time `json:"expires"`
	Hiring             []string   `json:"hiring"`
	Policy             []string   `json:"policy"`
	PreferredLanguages string     `json:"preferredLanguages"`
}

type SecurityTXTResponse struct {
	Response
	Result SecurityTXT `json:"result"`
}

func (api *API) GetSecurityTXT(ctx context.Context, rc *ResourceContainer) (SecurityTXT, error) {
	if rc.Identifier == "" {
		return SecurityTXT{}, ErrMissingZoneID
	}

	uri := buildURI(fmt.Sprintf("/zones/%s/security-center/securitytxt", rc.Identifier), nil)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return SecurityTXT{}, err
	}

	result := SecurityTXTResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return SecurityTXT{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	if !result.Success {
		return SecurityTXT{}, errors.New(result.Messages[0].Message)
	}

	return result.Result, nil
}

func (api *API) UpdateSecurityTXT(ctx context.Context, rc *ResourceContainer, securityTXT SecurityTXT) (SecurityTXT, error) {
	if rc.Identifier == "" {
		return SecurityTXT{}, ErrMissingZoneID
	}

	uri := buildURI(fmt.Sprintf("/zones/%s/security-center/securitytxt", rc.Identifier), nil)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, securityTXT)
	if err != nil {
		return SecurityTXT{}, err
	}

	result := SecurityTXTResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return SecurityTXT{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	if !result.Success {
		return SecurityTXT{}, errors.New(result.Messages[0].Message)
	}

	return result.Result, nil
}

func (api *API) DeleteSecurityTXT(ctx context.Context, rc *ResourceContainer) error {
	if rc.Identifier == "" {
		return ErrMissingZoneID
	}

	uri := buildURI(fmt.Sprintf("/zones/%s/security-center/securitytxt", rc.Identifier), nil)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	result := SecurityTXTResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	if !result.Success {
		return errors.New(result.Messages[0].Message)
	}

	return nil
}
