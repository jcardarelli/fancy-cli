#!/usr/bin/env bash
# Installation script for fancy-cli

SQLITE_DB=fancy-cli.db

if [[ ! -f $SQLITE_DB ]]; then
  sqlite3 $SQLITE_DB <sql/create-table.sql
fi
