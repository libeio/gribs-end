
include $(APPATH)/mk/common.mk
include $(APPATH)/mk/photo.mk

.PHONY: all

all: build install

clean:
	@for exefile in $(EXEFILES); do $(RM) $$exefile ; done;

build: $(GOFILES)
	$(GO-BUILD) -o $(SERVICENAME) $^

install:
	$(MV) $(SERVICENAME) $(APPATH)/bin
