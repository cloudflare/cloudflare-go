package cloudflare

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var testEmailRoutingRule = EmailRoutingRule{
	Tag:      "a7e6fb77503c41d8a7f3113c6918f10c",
	Name:     "Rule send to user@example.net",
	Priority: 0,
	Enabled:  true,
	Matchers: []EmailRoutingRuleMatcher{
		{
			Type:  "literal",
			Field: "to",
			Value: "test@example.com",
		},
	},
	Actions: []EmailRoutingRuleAction{
		{
			Type:  "forward",
			Value: []string{"destinationaddress@example.net"},
		},
	},
}

func TestEmailRouting_ListRoutingRules(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/"+testZoneID+"/email/routing/rules", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
  "success": true,
  "errors": [],
  "messages": [],
  "result": [
    {
      "tag": "a7e6fb77503c41d8a7f3113c6918f10c",
      "name": "Rule send to user@example.net",
      "priority": 0,
      "enabled": true,
      "matchers": [
        {
          "type": "literal",
          "field": "to",
          "value": "test@example.com"
        }
      ],
      "actions": [
        {
          "type": "forward",
          "value": [
            "destinationaddress@example.net"
          ]
        }
      ]
    }
  ],
  "result_info": {
    "page": 1,
    "per_page": 20,
    "count": 1,
    "total_count": 1
  }
}`)
	})

	_, _, err := client.ListEmailRoutingRules(context.Background(), AccountIdentifier(""), EmailRoutingListRulesParameters{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingZoneID, err)
	}

	res, resInfo, err := client.ListEmailRoutingRules(context.Background(), AccountIdentifier(testZoneID), EmailRoutingListRulesParameters{Enabled: true})
	if assert.NoError(t, err) {
		assert.Equal(t, resInfo.Page, 1)
		assert.Equal(t, testEmailRoutingRule, res[0])
	}
}

func TestEmailRouting_CreateRoutingRule(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/"+testZoneID+"/email/routing/rules", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
  "success": true,
  "errors": [],
  "messages": [],
  "result": {
    "tag": "a7e6fb77503c41d8a7f3113c6918f10c",
    "name": "Rule send to user@example.net",
    "priority": 0,
    "enabled": true,
    "matchers": [
      {
        "type": "literal",
        "field": "to",
        "value": "test@example.com"
      }
    ],
    "actions": [
      {
        "type": "forward",
        "value": [
          "destinationaddress@example.net"
        ]
      }
    ]
  }
}`)
	})

	_, err := client.CreateEmailRoutingRule(context.Background(), AccountIdentifier(""), EmailRoutingCreateRuleParameters{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingZoneID, err)
	}

	res, err := client.CreateEmailRoutingRule(context.Background(), AccountIdentifier(testZoneID), EmailRoutingCreateRuleParameters{Enabled: true})
	if assert.NoError(t, err) {
		assert.Equal(t, testEmailRoutingRule, res)
	}
}

func TestEmailRouting_GetRoutingRule(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/"+testZoneID+"/email/routing/rules/a7e6fb77503c41d8a7f3113c6918f10c", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
  "success": true,
  "errors": [],
  "messages": [],
  "result": {
    "tag": "a7e6fb77503c41d8a7f3113c6918f10c",
    "name": "Rule send to user@example.net",
    "priority": 0,
    "enabled": true,
    "matchers": [
      {
        "type": "literal",
        "field": "to",
        "value": "test@example.com"
      }
    ],
    "actions": [
      {
        "type": "forward",
        "value": [
          "destinationaddress@example.net"
        ]
      }
    ]
  }
}`)
	})

	_, err := client.GetEmailRoutingRule(context.Background(), ZoneIdentifier(""), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingZoneID, err)
	}

	res, err := client.GetEmailRoutingRule(context.Background(), AccountIdentifier(testZoneID), "a7e6fb77503c41d8a7f3113c6918f10c")
	if assert.NoError(t, err) {
		assert.Equal(t, testEmailRoutingRule, res)
	}
}

func TestEmailRouting_UpdateRoutingRule(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/"+testZoneID+"/email/routing/rules/a7e6fb77503c41d8a7f3113c6918f10c", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
  "success": true,
  "errors": [],
  "messages": [],
  "result": {
    "tag": "a7e6fb77503c41d8a7f3113c6918f10c",
    "name": "Rule send to user@example.net",
    "priority": 0,
    "enabled": true,
    "matchers": [
      {
        "type": "literal",
        "field": "to",
        "value": "test@example.com"
      }
    ],
    "actions": [
      {
        "type": "forward",
        "value": [
          "destinationaddress@example.net"
        ]
      }
    ]
  }
}`)
	})

	_, err := client.UpdateEmailRoutingRule(context.Background(), ZoneIdentifier(""), EmailRoutingUpdateRuleParameters{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingZoneID, err)
	}
	_, err = client.UpdateEmailRoutingRule(context.Background(), ZoneIdentifier(testZoneID), EmailRoutingUpdateRuleParameters{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingRuleID, err)
	}

	res, err := client.UpdateEmailRoutingRule(context.Background(), AccountIdentifier(testZoneID), EmailRoutingUpdateRuleParameters{RuleID: "a7e6fb77503c41d8a7f3113c6918f10c"})
	if assert.NoError(t, err) {
		assert.Equal(t, testEmailRoutingRule, res)
	}
}

