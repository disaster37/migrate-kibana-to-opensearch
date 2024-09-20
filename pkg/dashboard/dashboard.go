package dashboard

import (
	opensearchdashboard "github.com/disaster37/opensearch-dashboard/v2"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func ManageOpensearchGlobalParameters(c *cli.Context) (opensearchdashboard.Client, error) {

	log.Debug("Opensearch URL: ", c.StringSlice("dashboard-url"))
	log.Debug("Opensearch user: ", c.String("dashboard-user"))
	log.Debug("Opensearch password: XXX")
	log.Debug("Disable verify SSL: ", c.Bool("self-signed-certificate"))

	cfg := opensearchdashboard.Config{
		Address:          c.String("dashboard-url"),
		Username:         c.String("dashboard-user"),
		Password:         c.String("dashboard-password"),
		DisableVerifySSL: true,
	}

	dbClient, err := opensearchdashboard.NewClient(cfg)

	if err != nil {
		return nil, err
	}

	return dbClient, nil

}
