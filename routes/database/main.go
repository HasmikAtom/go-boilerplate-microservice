package database

import meta "github.com/HasmikAtom/meta-go"

func getDatabaseConfig() config {
	var conf config
	if err := meta.UnmarshalConfig(&conf); err != nil {
		panic(err)
	}
	return conf
}

// Conf stores the database configuration
var Conf = getDatabaseConfig()
