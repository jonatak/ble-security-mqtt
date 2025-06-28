build:
	@go build -o ./bin/ble-monitor cmd/ble-monitor/main.go

clean:
	@rm -rf ./bin/*

run:
	go run cmd/ble-monitor/main.go 