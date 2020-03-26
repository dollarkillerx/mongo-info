/**
*@program: mongo-info
*@description: https://github.com/dollarkillerx
*@author: dollarkiller [dollarkiller@dollarkiller.com]
*@create: 2020-03-25 15:55
 */
package examples

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"testing"
)

func TestInsert(t *testing.T) {
	db,err := New("mongodb://localhost:27017", "test")
	if err != nil {
		log.Fatalln(err)
	}

	type user struct {
		User string
		Password string
	}

	collection := db.Collection("colection1")

	insert, err := collection.Insert(&user{User: "abcd", Password: "21312qwewcfafsdf"})
	if err != nil {
		log.Fatalln("ss1 ",err)
	}
	log.Println("Id: ",insert.InsertedID)

	result,err := collection.Select(bson.M{"user": "abcd"})
	if err != nil {
		log.Fatalln("s1 ",err)
	}
	u := &user{}
	err = result.Decode(u)
	if err != nil {
		log.Fatalln("s2 ",err)
	}
	fmt.Println(u)
}
