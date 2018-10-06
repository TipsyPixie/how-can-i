GOCMD = go
MAIN_PACKAGE = howtodo
INSTALL = $(GOCMD) install
CLEAN = $(GOCMD) clean
TEST = $(GOCMD) test

.PHONY: all install test clean
all: install test

install:
	$(INSTALL) -v $(MAIN_PACKAGE) 

test:
	$(TEST) -v $(MAIN_PACKAGE)

clean:
	$(CLEAN) -v
	rm -rfv bin
