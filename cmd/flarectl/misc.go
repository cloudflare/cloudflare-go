package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

func initializeAPI(c *cli.Context) error {
	apiKey := os.Getenv("CF_API_KEY")
	if apiKey == "" {
		err := errors.New("No CF_API_KEY environment set")
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	apiEmail := os.Getenv("CF_API_EMAIL")
	if apiEmail == "" {
		err := errors.New("No CF_API_EMAIL environment set")
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	// Be aware the following code sets the global package `api` variable
	var err error
	api, err = cloudflare.New(apiKey, apiEmail)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cloudflare api: %s", err)
		return err
	}

	if c.IsSet("accountid") {
		cloudflare.UsingAccount(c.String("accountid"))(api)
	}

	return nil
}

// writeTableTabular outputs tabular data to STDOUT
func writeTableTabular(data [][]string, cols ...string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(cols)
	table.SetBorder(false)
	table.AppendBulk(data)

	table.Render()
}

// writeTableJSON outputs JSON data to STDOUT
func writeTableJSON(data [][]string, cols ...string) {
	mappedData := make([]map[string]string, 0)
	for i := range data {
		rowData := make(map[string]string)
		for j := range data[i] {
			rowData[cols[j]] = data[i][j]
		}
		mappedData = append(mappedData, rowData)
	}
	jsonData, err := json.Marshal(mappedData)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonData))
}

// writeTable outputs JSON or tabular data to STDOUT
func writeTable(c *cli.Context, data [][]string, cols ...string) {
	if c.GlobalBool("json") {
		writeTableJSON(data, cols...)
	} else {
		writeTableTabular(data, cols...)
	}
}

// Utility function to check if CLI flags were given.
func checkFlags(c *cli.Context, flags ...string) error {
	for _, flag := range flags {
		if c.String(flag) == "" {
			cli.ShowSubcommandHelp(c)
			err := errors.Errorf("error: the required flag %q was empty or not provided", flag)
			fmt.Fprintln(os.Stderr, err)
			return err
		}
	}

	return nil
}

func ips(c *cli.Context) {

	if c.String("ip-type") == "all" {
		_getIps("ipv4", c.Bool("ip-only"))
		_getIps("ipv6", c.Bool("ip-only"))
	} else {
		_getIps(c.String("ip-type"), c.Bool("ip-only"))
	}
}

//
// gets type of IPs to retrieve and returns results
//
func _getIps(ipType string, showMsgType bool) {
	ips, _ := cloudflare.IPs()

	switch ipType {
	case "ipv4":
		if showMsgType != true {
			fmt.Println("IPv4 ranges:")
		}
		for _, r := range ips.IPv4CIDRs {
			fmt.Println(" ", r)
		}
	case "ipv6":
		if showMsgType != true {
			fmt.Println("IPv6 ranges:")
		}
		for _, r := range ips.IPv6CIDRs {
			fmt.Println(" ", r)
		}
	}
}

func userInfo(c *cli.Context) {
	user, err := api.UserDetails()
	if err != nil {
		fmt.Println(err)
		return
	}
	var output [][]string
	output = append(output, []string{
		user.ID,
		user.Email,
		user.Username,
		user.FirstName + " " + user.LastName,
		fmt.Sprintf("%t", user.TwoFA),
	})
	writeTable(c, output, "ID", "Email", "Username", "Name", "2FA")
}

func userUpdate(*cli.Context) {
}

func pageRules(c *cli.Context) {
	if err := checkFlags(c, "zone"); err != nil {
		return
	}
	zone := c.String("zone")

	zoneID, err := api.ZoneIDByName(zone)
	if err != nil {
		fmt.Println(err)
		return
	}

	rules, err := api.ListPageRules(zoneID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%3s %-32s %-8s %s\n", "Pri", "ID", "Status", "URL")
	for _, r := range rules {
		var settings []string
		fmt.Printf("%3d %s %-8s %s\n", r.Priority, r.ID, r.Status, r.Targets[0].Constraint.Value)
		for _, a := range r.Actions {
			var s string
			switch v := a.Value.(type) {
			case int:
				s = fmt.Sprintf("%s: %d", cloudflare.PageRuleActions[a.ID], v)
			case float64:
				s = fmt.Sprintf("%s: %.f", cloudflare.PageRuleActions[a.ID], v)
			case map[string]interface{}:
				s = fmt.Sprintf("%s: %.f - %s", cloudflare.PageRuleActions[a.ID], v["status_code"], v["url"])
			case nil:
				s = fmt.Sprintf("%s", cloudflare.PageRuleActions[a.ID])
			default:
				vs := fmt.Sprintf("%s", v)
				s = fmt.Sprintf("%s: %s", cloudflare.PageRuleActions[a.ID], strings.Title(strings.Replace(vs, "_", " ", -1)))
			}
			settings = append(settings, s)
		}
		fmt.Println("   ", strings.Join(settings, ", "))
	}
}

func railgun(*cli.Context) {
}
