### Install
```sh
go get github.com/rizvn/registry
```

### Example Usage of the Library
Create a registry instance and initialize it:
```go
register := registry.Registry{}
register.Init()
```

Define a Registerable instance
```go
type Db struct {
	dbx      *sqlx.DB
	gorm     *gorm.DB
	DbHost   string
	DbPort   string
	DbUser   string
	DbPass   string
	DBName   string
	DbParams string
}

// Register implements the registry.Registerable interface
func (r *Db) Register(reg *registry.Registry) {
	r.DbUser = util.GetRequiredEnvVar("DB_USER")
	r.DbPass = util.GetRequiredEnvVar("DB_PASS")
	r.DbHost = util.GetRequiredEnvVar("DB_HOST")
	r.DbPort = util.GetRequiredEnvVar("DB_PORT")
	r.DBName = util.GetRequiredEnvVar("DB_NAME")
	r.DbParams = os.Getenv("DB_PARAMS")

	// .. additional initialization logic..
	
	// Set the instance in the registry
	reg.Set(r)
}
```


Add a registerable instance
```go
register.Add(&db.Db{})
```

Build Register
```go
register.Build()
```


Access the registered instance
```go
dbInstance := register.Get(&db.Db{}).(*db.Db)
```