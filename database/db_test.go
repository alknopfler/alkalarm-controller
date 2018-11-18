package database

import(
	_ "github.com/mattn/go-sqlite3"

	"testing"
	"github.com/stretchr/testify/assert"
)

func TestInitDB(t *testing.T) {
	db,err := InitDB("./test.db")
	defer db.Close()
	assert.NoError(t,err)
}

func TestCreateSchemas(t *testing.T) {
	db,_ := InitDB("./test.db")
	defer db.Close()
	err:=CreateSchemas(db)
	assert.NoError(t,err)
}


