#!/bin/bash

#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE USER gwp WITH PASSWORD 'gwp';
    GRANT ALL PRIVILEGES ON DATABASE gwp TO gwp;
    GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO gwp;

    GRANT ALL PRIVILEGES ON TABLE posts TO gwp;
    GRANT ALL PRIVILEGES ON TABLE comments TO gwp;

EOSQL
