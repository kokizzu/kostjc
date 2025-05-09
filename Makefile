
setup:
	go get -u -v github.com/kokizzu/gotro@latest
	go install github.com/air-verse/air@latest
	go install github.com/fatih/gomodifytags@latest
	go install github.com/kokizzu/replacer@latest
	go install github.com/akbarfa49/farify@latest
	go install golang.org/x/tools/cmd/goimports@latest

local-tarantool:
	docker exec -it kostjc-tarantool tarantoolctl connect userT:passT@127.0.0.1:3301
	# box.space -- list all tables
	# box.execute [[ SELECT * FROM "users" LIMIT 1 ]]
	# \set language sql
	# \set delimiter ;

local-clickhouse:
	docker exec -it kostjc-clickhouse clickhouse-client -u userC
	# SHOW TABLES -- list all tables
	# SELECT * FROM "actionLogs" LIMIT 1;

download-backup:
	@echo "Downloading backup database from Server"
	rsync -Wav -e 'ssh -p 22 -i ~/.ssh/$(ssh-key)' $(user)@194.233.65.174:../kostjc/backup/* ./backup

upload-backup:
	@echo "Uploading backup database at $(timestamp) to Server"
	rsync -avz --progress ./backup/*$(timestamp)*.jsonline.lz4 -e 'ssh -p 22 -i ~/.ssh/$(ssh-key)' $(user)@194.233.65.174:../kostjc/backup/

modtidy:
	sudo chmod -R a+rwx _tmpdb && go mod tidy

fixtags:
	# fix struct tags
	go generate ./domain

orm:
	# generate ORM
	./gen-orm.sh

views:
	# generate views and routes
	./gen-views.sh

migrate:
	go run main.go migrate

backup-db:
	go run main.go backup

restore-db:
	go run main.go restore $(word 2, $(MAKECMDGOALS))

svelte:
	pnpm i
	pnpm watch

web:
	air web

%:
	@:
