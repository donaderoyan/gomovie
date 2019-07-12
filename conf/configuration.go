package configuration

type Configuration struct {
  DB *DBConfiguration
}

type DBConfiguration struct {
  Dialect   string
  Host      string `required:"true"`
  Port      int
  User      string
  Dbname    string
  Password  string
}

func GetConfig() *Configuration {
  return &Configuration{
    DB: &DBConfiguration{
      Dialect:    "postgres",
      Host:       "127.0.0.1",
      Port:       5432,
      User:       "postgres",
      Dbname:     "gomovie",
      Password:   "ab123456"
    }
  }
}
