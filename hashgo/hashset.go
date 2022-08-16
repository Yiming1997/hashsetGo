package hashgo

import (
	"errors"
	"myapp/constVal"
	"myapp/utils"
	"sync"
)

type Option[T comparable] func(*Hashset[T])

type Hashset[T comparable] struct {
	sync.RWMutex
	data       [][]*HashItem[T, any]
	loadFactor float64
}

func (ht *Hashset[T]) GetLoadFactor() float64 {
	return ht.loadFactor
}

func WithLoadFactor[T comparable](loadFactor float64) Option[T] {
	return func(h *Hashset[T]) {
		h.loadFactor = loadFactor
	}
}

func (ht *Hashset[T]) GetData() [][]*HashItem[T, any] {
	return ht.data
}

func NewHashSet[T comparable](options ...func(hashset *Hashset[T])) *Hashset[T] {
	hashset := Hashset[T]{
		data:       make([][]*HashItem[T, any], 16),
		loadFactor: 0.75,
	}
	for _, option := range options {
		option(&hashset)
	}

	return &hashset
}

func (ht *Hashset[T]) Set(key T) (err error) {
	newKey2Bytes, err := utils.Encode(key)
	if err != nil {
		return err
	}

	newKeyHashcode := utils.Hashcode(newKey2Bytes)                // get hashcode
	newKeyOffset := utils.Hash(newKeyHashcode, len(ht.GetData())) //get target index to store

	keyExistence, existItemIndex, err := ht.checkKeyExistence(newKeyOffset, key)

	if err != nil {
		return err
	}

	newHashItem := &HashItem[T, any]{key, 1}
	if !keyExistence {
		ht.data[newKeyOffset] = append(ht.data[newKeyOffset], newHashItem)
	} else {
		ht.data[newKeyOffset][existItemIndex] = newHashItem //replace the old key with the new
	}

	currenLoadFactor := float64(ht.Count() / len(ht.data))

	if currenLoadFactor > 0.75 {
		ht.rehash(constVal.Extend)
	}

	return nil
}

func (ht *Hashset[T]) Del(key T) (err error) {
	Key2Bytes, err := utils.Encode(key)

	if err != nil {
		return err
	}

	targetKeyHashcode := utils.Hashcode(Key2Bytes)
	targetKeyOffset := targetKeyHashcode % len(ht.GetData())

	targetKeyArr := ht.data[targetKeyOffset]
	//targetKerArr2 := &ht.data[targetKeyOffset]

	if len(targetKeyArr) == 0 {
		return errors.New("the key does not exist")
	}

	for index, item := range targetKeyArr {
		if item.GetKey() == key {
			//here deleting the target item.....
			ht.sliceItemDeleteByIndex(targetKeyArr, index, targetKeyOffset)
			break
		}
	}

	currenLoadFactor := float64(ht.Count() / len(ht.data))

	if currenLoadFactor < 0.1 {
		ht.rehash(constVal.Shrink)
	}

	return nil
}

func (ht *Hashset[T]) DelAll() {
	ht.data = make([][]*HashItem[T, any], 16)
}

func (ht *Hashset[T]) Contains(Key T) bool {
	newKey2Bytes, _ := utils.Encode(Key)
	newKeyHashcode := utils.Hashcode(newKey2Bytes)           // get hashcode
	newKeyOffset := utils.Hash(newKeyHashcode, len(ht.data)) //get target index to store
	isExist, _, _ := ht.checkKeyExistence(newKeyOffset, Key)
	return isExist
}

func (ht *Hashset[T]) Count() int {
	totalCount := 0
	for i := 0; i < len(ht.data); i++ {
		totalCount += len(ht.data[i])
	}
	return totalCount
}

func (ht *Hashset[T]) Foreach(trans func(hashtable *HashItem[T, any])) {
	for i := 0; i < len(ht.data); i++ {
		for _, v := range ht.data[i] {
			trans(v)
		}
	}
}

func (ht *Hashset[T]) checkKeyExistence(row int, key T) (bool, int, error) {
	exist := false
	existItemIndex := -1

	Key2Bytes, err := utils.Encode(key)
	if err != nil {
		return false, -1, err
	}
	newKeyHashcode := utils.Hashcode(Key2Bytes)
	for index, item := range ht.data[row] {
		newKey2Bytes, err := utils.Encode(item.key)
		if err != nil {
			return false, -1, err
		}
		itemKeyHashcode := utils.Hashcode(newKey2Bytes)

		if itemKeyHashcode == newKeyHashcode && key == item.GetKey() {
			exist = true
			existItemIndex = index
			break
		}
	}
	return exist, existItemIndex, nil
}

func (ht *Hashset[T]) sliceItemDeleteByIndex(targetArr []*HashItem[T, any], index, targetKeyOffset int) {
	ht.data[targetKeyOffset] = append(targetArr[:index], targetArr[(index+1):]...)
	return
}

// opCode 1 extend ,2 shrink
func (ht *Hashset[T]) rehash(opCode constVal.OpCode) {
	oldSetArr := ht.GetData()
	if opCode == constVal.Extend { // too many elements may affect the set performance , so extend  its capability
		ht.data = make([][]*HashItem[T, any], len(ht.data)*2)
	} else if opCode == constVal.Shrink && len(ht.data) > 16 { // too fewer elements, shrink
		ht.data = make([][]*HashItem[T, any], len(ht.data)/2)
	} else {
		return
	}

	for _, arrItem := range oldSetArr {
		for _, item := range arrItem {
			ht.Set(item.GetKey())
		}
	}
	return
}
