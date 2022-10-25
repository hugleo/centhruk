#!/bin/sh

/usr/bin/go build
/usr/bin/cp centhruk /usr/local/bin
/usr/bin/cp centhruk.service /etc/systemd/system/
/usr/bin/cp -n centhruk.conf /etc/centhruk.conf
