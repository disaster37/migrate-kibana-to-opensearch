package dashboard

import (
	"encoding/json"

	"emperror.dev/errors"
	opensearchdashboard "github.com/disaster37/opensearch-dashboard/v2"
	"github.com/disaster37/opensearch-dashboard/v2/api"
	log "github.com/sirupsen/logrus"
)

func ImportDashboards(data []byte, tenant string, dbClient opensearchdashboard.Client) (err error) {

	res, err := dbClient.SavedObject().Import(tenant, api.SavedObjectImportOption{Overwrite: true}, data)
	if err != nil {
		return errors.Wrap(err, "Error when import objects on Opensearch")
	}

	b, err := json.Marshal(res)
	if err != nil {
		return err
	}

	if _, ok := res["errors"]; ok {
		log.Error(string(b))
		return errors.New("Error when import dashboards on Opensearch")
	}

	log.Info(string(b))

	return nil

}
