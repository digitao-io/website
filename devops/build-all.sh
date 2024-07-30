#!/usr/bin/env bash

VERSION_NUMBER=20240730

rm -rf digitao-website-images-*.tar.gz

docker buildx build \
  -f backend.Dockerfile \
  -t digitao-website-backend:${VERSION_NUMBER} \
  ..

docker buildx build \
  -f presentation.Dockerfile \
  -t digitao-website-presentation:${VERSION_NUMBER} \
  ..

cat > compose.override.yml <<- EOM
services:
  backend:
    image: digitao-website-backend:${VERSION_NUMBER}
  presentation:
    image: digitao-website-presentation:${VERSION_NUMBER}
EOM

docker save --output digitao-website-images-${VERSION_NUMBER}.tar \
  digitao-website-backend:${VERSION_NUMBER} \
  digitao-website-presentation:${VERSION_NUMBER}

gzip digitao-website-images-${VERSION_NUMBER}.tar
