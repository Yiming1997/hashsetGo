package hashgo

import (
	"errors"
	"myapp/utils"
)

type Hashset[T comparable] struct {
	data [][]*HashItem[T, any]
}

func (ht *Hashset[T]) GetData() [][]*HashItem[T, any] {
	return ht.data
}

func NewHashSet[T comparable]() *Hashset[T] {
	hashTable := &Hashset[T]{}
	hashTable.data = make([][]*HashItem[T, any], 16)
	return hashTable
}

func (ht *Hashset[T]) Set(key T) (err error) {
	newKey2Bytes, err := utils.Encode(key)
	if err != nil {
		return err
	}

	newKeyHashcode := utils.Hashcode(newKey2Bytes) // get hashcode
	newKeyOffset := utils.Hash(newKeyHashcode, 16) //get target index to store

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
	return nil
}

func (ht *Hashset[T]) Del(key T) (err error) {
	Key2Bytes, err := utils.Encode(key)

	if err != nil {
		return err
	}

	targetKeyHashcode := utils.Hashcode(Key2Bytes)
	targetKeyOffset := targetKeyHashcode % 16

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
	return nil
}

func (ht *Hashset[T]) DelAll() {
	ht.data = make([][]*HashItem[T, any], 16)
}

func (ht *Hashset[T]) Contains(Key T) bool {
	newKey2Bytes, _ := utils.Encode(Key)
	newKeyHashcode := utils.Hashcode(newKey2Bytes) // get hashcode
	newKeyOffset := utils.Hash(newKeyHashcode, 16) //get target index to store
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
