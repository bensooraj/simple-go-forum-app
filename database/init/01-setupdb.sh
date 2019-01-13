#!/bin/bash

#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE USER gwp;
    GRANT ALL PRIVILEGES ON DATABASE gwp TO gwp;
EOSQL
