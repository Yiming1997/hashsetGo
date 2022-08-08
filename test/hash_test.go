package test

import (
	"fmt"
	"myapp/hashgo"
	"myapp/utils"
	"testing"
)

type User struct {
	Id   int
	Name string
}

func TestHash2(*testing.T) {
	user := &User{Id: 1, Name: "xiaoming"}
	userdataByte, _ := utils.Encode(user)
	userHashCode := utils.Hashcode(userdataByte)
	userHash := utils.Hash(userHashCode, 16)
	fmt.Println("hash...", userHash)
}

func TestHash3(*testing.T) {
	hashtable := hashgo.NewHashSet[int]()
	//user1 := User{Id: 1, Name: "xiaoming"}
	//user2 := User{Id: 1, Name: "xiaoming2"}

	hashtable.Set(1)
	hashtable.Set(2)
	hashtable.Set(53)
	hashtable.Set(63)
	hashtable.Set(73)
	hashtable.Set(83)
	hashtable.Set(100)
	hashtable.Set(101)
	hashtable.Set(155)
	hashtable.Set(178)
	hashtable.Set(200)
	hashtable.Set(203)

	hashtable.Set(355)
	hashtable.Set(478)
	//hashtable.Set(500)
	//hashtable.Set(600)

	fmt.Println("hashtable...", hashtable.GetData())
	fmt.Println("hashtable count", hashtable.Count())

	fmt.Println("after delAll........")
	hashtable.DelAll()

	fmt.Println("hashtable...", hashtable.GetData())
	fmt.Println("hashtable count", hashtable.Count())

	//
	//fmt.Println(user1)

}

func TestHash4(*testing.T) {
	//hashtable := hashgo.NewHashSet[int]()
	//user1 := User{Id: 1, Name: "xiaoming"}
	//user2 := User{Id: 1, Name: "xiaoming2"}

	arr := make([]int, 10, 20)

	var a1 int = 1
	//a3 := &a1
	//
	//*a3 = 100
	//fmt.Println(a1)

	arr[1] = a1

	var a2 *int = new(int)
	*a2 = 4

	fmt.Println("a2", *a2)

	arr1Pointer := &arr[1]
	*arr1Pointer = 100

	fmt.Println("arr1Pointer", *arr1Pointer)

	fmt.Println("arr", arr)
}

func TestHash5(*testing.T) {
	hashtable := hashgo.NewHashSet[int]()
	//user1 := &User{Id: 1, Name: "xiaoming"}
	//user2 := User{Id: 1, Name: "xiaoming2"}
	//hashtable := hashgo.NewHashSet[int]()
	hashtable.Set(1)
	hashtable.Set(2)
	hashtable.Foreach(func(item *hashgo.HashItem[int, any]) {
		item.SetKey(10000)
	})

	hashtable.Foreach(func(item *hashgo.HashItem[int, any]) {
		fmt.Println(item)
	})
}

func TestHas6(*testing.T) {
	hashtable := hashgo.NewHashSet2[int]()
	//user1 := &User{Id: 1, Name: "xiaoming"}
	//user2 := User{Id: 1, Name: "xiaoming2"}
	//hashtable := hashgo.NewHashSet[int]()
	hashtable.Set(1)
	hashtable.Set(2)

	fmt.Println(hashtable.GetLoadFactor())

	hashtable.Foreach(func(item *hashgo.HashItem[int, any]) {
		item.SetKey(10000)
	})

	hashtable.Foreach(func(item *hashgo.HashItem[int, any]) {
		fmt.Println(item)
	})

}
