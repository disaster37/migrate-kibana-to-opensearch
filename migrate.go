package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/disaster37/migrate-kibana-to-opensearch/pkg/kibana"
	"github.com/disaster37/migrate-kibana-to-opensearch/pkg/migrate"
	"github.com/urfave/cli/v2"
)

func migrateDashboard(c *cli.Context) error {

	// Get kibana client
	kibanaClient, err := kibana.ManageKibanaGlobalParameters(c)
	if err != nil {
		return err
	}

	// Get dashboards
	exportedRawData, err := kibana.ExportDashboards(c.StringSlice("dashboard-id"), c.String("space"), kibanaClient)
	if err != nil {
		return err
	}

	// Split datas
	exportedDatas, err := migrate.SplitNewLineByteBuffer(exportedRawData)
	if err != nil {
		return err
	}

	// Migrate each objects
	convertedDatas := make([][]byte, 0, len(exportedDatas))
	for _, data := range exportedDatas {
		convertedData, err := migrate.ConvertObjectFromElasticsearchToOpensearch(data)
		if err != nil {
			return err
		}
		convertedDatas = append(convertedDatas, convertedData)
	}

	spew.Print(convertedDatas)

	return nil
}
