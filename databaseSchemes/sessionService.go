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

func GetBySessionId(sessionId string) *structs.ChannelHost {
	transaction := dbConn.Txn(false)
	defer transaction.Abort()

	sessionInfo, err := transaction.First("channelHost", "sessionId", sessionId)
	if err != nil {
		panic(err)
	}
	if sessionInfo != nil {
		return sessionInfo.(*structs.ChannelHost)
	} else {
		return nil
	}
}

func DeleteBySessionId(sessionId string) {
	sessionInfo := GetBySessionId(sessionId)

	if sessionInfo != nil {
		Delete(sessionInfo)
	} else {
		fmt.Printf("Could not delete session. Session-ID <%s> not found.", sessionId)
	}
}

func UpdateSession(sessionId string, guestUri string) {
	session := GetBySessionId(sessionId)
	if session != nil {
		transaction := dbConn.Txn(true)
		session.Guest = guestUri
		if err := transaction.Insert("channelHost", session); err != nil {
			panic(err)
		}
		transaction.Commit()
	}
}

func GetAllEntries() {
	transaction := dbConn.Txn(false)
	defer transaction.Abort()

	it, err := transaction.Get("channelHost", "id")
	if err != nil {
		panic(err)
	}

	fmt.Println("All the people:")
	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(*structs.ChannelHost)
		fmt.Printf("  %s\n", p)
	}
}

func Count() int {
	transaction := dbConn.Txn(false)
	defer transaction.Abort()

	it, err := transaction.Get("channelHost", "id")
	if err != nil {
		panic(err)
	}
	var count = 0
	for obj := it.Next(); obj != nil; obj = it.Next() {
		count++
	}
	return count
}
