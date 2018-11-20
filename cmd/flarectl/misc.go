package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"io/ioutil"

	"github.com/cloudflare/cloudflare-go"
	"github.com/codegangsta/cli"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/olekukonko/tablewriter"
	"gopkg.in/yaml.v2"

	"github.com/pkg/errors"
)

type conf struct {
	Credentials       []credential `yaml:"credentials"`
	CurrentCredential struct {
		Email string `yaml:"email"`
	} `yaml:"current-credential"`
}

type credential struct {
	Email string `yaml:"email"`
	Key   string `yaml:"key"`
}

// writeTable outputs tabular data to stdout.
func writeTable(data [][]string, cols ...string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(cols)
	table.SetBorder(false)
	table.AppendBulk(data)

	table.Render()
}

func checkEnv() error {
	home, _ := homedir.Dir()
	cfgFile := home + "/.flarectl"

	var APIKey string
	var APIEmail string

	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		APIKey = os.Getenv("CF_API_KEY")
		APIEmail = os.Getenv("CF_API_EMAIL")
		if APIKey == "" {
			return errors.New("API key not defined")
		}
		if APIEmail == "" {
			return errors.New("API email not defined")
		}

	} else {
		var fContext conf
		yamlFile, err := ioutil.ReadFile(cfgFile)
		if err != nil {
			log.Printf("yamlFile.Get err   #%v ", err)
		}

		err = yaml.Unmarshal(yamlFile, &fContext)
		if err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}

		for _, cred := range fContext.Credentials {
			if cred.Email == fContext.CurrentCredential.Email {
				APIEmail = cred.Email
				APIKey = cred.Key
			}
		}
	}

	if api == nil {
		var err error
		api, err = cloudflare.New(APIKey, APIEmail)
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
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

func userInfo(*cli.Context) {
	if err := checkEnv(); err != nil {
		fmt.Println(err)
		return
	}
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
	writeTable(output, "ID", "Email", "Username", "Name", "2FA")
}

func userUpdate(*cli.Context) {
}

func pageRules(c *cli.Context) {
	if err := checkEnv(); err != nil {
		fmt.Println(err)
		return
	}
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
