#!/bin/bash

# this code created by Andrian Latif 02-18-2020
# compile your code before using this code
# happy code 

__CURRENT_DIRECTORY=$(pwd)

logic(){
    rm -rf /etc/systemd/system/dockermonitorgo.service
    rm -rf /etc/systemd/system/dockermonitorbash.service
    # for golang app
    echo "[Unit]
Description= docker-monitor-golang
After=dockermonitorgo.service
Requires=dockermonitorgo.service

[Service]
User=root
ExecStart=/usr/bin/dockermonitorgo

[Install]
WantedBy=default.target 
    " > /etc/systemd/system/dockermonitorgo.service

    # bash app
    echo "[Unit]
Description= docker-monitor-bash
After=dockermonitorbash.service
Requires=dockermonitorbash.service

[Service]
User=root
ExecStart=/usr/bin/dockermonitorbash

[Install]
WantedBy=default.target 
    " > /etc/systemd/system/dockermonitorbash.service

    shell_command
}

shell_command(){
    rm -rf /usr/bin/dockermonitorbash
    rm -rf /usr/bin/dockermonitorgo
    rm -rf /etc/dockermonitor
    mkdir /etc/dockermonitor
    chmod 777 /etc/dockermonitor
    cp $__CURRENT_DIRECTORY/dockermonitorbash $__CURRENT_DIRECTORY/app
    cp $__CURRENT_DIRECTORY/app/dockermonitorgo /usr/bin/
    cp $__CURRENT_DIRECTORY/app/dockermonitorbash /usr/bin/
    cp $__CURRENT_DIRECTORY/config.toml /etc/dockermonitor/
    systemctl daemon-reload
    systemctl start dockermonitorgo
    systemctl daemon-reload
    systemctl start dockermonitorbash
    echo "successfuly install"
    echo "for see the log service of dockermonitorgo just type journalctl -f -u dockermonitorgo.service"
    echo "for see the log service of dockermonitorbash just type journalctl -f -u dockermonitorbash.service"
}

if [[ "$EUID" -ne 0 ]]; then
    echo "Sorry, you need to run this app as root"
    exit
else
    logic    
fi