func TestEmailRouting_DeleteRoutingRule(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/"+testZoneID+"/email/routing/rules/a7e6fb77503c41d8a7f3113c6918f10c", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
  "success": true,
  "errors": [],
  "messages": [],
  "result": {
    "tag": "a7e6fb77503c41d8a7f3113c6918f10c",
    "name": "Rule send to user@example.net",
    "priority": 0,
    "enabled": true,
    "matchers": [
      {
        "type": "literal",
        "field": "to",
        "value": "test@example.com"
      }
    ],
    "actions": [
      {
        "type": "forward",
        "value": [
          "destinationaddress@example.net"
        ]
      }
    ]
  }
}`)
	})

	_, err := client.DeleteEmailRoutingRule(context.Background(), ZoneIdentifier(""), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingZoneID, err)
	}

	res, err := client.DeleteEmailRoutingRule(context.Background(), AccountIdentifier(testZoneID), "a7e6fb77503c41d8a7f3113c6918f10c")
	if assert.NoError(t, err) {
		assert.Equal(t, testEmailRoutingRule, res)
	}
}

func TestEmailRouting_GetAllRule(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/"+testZoneID+"/email/routing/rules/catch_all", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
  "tag": "a7e6fb77503c41d8a7f3113c6918f10c",
  "name": "Rule send to user@example.net",
  "enabled": true,
  "matchers": [
    {
      "type": "all"
    }
  ],
  "actions": [
    {
      "type": "forward",
      "value": [
        "destinationaddress@example.net"
      ]
    }
  ]
}`)
	})

	_, err := client.GetEmailRoutingCatchAllRule(context.Background(), ZoneIdentifier(""))
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingZoneID, err)
	}

	want := EmailRoutingCatchAllRule{
		Tag:     "a7e6fb77503c41d8a7f3113c6918f10c",
		Name:    "Rule send to user@example.net",
		Enabled: true,
		Matchers: []EmailRoutingRuleMatcher{
			{
				Type: "all",
			},
		},
		Actions: []EmailRoutingRuleAction{
			{
				Type:  "forward",
				Value: []string{"destinationaddress@example.net"},
			},
		},
	}

	res, err := client.GetEmailRoutingCatchAllRule(context.Background(), AccountIdentifier(testZoneID))
	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}

func TestEmailRouting_UpdateAllRule(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/"+testZoneID+"/email/routing/rules/catch_all", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
  "tag": "a7e6fb77503c41d8a7f3113c6918f10c",
  "name": "Rule send to user@example.net",
  "enabled": true,
  "matchers": [
    {
      "type": "all"
    }
  ],
  "actions": [
    {
      "type": "forward",
      "value": [
        "destinationaddress@example.net"
      ]
    }
  ]
}`)
	})

	_, err := client.UpdateEmailRoutingCatchAllRule(context.Background(), ZoneIdentifier(""), EmailRoutingCatchAllRule{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingZoneID, err)
	}

	want := EmailRoutingCatchAllRule{
		Tag:     "a7e6fb77503c41d8a7f3113c6918f10c",
		Name:    "Rule send to user@example.net",
		Enabled: true,
		Matchers: []EmailRoutingRuleMatcher{
			{
				Type: "all",
			},
		},
		Actions: []EmailRoutingRuleAction{
			{
				Type:  "forward",
				Value: []string{"destinationaddress@example.net"},
			},
		},
	}

	res, err := client.UpdateEmailRoutingCatchAllRule(context.Background(), AccountIdentifier(testZoneID), EmailRoutingCatchAllRule{})
	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}
