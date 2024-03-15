package migrate

import (
	"encoding/json"
	"regexp"
	"strings"

	"emperror.dev/errors"
	"github.com/hashicorp/go-version"
	log "github.com/sirupsen/logrus"
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
		currentVersion, err := version.NewVersion(o.MigrationVersion.IndexPattern)
		if err != nil {
			return nil, err
		}
		maxVersion, err := version.NewVersion("7.6.0")
		if err != nil {
			return nil, err
		}
		if currentVersion.GreaterThan(maxVersion) {
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
		currentVersion, err := version.NewVersion(o.MigrationVersion.Search)
		if err != nil {
			return nil, err
		}
		maxVersion, err := version.NewVersion("7.9.3")
		if err != nil {
			return nil, err
		}
		if currentVersion.GreaterThan(maxVersion) {
			log.Infof("Patch version for object '%s' of type '%s'", o.Id, o.Type)
			o.MigrationVersion.Search = "7.9.3"
			o.CoreMigrationVersion = "7.9.3"

			// Remove grid property
			if _, ok := o.Attributes["grid"]; ok {
				delete(o.Attributes, "grid")
				log.Infof("Remove 'grid' for object '%s' of type '%s'", o.Id, o.Type)
			}

			// Remove hideChart property
			if _, ok := o.Attributes["hideChart"]; ok {
				delete(o.Attributes, "hideChart")
				log.Infof("Remove 'hideChart' for object '%s' of type '%s'", o.Id, o.Type)
			}

			// Remove isTextBasedQuery property
			if _, ok := o.Attributes["isTextBasedQuery"]; ok {
				delete(o.Attributes, "isTextBasedQuery")
				log.Infof("Remove 'isTextBasedQuery' for object '%s' of type '%s'", o.Id, o.Type)
			}

			// Remove timeRestore property
			if _, ok := o.Attributes["timeRestore"]; ok {
				delete(o.Attributes, "timeRestore")
				log.Infof("Remove 'timeRestore' for object '%s' of type '%s'", o.Id, o.Type)
			}

			// Remove usesAdHocDataView property
			if _, ok := o.Attributes["usesAdHocDataView"]; ok {
				delete(o.Attributes, "usesAdHocDataView")
				log.Infof("Remove 'usesAdHocDataView' for object '%s' of type '%s'", o.Id, o.Type)
			}
		}

	case "visualization":
		currentVersion, err := version.NewVersion(o.MigrationVersion.Visualization)
		if err != nil {
			return nil, err
		}
		maxVersion, err := version.NewVersion("7.10.0")
		if err != nil {
			return nil, err
		}
		if currentVersion.GreaterThan(maxVersion) {
			log.Infof("Patch version for object '%s' of type '%s'", o.Id, o.Type)
			o.MigrationVersion.Visualization = "7.10.0"
			o.CoreMigrationVersion = "7.10.0"
		}
	case "dashboard":
		currentVersion, err := version.NewVersion(o.MigrationVersion.Dashboard)
		if err != nil {
			return nil, err
		}
		maxVersion, err := version.NewVersion("7.9.3")
		if err != nil {
			return nil, err
		}
		if currentVersion.GreaterThan(maxVersion) {
			log.Infof("Patch version for object '%s' of type '%s'", o.Id, o.Type)
			o.MigrationVersion.Dashboard = "7.9.3"
			o.CoreMigrationVersion = "7.9.3"

			// We need to convert the name for somes references types and remove not supported object
			r := regexp.MustCompile(`^[^:]+:([^:]+)$`)
			finalReference := make([]map[string]any, 0, len(o.References))
			for _, reference := range o.References {
				if reference["type"] == "visualization" || reference["type"] == "search" {
					matches := r.FindStringSubmatch(reference["name"].(string))
					if len(matches) > 1 {
						reference["name"] = matches[1]
						log.Debugf("Patch reference name with new name '%s'", matches[1])
					}
				}
				if reference["type"] == "lens" || reference["type"] == "map" {

					log.Warnf("We remove reference of object '%s' because the type '%s' is not supported on Opensearch", reference["id"], reference["type"])
				} else {
					finalReference = append(finalReference, reference)
				}
			}
			o.References = finalReference

			// We need to remove not supported object from panelsJSON
			panels := new([]map[string]any)
			if err = json.Unmarshal([]byte(o.Attributes["panelsJSON"].(string)), panels); err != nil {
				return nil, err
			}
			finalPanels := make([]map[string]any, 0, len(*panels))
			for _, panel := range *panels {
				if panel["type"] == "lens" || panel["type"] == "map" {
					log.Warnf("We remove panel for object '%s' because the type '%s' is not supported on Opensearch", panel["panelIndex"], panel["type"])
				} else {
					finalPanels = append(finalPanels, panel)
				}
			}
			b, err := json.Marshal(finalPanels)
			if err != nil {
				return nil, err
			}
			o.Attributes["panelsJSON"] = string(b)

		}
	default:
		return nil, errors.Errorf("The object '%s' is on unsupported type '%s'", o.Id, o.Type)
	}

	dataConverted, err = json.Marshal(o)
	if err != nil {
		return nil, errors.Wrap(err, "Error when Marshall exported object")
	}

	log.Infof("Successfully migrate object contend '%s' of type '%s'", o.Id, o.Type)

	return dataConverted, err
}
