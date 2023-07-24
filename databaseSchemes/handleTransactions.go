package databaseSchemes

import (
	"fmt"
	"github.com/hashicorp/go-memdb"
	"websocketServer/structs"
)

var dbConn *memdb.MemDB

func SetDbConn(db *memdb.MemDB) {
	dbConn = db
}

func Write(userInfo *structs.ChannelHost) {
	transaction := dbConn.Txn(true)
	if err := transaction.Insert("channelHost", userInfo); err != nil {
		panic(err)
	}
	transaction.Commit()
}

func Delete(sessionInfo *structs.ChannelHost) {
	transaction := dbConn.Txn(true)
	if err := transaction.Delete("channelHost", sessionInfo); err != nil {
		fmt.Printf("Error while terminating session with ID '%s'. Error: %s", sessionInfo.SessionID, err)
	}
	transaction.Commit()
}

func DeleteBySessionId(sessionId string) {
	transaction := dbConn.Txn(false)
	defer transaction.Abort()

	sessionInfo, err := transaction.First("channelHost", "sessionId", sessionId)
	if err != nil {
		panic(err)
	}
	if sessionInfo != nil {
		Delete(sessionInfo.(*structs.ChannelHost))
	} else {
		fmt.Printf("Could not delete session. Session-ID %s not found.", sessionInfo)
	}

}
