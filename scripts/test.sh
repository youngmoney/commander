#!/usr/bin/env bash

diff <(go run . --config examples/config.yaml command simple) <(echo hi)
diff <(go run . --config examples/config.yaml command other) <(echo other; echo 3)
