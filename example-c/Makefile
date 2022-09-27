
.PHONY: build clean

install: build
	@cmake --install _build 

configure:
	@cmake . -B _build

build: configure
	@cmake --build _build

test: build
	@ctest --test-dir _build --output-on-failure --progress

clean:
	rm -Rf _build
