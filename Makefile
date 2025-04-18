run_dev: build_dev
	./build/in_reserv

run_prod: build_prod
	./build/in_reserv

build_dev:
	go build -tags dev -o build/in_reserv .

build_prod:
	go build -o build/in_reserv .
