package config


const (
	GLOBAL_STATE_TABLE = "CREATE TABLE IF NOT EXISTS global_state (" +
		"id TEXT PRIMARY KEY,"+
		"gstate TEXT CHECK( gstate IN ('"+STATE_FULL+"','"+STATE_PART+"','"+STATE_INAC+"') ) NOT NULL DEFAULT '');"
)

