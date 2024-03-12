package opensearch

import (
	"crypto/tls"
	"net/http"

	"github.com/disaster37/opensearch/v2"
	"github.com/disaster37/opensearch/v2/config"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"k8s.io/utils/ptr"
)

func manageOpensearchGlobalParameters(c *cli.Context) (*opensearch.Client, error) {

	log.Debug("Opensearch URL: ", c.StringSlice("dashboard-url"))
	log.Debug("Opensearch user: ", c.String("dashboard-user"))
	log.Debug("Opensearch password: XXX")
	log.Debug("Disable verify SSL: ", c.Bool("self-signed-certificate"))

	// Init opensearch client
	cfg := &config.Config{
		URLs:        c.StringSlice("dashboard-urls"),
		Username:    c.String("dashboard-user"),
		Password:    c.String("dashboard-password"),
		Sniff:       ptr.To[bool](false),
		Healthcheck: ptr.To[bool](false),
	}
	if c.Bool("self-signed-certificate") {
		cfg.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
	}

	osClient, err := opensearch.NewClientFromConfig(cfg)
	if err != nil {
		return nil, err
	}

	return osClient, nil

}
