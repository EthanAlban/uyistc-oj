#!/bin/sh

while ! nc -vzw 2 47.96.99.172 9000; do
    echo "Waiting for upcoming beego Server"
    sleep 2
done
#java -jar $GOPATH/src/unioj-judger/run/oj-0.0.1-SNAPSHOT.jar
$GOPATH/src/unioj-judger/run/run
