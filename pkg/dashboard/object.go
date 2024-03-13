package dashboard

import (
	"encoding/json"

	"emperror.dev/errors"
	"github.com/disaster37/go-kibana-rest/v8"
	log "github.com/sirupsen/logrus"
)

func ImportDashboards(data []byte, tenant string, kbClient *kibana.Client) (err error) {

	if tenant == "" || tenant == "default" {
		kbClient.Client.SetHeader("securitytenant", "global")
	} else {
		kbClient.Client.SetHeader("securitytenant", tenant)
	}
	res, err := kbClient.API.KibanaSavedObject.Import(data, true, "")
	if err != nil {
		return errors.Wrap(err, "Error when import objects on Opensearch")
	}

	b, err := json.Marshal(res)
	if err != nil {
		return err
	}

	log.Info(string(b))

	return nil

}
