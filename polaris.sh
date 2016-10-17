#!/bin/bash

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
