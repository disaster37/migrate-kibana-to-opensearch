package main

import (
	"fmt"
	"os"
	"sort"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var (
	version string
	commit  string
)

func run(args []string) error {

	// Logger setting
	formatter := new(prefixed.TextFormatter)
	formatter.FullTimestamp = true
	formatter.ForceFormatting = true
	log.SetFormatter(formatter)
	log.SetOutput(os.Stdout)

	// CLI settings
	app := cli.NewApp()
	app.Usage = "CLi to migrate Dashboard from kibana to Opensearch dashboard"
	app.Version = fmt.Sprintf("%s-%s", version, commit)
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "config",
			Usage: "Load configuration from `FILE`",
		},
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:     "kibana-url",
			Usage:    "The kibana URL",
			EnvVars:  []string{"KIBANA_URL"},
			Required: true,
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:     "kibana-user",
			Usage:    "The kibana user",
			EnvVars:  []string{"KIBANA_USER"},
			Required: true,
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:     "kibana-password",
			Usage:    "The Kibana password",
			EnvVars:  []string{"KIBANA_PASSWORD"},
			Required: true,
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:     "dashboard-url",
			Usage:    "The Opensearch dashboard URL",
			EnvVars:  []string{"DASHBOARD_URL"},
			Required: true,
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:     "dashboard-user",
			Usage:    "The Opensearch dashboard user",
			EnvVars:  []string{"DASHBOARD_USER"},
			Required: true,
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:     "dashboard-password",
			Usage:    "The Opensearch dashboard password",
			EnvVars:  []string{"DASHBOARD_PASSWORD"},
			Required: true,
		}),
		&cli.BoolFlag{
			Name:  "self-signed-certificate",
			Usage: "Disable the TLS certificate check",
		},
		&cli.BoolFlag{
			Name:  "debug",
			Usage: "Display debug output",
		},
	}
	app.Commands = []*cli.Command{
		{
			Name:  "migrate-dashboard",
			Usage: "Migrate dahsboards from Kibana to Opensearch Dashboard",
			Flags: []cli.Flag{
				&cli.StringSliceFlag{
					Name:  "dashboard-id",
					Usage: "The dashboard ids. If not provided, it migrate all dashboards",
				},
				&cli.StringFlag{
					Name:  "space",
					Usage: "The Kibana space where export dahsboards. If not provided is use default/global space/tenant",
				},
			},
			Action: migrateDashboard,
		},
	}

	app.Before = func(c *cli.Context) error {

		if c.Bool("debug") {
			log.SetLevel(log.DebugLevel)
		}

		if c.String("config") != "" {
			before := altsrc.InitInputSourceWithContext(app.Flags, altsrc.NewYamlSourceFromFlagFunc("config"))
			return before(c)
		}
		return nil
	}

	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(args)
	return err
}

func main() {
	err := run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
