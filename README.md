Magic the Gathering card archive (mtg-archive)
================================================

To Start
--------
First thing is to create local development database **mtg-local**. After local database
is created if can be migrated to latest version using **goose -env local up**.

Tools
-----

* **Goose** - Database migrations (`go get bitbucket.org/liamstask/goose/cmd/goose`)
* **postgres** - Database engine
* **gorilla/mux** - Request router (`go get github.com/gorilla/mux`)
* **go-logging** - Logger (`go get github.com/op/go-logging`)
* **gorm** - ORM package to use postgres from go (`go get github.com/jinzhu/gorm`)