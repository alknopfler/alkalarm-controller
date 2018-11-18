package config

/*
		global:  insert | select | delete (all) | update
 */

const (
		GLOBAL_STATE_INSERT= "INSERT INTO global_state(id,gstate) values(1,?)"
		GLOBAL_STATE_UPDATE = "UPDATE global_state SET gstate=? WHERE id=1"
		GLOBAL_STATE_QUERY="SELECT * FROM global_state WHERE id=1"

)
