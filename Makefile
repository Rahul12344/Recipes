EXEC=Recipes

.PHONY: test

run: bin/bconnect test
	./bin/$(EXEC)

bin/bconnect:
	go build -o bin/$(EXEC) *.go

clean:
	rm -rf bin/BConnect-backend

test:
	./test.sh
	go test -v ./...
