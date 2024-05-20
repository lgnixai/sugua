#!/bin/bash

# Usage: ./scripts/start.sh

set -e

cd "$(dirname "$0")/../"

air -c ./scripts/.xnote.toml
