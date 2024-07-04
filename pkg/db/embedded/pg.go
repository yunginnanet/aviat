package embedded

import (
	"errors"
	"io"

	"git.tcp.direct/kayos/common/entropy"
	"github.com/fergusstrange/embedded-postgres"
)

var (
	globalPostgres *embeddedpostgres.EmbeddedPostgres
	password       string
)

func StartEmbeddedPostgresql(logger io.Writer, dataPath, database string) (string, error) {
	password = entropy.RandStrWithUpper(16)

	opt := embeddedpostgres.DefaultConfig().
		DataPath(dataPath).Database(database).
		Username("aviat").Password(password).
		Logger(logger)

	globalPostgres = embeddedpostgres.NewDatabase()
	if err := globalPostgres.Start(); err != nil {
		return "", err
	}

	return opt.GetConnectionURL(), nil
}

func StopEmbeddedPostgresql() error {
	if globalPostgres == nil {
		return errors.New("embedded postgres not started")
	}
	return globalPostgres.Stop()
}
