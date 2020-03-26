/**
*@program: mongo-info
*@description: https://github.com/dollarkillerx
*@author: dollarkiller [dollarkiller@dollarkiller.com]
*@create: 2020-03-25 15:55
 */
package examples

import (
	"context"
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


	type ac struct {
		A string
	}
	type user struct {
		User string
		Password string
		A *ac
	}



	collection := db.Collection("colection1")

	insert, err := collection.Insert(&user{User: "abcd", Password: "21312qwewcfafsdf",A:&ac{A:"aaaaaaaaaaaa"}})
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

func TestUpdate(t *testing.T) {
	db,err := New("mongodb://localhost:27017", "test")
	if err != nil {
		log.Fatalln(err)
	}

	type user struct {
		User string
		Password string
	}

	collection := db.Collection("colection1")
	_, err = collection.Update(bson.M{"user": "abcd"}, bson.D{{
		"$set",
		bson.D{
			{"password","123456456"},
		},
	}})
	if err != nil {
		log.Fatalln("S1 ",err)
	}
}

func TestFindAll(t *testing.T) {
	db,err := New("mongodb://localhost:27017", "test")
	if err != nil {
		log.Fatalln(err)
	}

	type user struct {
		User string
		Password string
	}

	collection := db.Collection("colection1")
	all, err := collection.FindAll(bson.M{"user": "abcd"})
	if err != nil {
		log.Fatalln("S1: ",err)
	}

	nodes := make([]*user,0)
	for all.Next(context.Background()) {
		var node user
		if err = all.Decode(&node); err != nil {
			log.Println("S2: ",err)
			continue
		} else {
			fmt.Println("ID: ",all.ID())
			nodes = append(nodes, &node)
		}
	}

	log.Println(nodes)

}


func TestFindAll2(t *testing.T) {
	db,err := New("mongodb://localhost:27017", "test")
	if err != nil {
		log.Fatalln(err)
	}

	type user struct {
		User string
		Password string
	}

	collection := db.Collection("colection1")
	all, err := collection.FindAll(bson.M{"user": "abcd"})
	if err != nil {
		log.Fatalln("S1: ",err)
	}

	nodes := make([]user,0)
	for all.Next(context.Background()) {
		fmt.Println("ID ",all.ID())
		var node user
		if err = all.Decode(&node); err != nil {
			log.Println("S2: ",err)
			continue
		} else {
			nodes = append(nodes, node)
		}
	}

	log.Println(nodes)

}

func TestFindAll3(t *testing.T) {
	db,err := New("mongodb://localhost:27017", "test")
	if err != nil {
		log.Fatalln(err)
	}

	collection := db.Collection("colection12")
	_, err = collection.FindAll(bson.M{"user": "abcd"})
	if err != nil {
		log.Fatalln("S1: ",err)
	}
	log.Println("OK")
}
