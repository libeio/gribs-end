
export APPATH=$(PWD)

include $(APPATH)/mk/common.mk

.PHONY: all

all: deps build install

deps:
	@if [ ! -e go.mod ]; then 			\
		$(GO-MOD) init $(MODULENAME);	\
	fi

	@$(GO-GET) github.com/go-pg/pg/v10@v10.9.0
	@$(GO-GET) github.com/spf13/viper@v1.9.0

clean:
	$(MAKE) -C suites clean
	$(MAKE) -C bin    clean

build:
	$(MAKE) -C suites build

install:
	$(MAKE) -C suites install

photo:
	$(MAKE) -C suites photo
