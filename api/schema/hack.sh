#!/bin/sh

function sqldef_dry() {
  ./mysqldef -h 127.0.0.1 -u root playlist --dry-run < schema.sql
}

function sqldef_apply() {
  ./mysqldef -h 127.0.0.1 -u root playlist  < schema.sql
}

function sqlboiler() {
  ./sqlboiler mysql --no-tests --no-hooks --wipe
}

$1
