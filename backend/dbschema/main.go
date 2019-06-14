package dbschema

import (
	"github.com/lancer-kit/armory/db"
	"github.com/rubenv/sql-migrate"
)

//go get -u github.com/lancer-kit/forge

//go:generate go-bindata -ignore .+\.go$ -pkg dbschema -o bindata.go ./...
//go:generate gofmt -w bindata.go

func SetAssets() {
	db.SetAssets(migrate.AssetMigrationSource{
		Asset:    Asset,
		AssetDir: AssetDir,
		Dir:      "migrations",
	})
}
