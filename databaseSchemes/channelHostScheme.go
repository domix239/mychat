package databaseSchemes

import "github.com/hashicorp/go-memdb"

var channelHostDb *memdb.MemDB

// InitDatabaseScheme asd
func InitDatabaseScheme() (dbConn *memdb.MemDB) {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"channelHost": &memdb.TableSchema{
				Name: "channelHost",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.IntFieldIndex{Field: "Id"},
					},
					"uri": &memdb.IndexSchema{
						Name:    "uri",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "URI"},
					},
					"sessionId": &memdb.IndexSchema{
						Name:    "sessionId",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "SessionId"},
					},
					"alive": &memdb.IndexSchema{
						Name:    "alive",
						Unique:  false,
						Indexer: &memdb.BoolFieldIndex{Field: "Alive"},
					},
				},
			},
		},
	}

	db, err := memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}
	return db
}
