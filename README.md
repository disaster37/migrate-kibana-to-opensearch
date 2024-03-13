# migrate-kibana-to-opensearch

This project permit to migrate Kibana dashboard to Opensearch dashboard. It will include all objects references by Dashboard.

## Use it

## Global parameters

```
GLOBAL OPTIONS:
   --config FILE               Load configuration from FILE
   --kibana-url value          The kibana URL [$KIBANA_URL]
   --kibana-user value         The kibana user [$KIBANA_USER]
   --kibana-password value     The Kibana password [$KIBANA_PASSWORD]
   --dashboard-url value       The Opensearch dashboard URL [$DASHBOARD_URL]
   --dashboard-user value      The Opensearch dashboard user [$DASHBOARD_USER]
   --dashboard-password value  The Opensearch dashboard password [$DASHBOARD_PASSWORD]
   --self-signed-certificate   Disable the TLS certificate check (default: false)
   --debug                     Display debug output (default: false)
   --help, -h                  show help
   --version, -v               print the version
```

### Migrate dashboard

You can use the following parameters
```
OPTIONS:
   --dashboard-id value [ --dashboard-id value ]  The dashboard ids. If not provided, it migrate all dashboards
   --space value                                  The Kibana space where export dahsboards. If not provided is use default/global space/tenant
```

Sample to migrate only one dashboard
```bash
./migrate-kibana-to-opensearch --kibana-url https://kibana.domain.com --kibana-user elastic --kibana-password changme --dashboard-url https://dashboard.domain.com --dashboard-user admin --dashboard-password changeme migrate-dashboard --dashboard-id dashboard_id_1
```

Sample to migrate all dahsboards
```bash
./migrate-kibana-to-opensearch --kibana-url https://kibana.domain.com --kibana-user elastic --kibana-password changme --dashboard-url https://dashboard.domain.com --dashboard-user admin --dashboard-password changeme migrate-dashboard
```