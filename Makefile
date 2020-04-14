EXEC=Recipes


.PHONY: test

db:
	psql \
	   --host=$(HOST) \
	   --port=$(PORT) \
	   --username=$(USERNAME) \
	   --password \
	   --dbname=$(DATABASE) \

run: db bin/bconnect test
	./bin/$(EXEC)

bin/bconnect:
	go build -o bin/$(EXEC) *.go

clean:
	rm -rf bin/BConnect-backend

test:
	./test.sh
	go test -v ./...
