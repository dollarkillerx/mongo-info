/**
*@program: mongo-info
*@description: https://github.com/dollarkillerx
*@author: dollarkiller [dollarkiller@dollarkiller.com]
*@create: 2020-03-25 15:55
 */
package examples

import (
	"context"
	"encoding/json"
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


// 解码一些奇怪的东西
type A struct {
	Name string
	Data []c
}

type c interface {
	_a()
}

type b1 struct {
	Name string
	Pc string
}
func (b *b1) _a() {}

type b2 struct {
	Name string
	Age int
}
func (b *b2) _a() {}

func TestInAndOut(t *testing.T) {
	a := A{
		Name:"你好",
		Data:[]c{
			&b1{
				Name:"b1",
				Pc:"P2",
			},
			&b2{
				Name:"N2",
				Age:18,
			},
		},
	}

	db,err := New("mongodb://localhost:27017", "test")
	if err != nil {
		log.Fatalln(err)
	}
	collection := db.Collection("colection1")
	insert, err := collection.Insert(&a)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(insert)

	result, err := collection.Select(bson.M{"name": "你好"})
	if err != nil {
		log.Fatalln(err)
	}
	b := &A{}
	err = result.Decode(b)
	if err != nil {
		//log.Fatalln("1", err)
	}
	log.Println(b)



	marshal, err := json.Marshal(a)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(marshal))

	bb := &A{}
	err = json.Unmarshal(marshal, bb)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(bb)
}


func TestInterface(t *testing.T) {
	type s1 struct {
		Name string
		Ad interface{}
	}
	type a struct {
		Age int
	}

	// 如果是interface 将无法解析
	db,err := New("mongodb://localhost:27017", "test")
	if err != nil {
		log.Fatalln(err)
	}
	collection := db.Collection("colection1")

	b := s1{
		Name:"1112",
		Ad:&a{
			Age:16,
		},
	}

	insert, err := collection.Insert(b)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("OK: ",insert)

	result, err := collection.Select(bson.M{"name": "1112"})
	if err != nil {
		log.Fatalln(err)
	}

	c := &s1{}
	err = result.Decode(c)
	if err != nil {
		log.Fatalln(err)
	}

	_,ok := c.Ad.(*a)
	fmt.Println(c.Ad)
	if !ok {
		log.Fatalln("Eeeee")
	}

	log.Println("OK")
}
