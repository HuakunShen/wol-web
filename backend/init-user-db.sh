#!/bin/bash
set -e

[ ! -z "$WOL_WEB_USERNAME" ] && echo "default username exists" || echo "default username doesn't exist"
[ ! -z "$WOL_WEB_PASSWORD" ] && echo "default password exists" || echo "default password doesn't exist"

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    INSERT INTO users(username, password) VALUES ('$WOL_WEB_USERNAME', '$WOL_WEB_PASSWORD');
EOSQL