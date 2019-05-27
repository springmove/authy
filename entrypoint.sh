#!/bin/sh
set -e

if [ ! -f "/etc/authy/conf/config.yml" ];then
    cp /etc/authy/config.yml /etc/authy/conf/config.yml
fi

exec "$@"
