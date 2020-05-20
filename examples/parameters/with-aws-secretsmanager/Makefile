.PHONY: $(MAKECMDGOALS)

run:
	go run main.go --STAGE=DEV

run-flags:
	go run main.go --STAGE=DEV --HOST="some host" --PORT=1234 --DATABASE=database_name --LOCAL=true

run-env:
	STAGE=DEV HOST="some host" PORT=1234 DATABASE="database_name" LOCAL=true go run main.go