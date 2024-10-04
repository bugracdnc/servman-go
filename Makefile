install: servman

reinstall: clean servman

servman:
	go build servman
	sudo cp servman /usr/bin

clean:
	rm -f servman
	sudo rm -f /usr/bin/servman