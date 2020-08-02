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

bin/recipes: clean
	go build -o bin/$(EXEC) cmd/recipes/main.go

clean:
	rm -rf bin/Recipes

test:
	./test.sh
	go test -v ./...
