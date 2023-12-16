#!/bin/sh

sqlite3 db.db '.read init.sql'
echo 'built sqlite db'

echo $(sqlite3 db.db '.schema')
