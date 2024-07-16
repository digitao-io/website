#!/usr/bin/env bash

mysql --protocol=tcp -uroot -pno-password website-test < ./database/create_tables.sql
