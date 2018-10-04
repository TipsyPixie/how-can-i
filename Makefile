GOCMD = go
MAIN_PACKAGE = howtodo
INSTALL = $(GOCMD) install
CLEAN = $(GOCMD) clean
TEST = $(GOCMD) test
#GET = $(GOCMD) get

all: install test
.PHONY: install test clean

install:
	$(INSTALL) -v $(MAIN_PACKAGE) 

test:
	$(TEST) -v $(MAIN_PACKAGE)

clean:
	$(CLEAN) -v
	rm -rfv bin
