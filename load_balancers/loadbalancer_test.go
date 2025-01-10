// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancers_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/load_balancers"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

func TestLoadBalancerNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.LoadBalancers.New(context.TODO(), load_balancers.LoadBalancerNewParams{
		ZoneID:       cloudflare.F("699d98642c564d2e855e9661899b7252"),
		DefaultPools: cloudflare.F([]load_balancers.DefaultPoolsParam{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
		FallbackPool: cloudflare.F("fallback_pool"),
		Name:         cloudflare.F("www.example.com"),
		AdaptiveRouting: cloudflare.F(load_balancers.AdaptiveRoutingParam{
			FailoverAcrossPools: cloudflare.F(true),
		}),
		CountryPools: cloudflare.F(map[string][]string{
			"GB": {"abd90f38ced07c2e2f4df50b1f61d4194"},
			"US": {"de90f38ced07c2e2f4df50b1f61d4194", "00920f38ce07c2e2f4df50b1f61d4194"},
		}),
		Description: cloudflare.F("Load Balancer for www.example.com"),
		LocationStrategy: cloudflare.F(load_balancers.LocationStrategyParam{
			Mode:      cloudflare.F(load_balancers.LocationStrategyModePOP),
			PreferECS: cloudflare.F(load_balancers.LocationStrategyPreferECSAlways),
		}),
		Networks: cloudflare.F([]string{"string"}),
		POPPools: cloudflare.F(map[string][]string{
			"LAX": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
			"LHR": {"abd90f38ced07c2e2f4df50b1f61d4194", "f9138c5d07c2e2f4df57b1f61d4196"},
			"SJC": {"00920f38ce07c2e2f4df50b1f61d4194"},
		}),
		Proxied: cloudflare.F(true),
		RandomSteering: cloudflare.F(load_balancers.RandomSteeringParam{
			DefaultWeight: cloudflare.F(0.200000),
			PoolWeights: cloudflare.F(map[string]float64{
				"9290f38c5d07c2e2f4df57b1f61d4196": 0.500000,
				"de90f38ced07c2e2f4df50b1f61d4194": 0.300000,
			}),
		}),
		RegionPools: cloudflare.F(map[string][]string{
			"ENAM": {"00920f38ce07c2e2f4df50b1f61d4194"},
			"WNAM": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
		}),
		Rules: cloudflare.F([]load_balancers.RulesParam{{
			Condition: cloudflare.F("http.request.uri.path contains \"/testing\""),
			Disabled:  cloudflare.F(true),
			FixedResponse: cloudflare.F(load_balancers.RulesFixedResponseParam{
				ContentType: cloudflare.F("application/json"),
				Location:    cloudflare.F("www.example.com"),
				MessageBody: cloudflare.F("Testing Hello"),
				StatusCode:  cloudflare.F(int64(0)),
			}),
			Name: cloudflare.F("route the path /testing to testing datacenter."),
			Overrides: cloudflare.F(load_balancers.RulesOverridesParam{
				AdaptiveRouting: cloudflare.F(load_balancers.AdaptiveRoutingParam{
					FailoverAcrossPools: cloudflare.F(true),
				}),
				CountryPools: cloudflare.F(map[string][]string{
					"GB": {"abd90f38ced07c2e2f4df50b1f61d4194"},
					"US": {"de90f38ced07c2e2f4df50b1f61d4194", "00920f38ce07c2e2f4df50b1f61d4194"},
				}),
				DefaultPools: cloudflare.F([]load_balancers.DefaultPoolsParam{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
				FallbackPool: cloudflare.F("fallback_pool"),
				LocationStrategy: cloudflare.F(load_balancers.LocationStrategyParam{
					Mode:      cloudflare.F(load_balancers.LocationStrategyModePOP),
					PreferECS: cloudflare.F(load_balancers.LocationStrategyPreferECSAlways),
				}),
				POPPools: cloudflare.F(map[string][]string{
					"LAX": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
					"LHR": {"abd90f38ced07c2e2f4df50b1f61d4194", "f9138c5d07c2e2f4df57b1f61d4196"},
					"SJC": {"00920f38ce07c2e2f4df50b1f61d4194"},
				}),
				RandomSteering: cloudflare.F(load_balancers.RandomSteeringParam{
					DefaultWeight: cloudflare.F(0.200000),
					PoolWeights: cloudflare.F(map[string]float64{
						"9290f38c5d07c2e2f4df57b1f61d4196": 0.500000,
						"de90f38ced07c2e2f4df50b1f61d4194": 0.300000,
					}),
				}),
				RegionPools: cloudflare.F(map[string][]string{
					"ENAM": {"00920f38ce07c2e2f4df50b1f61d4194"},
					"WNAM": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
				}),
				SessionAffinity: cloudflare.F(load_balancers.SessionAffinityNone),
				SessionAffinityAttributes: cloudflare.F(load_balancers.SessionAffinityAttributesParam{
					DrainDuration:        cloudflare.F(100.000000),
					Headers:              cloudflare.F([]string{"x"}),
					RequireAllHeaders:    cloudflare.F(true),
					Samesite:             cloudflare.F(load_balancers.SessionAffinityAttributesSamesiteAuto),
					Secure:               cloudflare.F(load_balancers.SessionAffinityAttributesSecureAuto),
					ZeroDowntimeFailover: cloudflare.F(load_balancers.SessionAffinityAttributesZeroDowntimeFailoverNone),
				}),
				SessionAffinityTTL: cloudflare.F(1800.000000),
				SteeringPolicy:     cloudflare.F(load_balancers.SteeringPolicyOff),
				TTL:                cloudflare.F(30.000000),
			}),
			Priority:   cloudflare.F(int64(0)),
			Terminates: cloudflare.F(true),
		}}),
		SessionAffinity: cloudflare.F(load_balancers.SessionAffinityNone),
		SessionAffinityAttributes: cloudflare.F(load_balancers.SessionAffinityAttributesParam{
			DrainDuration:        cloudflare.F(100.000000),
			Headers:              cloudflare.F([]string{"x"}),
			RequireAllHeaders:    cloudflare.F(true),
			Samesite:             cloudflare.F(load_balancers.SessionAffinityAttributesSamesiteAuto),
			Secure:               cloudflare.F(load_balancers.SessionAffinityAttributesSecureAuto),
			ZeroDowntimeFailover: cloudflare.F(load_balancers.SessionAffinityAttributesZeroDowntimeFailoverNone),
		}),
		SessionAffinityTTL: cloudflare.F(1800.000000),
		SteeringPolicy:     cloudflare.F(load_balancers.SteeringPolicyOff),
		TTL:                cloudflare.F(30.000000),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLoadBalancerUpdateWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.LoadBalancers.Update(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		load_balancers.LoadBalancerUpdateParams{
			ZoneID:       cloudflare.F("699d98642c564d2e855e9661899b7252"),
			DefaultPools: cloudflare.F([]load_balancers.DefaultPoolsParam{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
			FallbackPool: cloudflare.F("fallback_pool"),
			Name:         cloudflare.F("www.example.com"),
			AdaptiveRouting: cloudflare.F(load_balancers.AdaptiveRoutingParam{
				FailoverAcrossPools: cloudflare.F(true),
			}),
			CountryPools: cloudflare.F(map[string][]string{
				"GB": {"abd90f38ced07c2e2f4df50b1f61d4194"},
				"US": {"de90f38ced07c2e2f4df50b1f61d4194", "00920f38ce07c2e2f4df50b1f61d4194"},
			}),
			Description: cloudflare.F("Load Balancer for www.example.com"),
			Enabled:     cloudflare.F(true),
			LocationStrategy: cloudflare.F(load_balancers.LocationStrategyParam{
				Mode:      cloudflare.F(load_balancers.LocationStrategyModePOP),
				PreferECS: cloudflare.F(load_balancers.LocationStrategyPreferECSAlways),
			}),
			Networks: cloudflare.F([]string{"string"}),
			POPPools: cloudflare.F(map[string][]string{
				"LAX": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
				"LHR": {"abd90f38ced07c2e2f4df50b1f61d4194", "f9138c5d07c2e2f4df57b1f61d4196"},
				"SJC": {"00920f38ce07c2e2f4df50b1f61d4194"},
			}),
			Proxied: cloudflare.F(true),
			RandomSteering: cloudflare.F(load_balancers.RandomSteeringParam{
				DefaultWeight: cloudflare.F(0.200000),
				PoolWeights: cloudflare.F(map[string]float64{
					"9290f38c5d07c2e2f4df57b1f61d4196": 0.500000,
					"de90f38ced07c2e2f4df50b1f61d4194": 0.300000,
				}),
			}),
			RegionPools: cloudflare.F(map[string][]string{
				"ENAM": {"00920f38ce07c2e2f4df50b1f61d4194"},
				"WNAM": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
			}),
			Rules: cloudflare.F([]load_balancers.RulesParam{{
				Condition: cloudflare.F("http.request.uri.path contains \"/testing\""),
				Disabled:  cloudflare.F(true),
				FixedResponse: cloudflare.F(load_balancers.RulesFixedResponseParam{
					ContentType: cloudflare.F("application/json"),
					Location:    cloudflare.F("www.example.com"),
					MessageBody: cloudflare.F("Testing Hello"),
					StatusCode:  cloudflare.F(int64(0)),
				}),
				Name: cloudflare.F("route the path /testing to testing datacenter."),
				Overrides: cloudflare.F(load_balancers.RulesOverridesParam{
					AdaptiveRouting: cloudflare.F(load_balancers.AdaptiveRoutingParam{
						FailoverAcrossPools: cloudflare.F(true),
					}),
					CountryPools: cloudflare.F(map[string][]string{
						"GB": {"abd90f38ced07c2e2f4df50b1f61d4194"},
						"US": {"de90f38ced07c2e2f4df50b1f61d4194", "00920f38ce07c2e2f4df50b1f61d4194"},
					}),
					DefaultPools: cloudflare.F([]load_balancers.DefaultPoolsParam{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
					FallbackPool: cloudflare.F("fallback_pool"),
					LocationStrategy: cloudflare.F(load_balancers.LocationStrategyParam{
						Mode:      cloudflare.F(load_balancers.LocationStrategyModePOP),
						PreferECS: cloudflare.F(load_balancers.LocationStrategyPreferECSAlways),
					}),
					POPPools: cloudflare.F(map[string][]string{
						"LAX": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
						"LHR": {"abd90f38ced07c2e2f4df50b1f61d4194", "f9138c5d07c2e2f4df57b1f61d4196"},
						"SJC": {"00920f38ce07c2e2f4df50b1f61d4194"},
					}),
					RandomSteering: cloudflare.F(load_balancers.RandomSteeringParam{
						DefaultWeight: cloudflare.F(0.200000),
						PoolWeights: cloudflare.F(map[string]float64{
							"9290f38c5d07c2e2f4df57b1f61d4196": 0.500000,
							"de90f38ced07c2e2f4df50b1f61d4194": 0.300000,
						}),
					}),
					RegionPools: cloudflare.F(map[string][]string{
						"ENAM": {"00920f38ce07c2e2f4df50b1f61d4194"},
						"WNAM": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
					}),
					SessionAffinity: cloudflare.F(load_balancers.SessionAffinityNone),
					SessionAffinityAttributes: cloudflare.F(load_balancers.SessionAffinityAttributesParam{
						DrainDuration:        cloudflare.F(100.000000),
						Headers:              cloudflare.F([]string{"x"}),
						RequireAllHeaders:    cloudflare.F(true),
						Samesite:             cloudflare.F(load_balancers.SessionAffinityAttributesSamesiteAuto),
						Secure:               cloudflare.F(load_balancers.SessionAffinityAttributesSecureAuto),
						ZeroDowntimeFailover: cloudflare.F(load_balancers.SessionAffinityAttributesZeroDowntimeFailoverNone),
					}),
					SessionAffinityTTL: cloudflare.F(1800.000000),
					SteeringPolicy:     cloudflare.F(load_balancers.SteeringPolicyOff),
					TTL:                cloudflare.F(30.000000),
				}),
				Priority:   cloudflare.F(int64(0)),
				Terminates: cloudflare.F(true),
			}}),
			SessionAffinity: cloudflare.F(load_balancers.SessionAffinityNone),
			SessionAffinityAttributes: cloudflare.F(load_balancers.SessionAffinityAttributesParam{
				DrainDuration:        cloudflare.F(100.000000),
				Headers:              cloudflare.F([]string{"x"}),
				RequireAllHeaders:    cloudflare.F(true),
				Samesite:             cloudflare.F(load_balancers.SessionAffinityAttributesSamesiteAuto),
				Secure:               cloudflare.F(load_balancers.SessionAffinityAttributesSecureAuto),
				ZeroDowntimeFailover: cloudflare.F(load_balancers.SessionAffinityAttributesZeroDowntimeFailoverNone),
			}),
			SessionAffinityTTL: cloudflare.F(1800.000000),
			SteeringPolicy:     cloudflare.F(load_balancers.SteeringPolicyOff),
			TTL:                cloudflare.F(30.000000),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLoadBalancerList(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.LoadBalancers.List(context.TODO(), load_balancers.LoadBalancerListParams{
		ZoneID: cloudflare.F("699d98642c564d2e855e9661899b7252"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLoadBalancerDelete(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.LoadBalancers.Delete(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		load_balancers.LoadBalancerDeleteParams{
			ZoneID: cloudflare.F("699d98642c564d2e855e9661899b7252"),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLoadBalancerEditWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.LoadBalancers.Edit(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		load_balancers.LoadBalancerEditParams{
			ZoneID: cloudflare.F("699d98642c564d2e855e9661899b7252"),
			AdaptiveRouting: cloudflare.F(load_balancers.AdaptiveRoutingParam{
				FailoverAcrossPools: cloudflare.F(true),
			}),
			CountryPools: cloudflare.F(map[string][]string{
				"GB": {"abd90f38ced07c2e2f4df50b1f61d4194"},
				"US": {"de90f38ced07c2e2f4df50b1f61d4194", "00920f38ce07c2e2f4df50b1f61d4194"},
			}),
			DefaultPools: cloudflare.F([]load_balancers.DefaultPoolsParam{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
			Description:  cloudflare.F("Load Balancer for www.example.com"),
			Enabled:      cloudflare.F(true),
			FallbackPool: cloudflare.F("fallback_pool"),
			LocationStrategy: cloudflare.F(load_balancers.LocationStrategyParam{
				Mode:      cloudflare.F(load_balancers.LocationStrategyModePOP),
				PreferECS: cloudflare.F(load_balancers.LocationStrategyPreferECSAlways),
			}),
			Name: cloudflare.F("www.example.com"),
			POPPools: cloudflare.F(map[string][]string{
				"LAX": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
				"LHR": {"abd90f38ced07c2e2f4df50b1f61d4194", "f9138c5d07c2e2f4df57b1f61d4196"},
				"SJC": {"00920f38ce07c2e2f4df50b1f61d4194"},
			}),
			Proxied: cloudflare.F(true),
			RandomSteering: cloudflare.F(load_balancers.RandomSteeringParam{
				DefaultWeight: cloudflare.F(0.200000),
				PoolWeights: cloudflare.F(map[string]float64{
					"9290f38c5d07c2e2f4df57b1f61d4196": 0.500000,
					"de90f38ced07c2e2f4df50b1f61d4194": 0.300000,
				}),
			}),
			RegionPools: cloudflare.F(map[string][]string{
				"ENAM": {"00920f38ce07c2e2f4df50b1f61d4194"},
				"WNAM": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
			}),
			Rules: cloudflare.F([]load_balancers.RulesParam{{
				Condition: cloudflare.F("http.request.uri.path contains \"/testing\""),
				Disabled:  cloudflare.F(true),
				FixedResponse: cloudflare.F(load_balancers.RulesFixedResponseParam{
					ContentType: cloudflare.F("application/json"),
					Location:    cloudflare.F("www.example.com"),
					MessageBody: cloudflare.F("Testing Hello"),
					StatusCode:  cloudflare.F(int64(0)),
				}),
				Name: cloudflare.F("route the path /testing to testing datacenter."),
				Overrides: cloudflare.F(load_balancers.RulesOverridesParam{
					AdaptiveRouting: cloudflare.F(load_balancers.AdaptiveRoutingParam{
						FailoverAcrossPools: cloudflare.F(true),
					}),
					CountryPools: cloudflare.F(map[string][]string{
						"GB": {"abd90f38ced07c2e2f4df50b1f61d4194"},
						"US": {"de90f38ced07c2e2f4df50b1f61d4194", "00920f38ce07c2e2f4df50b1f61d4194"},
					}),
					DefaultPools: cloudflare.F([]load_balancers.DefaultPoolsParam{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
					FallbackPool: cloudflare.F("fallback_pool"),
					LocationStrategy: cloudflare.F(load_balancers.LocationStrategyParam{
						Mode:      cloudflare.F(load_balancers.LocationStrategyModePOP),
						PreferECS: cloudflare.F(load_balancers.LocationStrategyPreferECSAlways),
					}),
					POPPools: cloudflare.F(map[string][]string{
						"LAX": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
						"LHR": {"abd90f38ced07c2e2f4df50b1f61d4194", "f9138c5d07c2e2f4df57b1f61d4196"},
						"SJC": {"00920f38ce07c2e2f4df50b1f61d4194"},
					}),
					RandomSteering: cloudflare.F(load_balancers.RandomSteeringParam{
						DefaultWeight: cloudflare.F(0.200000),
						PoolWeights: cloudflare.F(map[string]float64{
							"9290f38c5d07c2e2f4df57b1f61d4196": 0.500000,
							"de90f38ced07c2e2f4df50b1f61d4194": 0.300000,
						}),
					}),
					RegionPools: cloudflare.F(map[string][]string{
						"ENAM": {"00920f38ce07c2e2f4df50b1f61d4194"},
						"WNAM": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
					}),
					SessionAffinity: cloudflare.F(load_balancers.SessionAffinityNone),
					SessionAffinityAttributes: cloudflare.F(load_balancers.SessionAffinityAttributesParam{
						DrainDuration:        cloudflare.F(100.000000),
						Headers:              cloudflare.F([]string{"x"}),
						RequireAllHeaders:    cloudflare.F(true),
						Samesite:             cloudflare.F(load_balancers.SessionAffinityAttributesSamesiteAuto),
						Secure:               cloudflare.F(load_balancers.SessionAffinityAttributesSecureAuto),
						ZeroDowntimeFailover: cloudflare.F(load_balancers.SessionAffinityAttributesZeroDowntimeFailoverNone),
					}),
					SessionAffinityTTL: cloudflare.F(1800.000000),
					SteeringPolicy:     cloudflare.F(load_balancers.SteeringPolicyOff),
					TTL:                cloudflare.F(30.000000),
				}),
				Priority:   cloudflare.F(int64(0)),
				Terminates: cloudflare.F(true),
			}}),
			SessionAffinity: cloudflare.F(load_balancers.SessionAffinityNone),
			SessionAffinityAttributes: cloudflare.F(load_balancers.SessionAffinityAttributesParam{
				DrainDuration:        cloudflare.F(100.000000),
				Headers:              cloudflare.F([]string{"x"}),
				RequireAllHeaders:    cloudflare.F(true),
				Samesite:             cloudflare.F(load_balancers.SessionAffinityAttributesSamesiteAuto),
				Secure:               cloudflare.F(load_balancers.SessionAffinityAttributesSecureAuto),
				ZeroDowntimeFailover: cloudflare.F(load_balancers.SessionAffinityAttributesZeroDowntimeFailoverNone),
			}),
			SessionAffinityTTL: cloudflare.F(1800.000000),
			SteeringPolicy:     cloudflare.F(load_balancers.SteeringPolicyOff),
			TTL:                cloudflare.F(30.000000),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLoadBalancerGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.LoadBalancers.Get(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		load_balancers.LoadBalancerGetParams{
			ZoneID: cloudflare.F("699d98642c564d2e855e9661899b7252"),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
