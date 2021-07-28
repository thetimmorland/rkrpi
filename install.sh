#!/usr/bin/env bash
set -euo pipefail

WHEEL=rkrpi-0.1.0-py3-none-any.whl

poetry build
scp ./dist/$WHEEL root@rkrpi.local:/tmp/

ssh -T root@rkrpi.local <<EOF
    pip3 install -I /tmp/$WHEEL

    echo Restarting rkrpi.nmeaclient.service...
    systemctl restart rkrpi.nmeaclient.service

    echo Restarting rkrpi.httpserver.service...
    systemctl restart rkrpi.httpserver.service

    echo Done!
EOF
