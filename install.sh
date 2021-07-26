#!/usr/bin/env bash
set -euo pipefail

GOOS=linux GOARCH=arm GOARM=5 CGO_ENABLED=1 \
    go build -o . ./...

scp rkrpi_nmeaclient root@raspberrypi.local:/usr/local/bin
scp rkrpi_httpserver root@raspberrypi.local:/usr/local/bin

ssh root@raspberrypi.local "bash -s" <<EOF
    systemctl restart rkrpi_nmeaclient.service
    systemctl restart rkrpi_httpserver.service
EOF
