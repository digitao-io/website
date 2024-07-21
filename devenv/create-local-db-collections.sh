#!/usr/bin/env bash

mongosh "mongodb://localhost:27017/website_test" --username root --password no-password --authenticationDatabase admin ./database/create-collections.js
