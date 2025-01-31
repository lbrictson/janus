# Runtime Config

Janus is designed to be simple to get started with so there are no config files and the list of configuration
options are limited.  All configuration items are set by environment variables.  Settings these can vary based on your
environment.

| Variable              | Explanation                                              | Default Value         |
|-----------------------|----------------------------------------------------------|-----------------------|
| JANUS_PORT            | Port Janus runs on                                       | 8080                  |
| JANUS_DB_TYPE         | Database type (sqlite or postgres)                       | sqlite                |
| JANUS_DB_PATH         | Path to sqlite DB file, only used when DB_TYPE is sqlite | janus.db              |
| JANUS_DB_HOSTNAME     | Postgres database hostname                               | localhost             |
| JANUS_DB_PORT         | Postgres database port                                   | 5432                  |
| JANUS_DB_NAME         | Postgres database name                                   | postgres              |
| JANUS_DB_USER         | Postgres database username                               | postgres              |
| JANUS_DB_PASSWORD     | Postgres database password                               | postgres              |
| JANUS_DB_SSL_MODE     | Postgres SSL connection mode                             | disable               |
| JANUS_URL             | URL to access janus on, should be set to your domain     | http://localhost:8080 |
| JANUS_DISABLE_METRICS | Disable prometheus metric listener                       | false                 |
| JANUS_METRICS_PORT    | Port for prometheus metrics                              | 8081                  |
| JANUS_SESSION_NAME    | Name of the cookie for authentication tracking           | janus                 |
| JANUS_BRAND_NAME      | Controls the site name and brand name                    | Janus                 |