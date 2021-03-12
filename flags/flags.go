package flags

import (
	"flag"
)

// Config Store the data
type Config struct {
	Host     string
	User     string
	Pass     string
	Database string
	Table    string
	Field    string
	Value    string
	Port     int64
	Output   bool
}

// Init gather input
func Init() *Config {

	const (
		defaultHost = "localhost"
		usageHost   = "Database hostname"

		defaultUser = "root"
		usageUser   = "Database User"

		defaultPass = ""
		usagePass   = "Database User"

		defaultDatabase = "test-database"
		usageDatabase   = "Database name"

		defaultTable = "user"
		defaultField = "id"
		defaultValue = "1"
		usageTable   = "Database name"

		defaultPort = 3306
	)

	config := &Config{}

	flag.StringVar(&config.Host, "host", defaultHost, usageHost)
	flag.StringVar(&config.Host, "h", defaultHost, usageHost+" (shorthand)")

	flag.StringVar(&config.User, "user", defaultUser, usageUser)
	flag.StringVar(&config.User, "u", defaultUser, usageUser+" (shorthand)")

	flag.StringVar(&config.Pass, "passwd", defaultPass, usagePass)
	flag.StringVar(&config.Pass, "p", defaultPass, usagePass+" (shorthand)")

	flag.StringVar(&config.Database, "database", defaultDatabase, usageDatabase)
	flag.StringVar(&config.Database, "d", defaultDatabase, usageDatabase+" (shorthand)")

	flag.StringVar(&config.Table, "table", defaultTable, usageDatabase)
	flag.StringVar(&config.Table, "t", defaultTable, usageDatabase+" (shorthand)")

	flag.StringVar(&config.Field, "field", defaultField, usageDatabase)
	flag.StringVar(&config.Field, "f", defaultField, usageDatabase+" (shorthand)")

	flag.StringVar(&config.Value, "value", defaultValue, usageDatabase)
	flag.StringVar(&config.Value, "v", defaultValue, usageDatabase+" (shorthand)")

	flag.BoolVar(&config.Output, "output", true, usageDatabase)
	flag.BoolVar(&config.Output, "o", true, usageDatabase+" (shorthand)")

	flag.Int64Var(&config.Port, "port", defaultPort, usageUser)
	//flag.Int64Var(&config.dbPort, "p", defaultPort, usageUser+" (shorthand)")

	flag.Parse()

	return config
}
