#!/usr/bin/env bash

# This is a script to generate the api client for the Meta conversions api.
# https://developers.facebook.com/docs/marketing-api/conversions-api/using-the-api#swagger
# https://github.com/facebookincubator/Facebook-Server-Side-API-Swagger

script_folder="$(realpath "$(dirname "$0")")"

user="$(id -u):$(id -g)"
commands="
cd /opt/swagger-codegen-cli
apk add --no-cache git
java -jar swagger-codegen-cli.jar generate -i scripts/server-side-api.yaml -l go -o /api
chown $user -R /api
"

outfolder="facebookgen"

docker run --rm --entrypoint="/bin/sh" \
  -v "$script_folder:/opt/swagger-codegen-cli/scripts" \
  -v "$script_folder/../pkg/$outfolder:/api" swaggerapi/swagger-codegen-cli \
  -c "$commands"

# Change package name.
for file in ../pkg/"$outfolder"/*.go
do
  sed -i 's/package swagger/package facebookgen/' "$file"
done
