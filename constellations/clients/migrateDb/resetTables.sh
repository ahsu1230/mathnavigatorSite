#!/bin/bash
#
# This script will "reset" the `mathnavdb` database of a specified mysql instance.
# The database reset will drop all tables and build them up again by following the migration files.
#
# Example Usage:
# ./resetTables.sh user password localhost 3308
# will reset the `mathnavdb` database at mysql://user:password@(localhost:3308)/mathnavdb

USERNAME=$1
PASSWORD=$2
HOST=$3
PORT=$4

MIGRATIONS_PATH="file://../../orion/src/repos/migrations"
MIGRATE_MYSQL_PATH="mysql://$USERNAME:$PASSWORD@\($HOST:$PORT\)/mathnavdb"

echo "Migrations Path: $MIGRATIONS_PATH"
echo "Target Migrate MySQL Path: $MIGRATE_MYSQL_PATH"

eval "migrate -source $MIGRATIONS_PATH -database $MIGRATE_MYSQL_PATH down"
eval "migrate -source $MIGRATIONS_PATH -database $MIGRATE_MYSQL_PATH up"
