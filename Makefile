GOCMD = go
BUILD = $(GOCMD) build
CLEAN = $(GOCMD) clean
TEST = $(GOCMD) test
#GET = $(GOCMD) get

all: test build
.PHONY: clean

build:
	$(BUILD) -v

test:
	$(TEST) -v

clean:
	$(CLEAN) -v
	rm -rfv bin
