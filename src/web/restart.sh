#!/bin/sh
cd /opt/web
pgrep web | xargs kill -9
nohup ./web &