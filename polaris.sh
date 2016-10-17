#!/bin/bash

### BEGIN INIT INFO
# Provides:          polaris
# Required-Start:    $local_fs $network
# Required-Stop:     $local_fs
# Default-Start:     2 3 4 5
# Default-Stop:      0 1 6
# Short-Description: polaris
# Description:       polaris webapp
### END INIT INFO

GOPATH=

case $1 in
    start)
        echo "Starting polaris service"
        cd $GOPATH/src/github.com/droff/polaris
        nohup ./polaris > /dev/null 2>&1 & echo $! > /tmp/polaris.pid
        ;;
    stop)
        echo "Stopping polaris service"
        kill "$(cat /tmp/polaris.pid)"
        ;;
    *)
        echo "polaris service."
        echo $"Usage $0 {start|stop}"
        exit 1
esac
exit 0
