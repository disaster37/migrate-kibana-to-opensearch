package dashboard

import (
	"github.com/disaster37/go-kibana-rest/v8"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func ManageOpensearchGlobalParameters(c *cli.Context) (*kibana.Client, error) {

	log.Debug("Opensearch URL: ", c.StringSlice("dashboard-url"))
	log.Debug("Opensearch user: ", c.String("dashboard-user"))
	log.Debug("Opensearch password: XXX")
	log.Debug("Disable verify SSL: ", c.Bool("self-signed-certificate"))

	cfg := kibana.Config{
		Address:          c.String("dashboard-url"),
		Username:         c.String("dashboard-user"),
		Password:         c.String("dashboard-password"),
		DisableVerifySSL: true,
	}

	kbClient, err := kibana.NewClient(cfg)

	if err != nil {
		return nil, err
	}

	return kbClient, nil

}
