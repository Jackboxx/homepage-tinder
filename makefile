build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm ./src/main
	go build -o homepage-tinder ./src/main

run: build
	echo "http://localhost:5050"
	./homepage-tinder
