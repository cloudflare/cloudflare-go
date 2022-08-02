package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type PhishingScanSubmission struct {
	ExcludedUrls     []URLSubmission `json:"excluded_urls"`
	SkippedUrls      []URLSubmission `json:"skipped_urls"`
	SubmittedDomains []interface{}   `json:"submitted_domains"`
	SubmittedUrls    []URLSubmission `json:"submitted_urls"`
}

type URLSubmission struct {
	URL   string `json:"url"`
	URLID int    `json:"url_id"`
}

// PhishingScanParameters represent parameters for a phishing scan request.
type PhishingScanSubmissionParameters struct {
	AccountID string `url:"-"`
	URL       string `url:"url,omitempty"`
	Skip      bool   `url:"skip,omitempty"`
}

// PhishingScanResponse represent an API response for a phishing scan.
type PhishingScanSubmissionResponse struct {
	Response
	Result []PhishingScanSubmission `json:"result,omitempty"`
}

// IntelligencePhishingScan scans a URL for suspected phishing
//
// API Reference: https://api.cloudflare.com/#phishing-url-scanner-scan-suspicious-url
func (api *API) IntelligencePhishingScanSubmission(ctx context.Context, params PhishingScanSubmissionParameters) (PhishingScanSubmissionResponse, error) {
	if params.AccountID == "" {
		return PhishingScanSubmissionResponse{}, ErrMissingAccountID
	}

	if params.URL == "" {
		return PhishingScanSubmissionResponse{}, ErrMissingURL
	}

	uri := buildURI(fmt.Sprintf("/accounts/%s/brand-protection/submit", params.AccountID), params)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return PhishingScanSubmissionResponse{}, err
	}

	var phishingScanResponse PhishingScanSubmissionResponse
	if err := json.Unmarshal(res, &phishingScanResponse); err != nil {
		return PhishingScanSubmissionResponse{}, err
	}
	return phishingScanResponse, nil
}

// PhishingScanResultsParameters represent parameters for a phishing scan request.
type PhishingScanResultsParameters struct {
	AccountID string `url:"-"`
	URL       string `url:"url,omitempty"`
}

// PhishingScanResultsResponse represent an API response for a phishing scan.
type PhishingScanResultsResponse struct {
	Response
	Result []PhishingScan `json:"result"`
}

type PhishingScan struct {
	Categorizations []interface{} `json:"categorizations"`
	ModelResults    []struct {
		ModelName  string  `json:"model_name"`
		ModelScore float64 `json:"model_score"`
	} `json:"model_results"`
	RuleMatches    []interface{} `json:"rule_matches"`
	ScanStatus     interface{}   `json:"scan_status"`
	ScreenshotPath string        `json:"screenshot_path"`
	Domain         string        `json:"domain"`
	URL            string        `json:"url"`
}

// IntelligencePhishingScanResults scans a URL for suspected phishing
//
// API Reference: https://api.cloudflare.com/#phishing-url-scanner-scan-suspicious-url
func (api *API) IntelligencePhishingScanResults(ctx context.Context, params PhishingScanResultsParameters) (PhishingScanResultsResponse, error) {
	if params.AccountID == "" {
		return PhishingScanResultsResponse{}, ErrMissingAccountID
	}

	if params.URL == "" {
		return PhishingScanResultsResponse{}, ErrMissingURL
	}

	uri := buildURI(fmt.Sprintf("/accounts/%s/brand-protection/url-info", params.AccountID), params)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return PhishingScanResultsResponse{}, err
	}

	var phishingScanResponse PhishingScanResultsResponse
	if err := json.Unmarshal(res, &phishingScanResponse); err != nil {
		return PhishingScanResultsResponse{}, err
	}
	return phishingScanResponse, nil
}
