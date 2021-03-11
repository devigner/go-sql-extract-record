package flags

import (
	"flag"
)

type Config struct {
	DbHost     string
	DbUser     string
	DbPass     string
	DbDatabase string
	Table      string
	Field      string
	Value      string
	DbPort     int64
	Output     bool
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

	flag.StringVar(&config.DbHost, "host", defaultHost, usageHost)
	flag.StringVar(&config.DbHost, "h", defaultHost, usageHost+" (shorthand)")

	flag.StringVar(&config.DbUser, "user", defaultUser, usageUser)
	flag.StringVar(&config.DbUser, "u", defaultUser, usageUser+" (shorthand)")

	flag.StringVar(&config.DbPass, "passwd", defaultPass, usagePass)
	flag.StringVar(&config.DbPass, "p", defaultPass, usagePass+" (shorthand)")

	flag.StringVar(&config.DbDatabase, "database", defaultDatabase, usageDatabase)
	flag.StringVar(&config.DbDatabase, "d", defaultDatabase, usageDatabase+" (shorthand)")

	flag.StringVar(&config.Table, "table", defaultTable, usageDatabase)
	flag.StringVar(&config.Table, "t", defaultTable, usageDatabase+" (shorthand)")

	flag.StringVar(&config.Field, "field", defaultField, usageDatabase)
	flag.StringVar(&config.Field, "f", defaultField, usageDatabase+" (shorthand)")

	flag.StringVar(&config.Value, "value", defaultValue, usageDatabase)
	flag.StringVar(&config.Value, "v", defaultValue, usageDatabase+" (shorthand)")

	flag.BoolVar(&config.Output, "output", true, usageDatabase)
	flag.BoolVar(&config.Output, "o", true, usageDatabase+" (shorthand)")

	flag.Int64Var(&config.DbPort, "port", defaultPort, usageUser)
	//flag.Int64Var(&config.dbPort, "p", defaultPort, usageUser+" (shorthand)")

	flag.Parse()

	return config
}
