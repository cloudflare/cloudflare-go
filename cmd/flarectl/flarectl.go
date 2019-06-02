package main

import (
	"os"

	"github.com/cloudflare/cloudflare-go"
	"github.com/codegangsta/cli"
)

var api *cloudflare.API

func main() {
	app := cli.NewApp()
	app.Name = "flarectl"
	app.Usage = "Cloudflare CLI"
	app.Version = "2019.06.01"
	app.Before = authenticate

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "api-key, k",
			Usage:  "api key",
			EnvVar: "CF_API_KEY",
		},
		cli.StringFlag{
			Name:   "email, e",
			Usage:  "account email",
			EnvVar: "CF_ACCOUNT_EMAIL, CF_API_EMAIL",
		},
		cli.StringFlag{
			Name:   "account-id, a",
			Usage:  "account id (previously: organization id)",
			EnvVar: "CF_ACCOUNT_ID",
		},
		cli.StringFlag{
			Name:   "zone, z",
			Usage:  "zone name",
			EnvVar: "CF_ZONE_NAME",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "ips",
			Aliases: []string{"i"},
			Action:  getIps,
			Usage:   "Print Cloudflare IP ranges",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "ip-type",
					Usage: "type of IPs ( ipv4 | ipv6 | all )",
					Value: "all",
				},
				cli.BoolFlag{
					Name:  "ip-only",
					Usage: "show only addresses",
				},
			},
		},
		{
			Name:    "user",
			Aliases: []string{"u"},
			Usage:   "User information",
			Subcommands: []cli.Command{
				{
					Name:    "info",
					Aliases: []string{"i"},
					Action:  userInfo,
					Usage:   "User details",
				},
				{
					Name:    "update",
					Aliases: []string{"u"},
					Action:  userUpdate,
					Usage:   "Update user details",
				},
			},
		},

		{
			Name:    "zone",
			Aliases: []string{"z"},
			Usage:   "Zone information",
			Subcommands: []cli.Command{
				{
					Name:    "list",
					Aliases: []string{"l"},
					Action:  zoneList,
					Usage:   "List all zones on an account",
				},
				{
					Name:    "create",
					Aliases: []string{"c"},
					Action:  zoneCreate,
					Usage:   "Create a new zone",
					Flags: []cli.Flag{
						cli.BoolFlag{
							Name:  "jumpstart",
							Usage: "automatically fetch DNS records",
						},
					},
				},
				{
					Name:   "check",
					Action: zoneCheck,
					Usage:  "Initiate a zone activation check",
				},
				{
					Name:    "info",
					Aliases: []string{"i"},
					Action:  zoneInfo,
					Usage:   "Information on one zone",
				},
				{
					Name:    "plan",
					Aliases: []string{"p"},
					Action:  zonePlan,
					Usage:   "Plan information for one zone",
				},
				{
					Name:    "settings",
					Aliases: []string{"s"},
					Action:  zoneSettings,
					Usage:   "Settings for one zone",
				},
				{
					Name:   "purge",
					Action: zoneCachePurge,
					Usage:  "(Selectively) Purge the cache for a zone",
					Flags: []cli.Flag{
						cli.BoolFlag{
							Name:  "everything",
							Usage: "purge everything from cache for the zone",
						},
						cli.StringSliceFlag{
							Name:  "hosts",
							Usage: "a list of hostnames to purge the cache for",
						},
						cli.StringSliceFlag{
							Name:  "tags",
							Usage: "the cache tags to purge (Enterprise only)",
						},
						cli.StringSliceFlag{
							Name:  "files",
							Usage: "a list of [exact] URLs to purge",
						},
					},
				},
				{
					Name:    "dns",
					Aliases: []string{"d"},
					Action:  zoneRecords,
					Usage:   "DNS records for a zone",
				},
				{
					Name:    "railgun",
					Aliases: []string{"r"},
					Action:  zoneRailgun,
					Usage:   "Railguns for a zone",
				},
				{
					Name:    "certs",
					Aliases: []string{"c"},
					Action:  zoneCerts,
					Usage:   "Custom SSL certificates for a zone",
				},
				{
					Name:    "keyless",
					Aliases: []string{"k"},
					Action:  zoneKeyless,
					Usage:   "Keyless SSL for a zone",
				},
			},
		},

		{
			Name:    "dns",
			Aliases: []string{"d"},
			Usage:   "DNS records",
			Subcommands: []cli.Command{
				{
					Name:    "list",
					Aliases: []string{"l"},
					Action:  zoneRecords,
					Usage:   "List DNS records for a zone",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "id",
							Usage: "record id",
						},
						cli.StringFlag{
							Name:  "name",
							Usage: "record name",
						},
						cli.StringFlag{
							Name:  "content",
							Usage: "record content",
						},
					},
				},
				{
					Name:    "create",
					Aliases: []string{"c"},
					Action:  dnsCreate,
					Usage:   "Create a DNS record",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "name",
							Usage: "record name",
						},
						cli.StringFlag{
							Name:  "type",
							Usage: "record type",
						},
						cli.StringFlag{
							Name:  "content",
							Usage: "record content",
						},
						cli.IntFlag{
							Name:  "ttl",
							Usage: "TTL (1 = automatic)",
							Value: 1,
						},
						cli.BoolFlag{
							Name:  "proxy",
							Usage: "proxy through Cloudflare (orange cloud)",
						},
					},
				},
				{
					Name:    "update",
					Aliases: []string{"u"},
					Action:  dnsUpdate,
					Usage:   "Update a DNS record",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "id",
							Usage: "record id",
						},
						cli.StringFlag{
							Name:  "name",
							Usage: "record name",
						},
						cli.StringFlag{
							Name:  "content",
							Usage: "record content",
						},
						cli.IntFlag{
							Name:  "ttl",
							Usage: "TTL (1 = automatic)",
							Value: 1,
						},
						cli.BoolFlag{
							Name:  "proxy",
							Usage: "proxy through Cloudflare (orange cloud)",
						},
					},
				},
				{
					Name:    "create-or-update",
					Aliases: []string{"o"},
					Action:  dnsCreateOrUpdate,
					Usage:   "Create a DNS record, or update if it exists",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "name",
							Usage: "record name",
						},
						cli.StringFlag{
							Name:  "content",
							Usage: "record content",
						},
						cli.StringFlag{
							Name:  "type",
							Usage: "record type",
						},
						cli.IntFlag{
							Name:  "ttl",
							Usage: "TTL (1 = automatic)",
							Value: 1,
						},
						cli.BoolFlag{
							Name:  "proxy",
							Usage: "proxy through Cloudflare (orange cloud)",
						},
					},
				},
				{
					Name:    "delete",
					Aliases: []string{"d"},
					Action:  dnsDelete,
					Usage:   "Delete a DNS record",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "id",
							Usage: "record id",
						},
					},
				},
			},
		},
		{
			Name:    "user-agents",
			Aliases: []string{"ua"},
			Usage:   "User-Agent blocking",
			Subcommands: []cli.Command{
				{
					Name:    "list",
					Aliases: []string{"l"},
					Action:  userAgentList,
					Usage:   "List User-Agent blocks for a zone",
					Flags: []cli.Flag{
						cli.IntFlag{
							Name:  "page",
							Usage: "result page to return",
						},
					},
				},
				{
					Name:    "create",
					Aliases: []string{"c"},
					Action:  userAgentCreate,
					Usage:   "Create a User-Agent blocking rule",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "mode",
							Usage: "the blocking mode: block, challenge, js_challenge, whitelist",
						},
						cli.StringFlag{
							Name:  "value",
							Usage: "the exact User-Agent to block",
						},
						cli.BoolFlag{
							Name:  "paused",
							Usage: "whether the rule should be paused (default: false)",
						},
						cli.StringFlag{
							Name:  "description",
							Usage: "a description for the rule",
						},
					},
				},
				{
					Name:    "update",
					Aliases: []string{"u"},
					Action:  userAgentUpdate,
					Usage:   "Update an existing User-Agent block",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "id",
							Usage: "User-Agent blocking rule ID",
						},
						cli.StringFlag{
							Name:  "mode",
							Usage: "the blocking mode: block, challenge, js_challenge, whitelist",
						},
						cli.StringFlag{
							Name:  "value",
							Usage: "the exact User-Agent to block",
						},
						cli.BoolFlag{
							Name:  "paused",
							Usage: "whether the rule should be paused (default: false)",
						},
						cli.StringFlag{
							Name:  "description",
							Usage: "a description for the rule",
						},
					},
				},
				{
					Name:    "delete",
					Aliases: []string{"d"},
					Action:  userAgentDelete,
					Usage:   "Delete a User-Agent block",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "id",
							Usage: "User-Agent blocking rule ID",
						},
					},
				},
			},
		},
		{
			Name:    "pagerules",
			Aliases: []string{"p"},
			Usage:   "Page Rules",
			Subcommands: []cli.Command{
				{
					Name:    "list",
					Aliases: []string{"l"},
					Action:  pageRules,
					Usage:   "List Page Rules for a zone",
				},
			},
		},

		{
			Name:    "key-value",
			Aliases: []string{"kv"},
			Usage:   "Workers KV",
			Subcommands: []cli.Command{
				{
					Name:    "write",
					Aliases: []string{"w"},
					Action:  kvWrite,
					Usage:   "Write a key-value pair to a namespace",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "namespace, n",
							Usage: "name or ID of the namespace",
						},
						cli.StringFlag{
							Name:  "key, k",
							Usage: "key name of the pair",
						},
						cli.StringFlag{
							Name:  "value, v",
							Usage: "value of the pair",
						},
						cli.IntFlag{
							Name:  "expiration, e",
							Usage: "expiration of the pair",
						},
						cli.DurationFlag{
							Name:  "ttl, t",
							Usage: "expiration ttl of the pair",
						},
					},
				},
				{
					Name:    "upload",
					Aliases: []string{"u"},
					Action:  kvUpload,
					Usage:   "Upload file(s) to a namespace",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "namespace, n",
							Usage: "name or ID of the namespace",
						},
						cli.StringFlag{
							Name:  "key, k",
							Usage: "format string to make keys from their path",
							Value: "%s",
						},
						cli.StringFlag{
							Name:  "filter, f",
							Usage: "regex expression to filter for paths",
							Value: ".*",
						},
						cli.StringFlag{
							Name:  "path, p",
							Usage: "path of the file or directory to upload",
						},
						cli.BoolFlag{
							Name:  "recursive, r",
							Usage: "whether to search for files recursively",
						},
						cli.IntFlag{
							Name:  "expiration, e",
							Usage: "expiration of each file",
						},
						cli.DurationFlag{
							Name:  "ttl, t",
							Usage: "expiration ttl of each file",
						},
						cli.BoolFlag{
							Name:  "unsafe, u",
							Usage: "whether to NOT check for max upload limits",
						},
					},
				},
				{
					Name:    "read",
					Aliases: []string{"r"},
					Action:  kvRead,
					Usage:   "Read key-value pair(s) from a namespace",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "namespace, n",
							Usage: "name or ID of the namespace",
						},
						cli.StringSliceFlag{
							Name:  "key, k",
							Usage: "key name of the pair(s)",
						},
						cli.BoolFlag{
							Name:  "raw, r",
							Usage: "whether to print the raw value",
						},
					},
				},
				{
					Name:    "delete",
					Aliases: []string{"d"},
					Action:  kvDelete,
					Usage:   "Delete key-value pair(s) from a namespace",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "namespace, n",
							Usage: "name or ID of the namespace",
						},
						cli.StringSliceFlag{
							Name:  "key, k",
							Usage: "key name of the pair(s)",
						},
					},
				}, {
					Name:    "list",
					Aliases: []string{"l"},
					Action:  kvList,
					Usage:   "List key-value pairs in a namespace",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "namespace, n",
							Usage: "name or ID of the namespace",
						},
						cli.StringFlag{
							Name:  "cursor, c",
							Usage: "cursor from the previous list operation",
							Value: "",
						},
						cli.IntFlag{
							Name:  "limit, l",
							Usage: "maximum number of pairs to list",
							Value: 20,
						},
						cli.BoolFlag{
							Name:  "values, v",
							Usage: "whether to also fetch values",
						},
					},
				},
			},
		},

		{
			Name:    "railgun",
			Aliases: []string{"r"},
			Usage:   "Railgun information",
			Action:  railgun,
		},

		{
			Name:    "firewall",
			Aliases: []string{"f"},
			Usage:   "Firewall",
			Subcommands: []cli.Command{
				{
					Name:    "rules",
					Aliases: []string{"r"},
					Usage:   "Access Rules",
					Subcommands: []cli.Command{
						{
							Name:    "list",
							Aliases: []string{"l"},
							Action:  firewallAccessRules,
							Usage:   "List firewall access rules",
							Flags: []cli.Flag{
								cli.StringFlag{
									Name:  "value",
									Usage: "rule value",
								},
								cli.StringFlag{
									Name:  "scope-type",
									Usage: "rule scope",
								},
								cli.StringFlag{
									Name:  "mode",
									Usage: "rule mode",
								},
								cli.StringFlag{
									Name:  "notes",
									Usage: "rule notes",
								},
							},
						},
						{
							Name:    "create",
							Aliases: []string{"c"},
							Action:  firewallAccessRuleCreate,
							Usage:   "Create a firewall access rule",
							Flags: []cli.Flag{
								cli.StringFlag{
									Name:  "value",
									Usage: "rule value",
								},
								cli.StringFlag{
									Name:  "mode",
									Usage: "rule mode",
								},
								cli.StringFlag{
									Name:  "notes",
									Usage: "rule notes",
								},
							},
						},
						{
							Name:    "update",
							Aliases: []string{"u"},
							Action:  firewallAccessRuleUpdate,
							Usage:   "Update a firewall access rule",
							Flags: []cli.Flag{
								cli.StringFlag{
									Name:  "id",
									Usage: "rule id",
								},
								cli.StringFlag{
									Name:  "mode",
									Usage: "rule mode",
								},
								cli.StringFlag{
									Name:  "notes",
									Usage: "rule notes",
								},
							},
						},
						{
							Name:    "create-or-update",
							Aliases: []string{"o"},
							Action:  firewallAccessRuleCreateOrUpdate,
							Usage:   "Creatae a firewall access rule, or update it if it exists",
							Flags: []cli.Flag{
								cli.StringFlag{
									Name:  "value",
									Usage: "rule value",
								},
								cli.StringFlag{
									Name:  "mode",
									Usage: "rule mode",
								},
								cli.StringFlag{
									Name:  "notes",
									Usage: "rule notes",
								},
							},
						},
						{
							Name:    "delete",
							Aliases: []string{"d"},
							Action:  firewallAccessRuleDelete,
							Usage:   "Delete a firewall access rule",
							Flags: []cli.Flag{
								cli.StringFlag{
									Name:  "id",
									Usage: "rule id",
								},
							},
						},
					},
				},
			},
		},
	}
	app.Run(os.Args)
}
