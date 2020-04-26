all: stats-summary

stats-summary: main.go go.mod go.sum stats/stats.go
	go build -o stats-summary main.go

test:
	go test ./...

clean:
	rm stats-summary

