
export PKG_CONFIG_PATH=$(shell pwd)

install:
	go install

test:
	echo "PKG_CONFIG_PATH=${PKG_CONFIG_PATH}"
	go test -v
