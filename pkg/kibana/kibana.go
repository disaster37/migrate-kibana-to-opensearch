package kibana

import (
	"github.com/disaster37/go-kibana-rest/v8"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func ManageKibanaGlobalParameters(c *cli.Context) (*kibana.Client, error) {

	log.Debug("Kibana URL: ", c.String("kibana-url"))
	log.Debug("Elasticsearch user: ", c.String("kibana-user"))
	log.Debug("Elasticsearch password: XXX")
	log.Debug("Disable verify SSL: ", c.Bool("self-signed-certificate"))

	cfg := kibana.Config{
		Address:          c.String("kibana-url"),
		Username:         c.String("kibana-user"),
		Password:         c.String("kibana-password"),
		DisableVerifySSL: true,
	}

	kbClient, err := kibana.NewClient(cfg)

	if err != nil {
		return nil, err
	}

	return kbClient, nil

}
