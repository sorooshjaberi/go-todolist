package constants

type envKeysType struct {
	JWTSecret  string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPass     string
	DBDatabase string
}

var EnvKeys = envKeysType{
	JWTSecret:  "JWT_SECRET",
	DBHost:     "DB.HOST",
	DBPort:     "DB.PORT",
	DBDatabase: "DB.DATABASE",
	DBUser:     "DB.USER",
	DBPass:     "DB.PASSWORD",
}
