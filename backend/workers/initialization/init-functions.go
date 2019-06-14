package initialization

import (
	"github.com/jaxmef/booklover-backend/config"
	"github.com/lancer-kit/armory/db"
	"github.com/sirupsen/logrus"
)

type initModule string

var (
	DB   initModule = "database connection"
)

func initDatabase(cfg *config.Cfg, entry *logrus.Entry) error {
	return db.Init(cfg.DB, entry)
}
