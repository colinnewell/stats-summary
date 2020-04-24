all: stats-summary

stats-summary: main.go go.mod go.sum
	go build -o stats-summary main.go

clean:
	rm stats-summary

