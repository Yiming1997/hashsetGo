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

	hashtable.Set(3)
	hashtable.Set(4)

	hashtable.Set(5)
	hashtable.Set(6)

	hashtable.Set(7)
	hashtable.Set(8)

	hashtable.Set(9)
	hashtable.Set(10)

	hashtable.Set(11)
	hashtable.Set(12)

	hashtable.Set(13)
	hashtable.Set(14)

	hashtable.Set(15)
	hashtable.Set(16)

	hashtable.Set(17)
	hashtable.Set(18)

	hashtable.Set(19)
	hashtable.Set(20)

	//hashtable.Set(500)
	//hashtable.Set(600)

	fmt.Println("hashtable...", hashtable.GetData())
	fmt.Println("hashtable count", hashtable.Count())
	fmt.Println("bucket len", len(hashtable.GetData()))

	hashtable.Del(1)
	hashtable.Del(2)
	hashtable.Del(3)
	hashtable.Del(4)
	hashtable.Del(5)
	hashtable.Del(6)
	hashtable.Del(7)
	hashtable.Del(8)
	hashtable.Del(9)
	hashtable.Del(10)
	hashtable.Del(11)
	hashtable.Del(12)
	hashtable.Del(13)
	hashtable.Del(14)
	hashtable.Del(15)
	hashtable.Del(16)
	hashtable.Del(17)
	hashtable.Del(18)
	fmt.Println("after deleted...............................................................................")

	fmt.Println("hashtable...", hashtable.GetData())
	fmt.Println("hashtable count", hashtable.Count())
	fmt.Println("bucket len", len(hashtable.GetData()))

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
	hashtable := hashgo.NewHashSet[int]()
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

func TestHas7(*testing.T) {
	data := make([]any, 16)

	data[1] = make([]int, 16)

	data[2] = "v"

}
