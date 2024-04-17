test:
	go test ./...

bench:
	go test -benchmem -run=^$$ -bench . github.com/Bot-Hive-Trading/gws

cover:
	go test -coverprofile=./bin/cover.out --cover ./...
