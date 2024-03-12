package migrate

import (
	"encoding/json"
	"strings"

	"emperror.dev/errors"
	log "github.com/sirupsen/logrus"
	"golang.org/x/mod/semver"
)

func ConvertObjectFromElasticsearchToOpensearch(data []byte) (dataConverted []byte, err error) {
	o := new(Object)
	if err = json.Unmarshal(data, o); err != nil {
		return nil, errors.Wrap(err, "Error when unmarshall exported object")
	}

	switch o.Type {
	case "lens":
		log.Warnf("Found object '%s' of type 'lens'. We exclude it from import because is not supported on Opensearch", o.Id)
		return nil, nil
	case "map":
		log.Warnf("Found object '%s' of type 'map'. We exclude it from import because is not supported on Opensearch", o.Id)
		return nil, nil
	case "index-pattern":
		if semver.Compare("7.6.0", o.MigrationVersion.IndexPattern) == -1 {
			log.Infof("Patch version for object '%s' of type '%s'", o.Id, o.Type)
			o.MigrationVersion.IndexPattern = "7.6.0"
			o.CoreMigrationVersion = "7.6.0"
		}

		indexes := strings.Split(o.Attributes["title"].(string), ",")
		if len(indexes) > 1 {
			log.Infof("Patch title for object '%s' of type '%s'. Multiple index is not supported on Opensearch. We keep only the first entry '%s'", o.Id, o.Type, indexes[0])
			o.Attributes["title"] = indexes[0]
		}
	case "search":
		if semver.Compare("7.6.0", o.MigrationVersion.Search) == -1 {
			log.Infof("Patch version for object '%s' of type '%s'", o.Id, o.Type)
			o.MigrationVersion.IndexPattern = "7.6.0"
			o.CoreMigrationVersion = "7.6.0"
		}
	case "visualization":
		if semver.Compare("7.10.0", o.MigrationVersion.Search) == -1 {
			log.Infof("Patch version for object '%s' of type '%s'", o.Id, o.Type)
			o.MigrationVersion.IndexPattern = "7.10.0"
			o.CoreMigrationVersion = "7.10.0"
		}
	case "dashboard":
		if semver.Compare("7.9.3", o.MigrationVersion.Search) == -1 {
			log.Infof("Patch version for object '%s' of type '%s'", o.Id, o.Type)
			o.MigrationVersion.IndexPattern = "7.9.3"
			o.CoreMigrationVersion = "7.9.3"
		}
	default:
		return nil, errors.Errorf("The object '%s' is on unsupported type '%s'", o.Id, o.Type)
	}

	dataConverted, err = json.Marshal(o)
	if err != nil {
		return nil, errors.Wrap(err, "Error when Marshall exported object")
	}

	return dataConverted, err
}
