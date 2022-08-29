all: build-server build-client


build-server:
	cd pow && go build -o ../build/server ./cmd/server && cd ..

build-client:
	cd pow && go build -o ../build/client ./cmd/client && cd ..
