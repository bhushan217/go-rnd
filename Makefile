# .PHONY: clean, test, run

# clean:
#   rm -rf *.out

run:
	go run main.go

test:
	go test -coverprofile=coverage.out && go tool cover -html=coverage.out
