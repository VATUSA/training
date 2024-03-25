package env

type EnvironmentVariables struct {
	Database DatabaseEnvironment
	Test     string
}

type DatabaseEnvironment struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}
