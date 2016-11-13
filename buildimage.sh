#!/bin/bash
rm catseye-linux-amd64.aci
acbuild begin
acbuild set-name catseye
chmod +x scripts/*
acbuild copy scripts/repopulate.sh /bin/repopulate
acbuild copy scripts/run.sh /bin/run
acbuild copy snapOne.json /bin/snapOne.json
acbuild label add os linux
acbuild label add arch amd64
acbuild set-exec /bin/run
acbuild write catseye-linux-amd64.aci
acbuild end
