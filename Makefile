# create by Andrian Latif on 02-18-2020

# make
# for see the log just type on terminal
# journalctl -f -u dockermonitorgo.service
# journalctl -f -u dockermonitorbash.service

install: 
	@./install.sh

# make compile
# this stage will compile the app from source code 
compile:
	@mkdir app
	@go mod download
	@go build -o app/dockermonitorgo main.go
	@chmod +x app/dockermonitorgo
	@cp dockermonitorbash app/ 


# for delete all of services
uninstall:
	@./uninstall.sh