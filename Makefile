all: build-server build-client


build-server:
	cd src && go build -o ../build/server ./cmd/server && cd ..

build-client:
	cd src && go build -o ../build/client ./cmd/client && cd ..
