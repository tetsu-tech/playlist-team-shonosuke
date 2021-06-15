#!/bin/sh

function sqldef_dry() {
  ./mysqldef -h localhost -u root kotopa --dry-run < schema.sql
}

function sqldef_apply() {
  ./mysqldef -h localhost -u root kotopa  < schema.sql
}

function sqlboiler() {
  ./sqlboiler mysql --no-tests --no-hooks --wipe
}

$1
