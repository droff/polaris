#!/bin/bash

case $1 in
    start)
        echo "Starting polaris service"
        nohup $GOPATH/bin/polaris > /dev/null 2>&1 & echo $! > /tmp/ploaris.pid
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
