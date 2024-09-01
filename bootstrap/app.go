package bootstrap

import "github.com/Sebastien-jo/SportTrackGo/mysql"

type Application struct {
  Env *Env
  Mysql mysql.Client
}

func App() Application {
  app := &Application{}
  app.Env = NewEnv()
  app.Mysql = NewMysqlDatabase(app.Env)
  return *app
}

func (app *Application) CloseDBConnection() {
  CloseMySqlDBConnection(app.Mysql)
} 
