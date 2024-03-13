package kibana

import (
	"bytes"

	"emperror.dev/errors"
	"github.com/disaster37/go-kibana-rest/v8"
	log "github.com/sirupsen/logrus"
)

func ExportDashboards(dashboards []string, userSpace string, kbClient *kibana.Client) (data *bytes.Buffer, err error) {
	if kbClient == nil {
		return nil, errors.New("You must provide kb client")
	}
	log.Debug("Dashboards: ", dashboards)
	log.Debug("UserSpace: ", userSpace)

	var (
		objects     []map[string]string
		objectTypes []string
	)

	if len(dashboards) > 0 {
		objects = make([]map[string]string, 0, len(dashboards))
		for _, dashboardId := range dashboards {
			objects = append(objects, map[string]string{
				"type": "dashboard",
				"id":   dashboardId,
			})
		}
	} else {
		objectTypes = []string{"dashboard"}
	}

	// Exports all dashboard and includes all references
	exportByte, err := kbClient.API.KibanaSavedObject.Export(
		objectTypes,
		objects,
		true,
		userSpace,
	)
	if err != nil {
		return nil, errors.Wrap(err, "Error when export dashboards from Kibana")
	}

	exportBuffer := bytes.NewBuffer(exportByte)

	return exportBuffer, nil
}
