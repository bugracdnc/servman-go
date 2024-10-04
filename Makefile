clear:
	rm servman
	sudo rm /usr/bin/servman

servman:
	go build servman
	sudo cp servman /usr/bin