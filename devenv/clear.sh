#!/usr/bin/env bash

mongosh \
  "mongodb://localhost:27017/website_test" \
  --username root \
  --password no-password \
  --authenticationDatabase admin \
  --eval "db.getCollectionNames().forEach((name) => { db[name].deleteMany({}) })"

mc rb --force website-test/website-test
