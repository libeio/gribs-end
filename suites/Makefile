
include $(APPATH)/mk/common.mk

.PHONY: all

all: build install

clean:
	$(MAKE) -C photo clean

build:
	$(MAKE) -C photo build

install:
	$(MAKE) -C photo install

photo:
	$(MAKE) -C photo all
