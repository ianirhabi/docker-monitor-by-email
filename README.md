add your docker container service list to variable list_of_containers on dockermonitorbash
and add some of your email information on config.toml 

How to install this app:

- download and install golang sdk https://dl.google.com/go/go1.14.linux-amd64.tar.gz 
- $ go mod init docker-monitor-by-email 
- $ make compile
- $ sudo make install

Check the services
- $ sudo systemctl status dockermonitorgo
- $ sudo systemctl status dockermonitorbash
- $ sudo systemctl enable dockermonitorgo
- $ sudo systemctl enable dockermonitorbash

for uninstall is just simple :
- $ sudo make uninstall

for see logs :
log monitor-go just type:
- $ journalctl -f -u dockermonitorgo.service
log monitor-bash just type:
- $ journalctl -f -u dockermonitorbash.service