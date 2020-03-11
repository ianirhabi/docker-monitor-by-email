#!/bin/bash

#CREATED BY ANDRIAN LATIF on 02-18-2020

logic(){
    systemctl stop dockermonitorgo
    systemctl stop dockermonitorbash
    rm -rf /etc/systemd/system/dockermonitorgo.service
    rm -rf /etc/systemd/system/dockermonitorbash.service
    rm -rf /usr/bin/dockermonitorbash
    rm -rf /usr/bin/dockermonitorgo
    rm -rf /etc/dockermonitor

    echo "successfuly uninstall "

}

if [[ "$EUID" -ne 0 ]]; then
    echo "Sorry, you need to run this app as root"
    exit
else
    logic    
fi
