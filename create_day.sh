#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

read -p "Create a dir for which day?: " DAY

SCRIPT_DIR="$(pwd)"
if [ ! -d "$SCRIPT_DIR"/"$DAY" ]; then
  mkdir -p "$SCRIPT_DIR"/day"$DAY";
  (
    cd "$SCRIPT_DIR"/day"$DAY";
      mkdir -p input;
      touch main.go;
  )
fi