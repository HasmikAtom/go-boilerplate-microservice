package database

//** Structs for Mutable's config.json **//

// Mutable Configs
type config struct {
	Prod envConfig
	Dev  envConfig
}

//** Example Environment Configs //
type envConfig struct {
	// Postgres postgresConfig
}
