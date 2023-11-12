package boot

import (
	"database/sql"
)

type Application struct {
	Env *Env
	DB  *sql.DB
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.DB = NewPostgresDB(app.Env)
	return *app
}

func (a *Application) Close() {
	ClosePostgresConnection(a.DB)
}
