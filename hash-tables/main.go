package main

import "fmt"

//struct hash table with an array with fixed size

//struct bucket to store nodes

//struct node with value to be stored

// hash function to convert string into index number
const ArraySize = 7

var hashTable *HashTable

type HashTable struct {
	array [ArraySize]*Bucket
}

type Bucket struct {
	head *BucketNode
}

type BucketNode struct {
	key  string
	next *BucketNode
}

func (h *HashTable) Insert(value string) {
	index := hash(value)
	h.array[index].Insert(value)
}

func (h *HashTable) Search(key string) {
	index := hash(key)

	h.array[index].Search(key)
}

func (b *Bucket) Insert(key string) {
	newNode := &BucketNode{key: key}
	newNode.next = b.head
	b.head = newNode
}

func (b *Bucket) Search(key string) {
	current := b.head

	for current != nil {
		if current.key == key {
			fmt.Println("Encontrei o node", current, "com valor", current.key)
			return
		}

		current = current.next
	}

	fmt.Println("Não entrei nenhum node com a chave", key)
}

func hash(value string) int {
	sum := 0

	for _, v := range value {
		sum += int(v)
	}
	//fmt.Println("O hash index é", sum%ArraySize)
	return sum % ArraySize
}

func init() {
	hashTable = &HashTable{}

	for i := range hashTable.array {
		hashTable.array[i] = &Bucket{}
	}
}

func main() {
	fmt.Println(hashTable)

	hashTable.Insert("RANDY")
	hashTable.Insert("OTIS")
	hashTable.Insert("FRANCIS")
	hashTable.Insert("ERIC")
	hashTable.Insert("WILIAM")

	fmt.Println(
		hashTable.array[4].head,
		hashTable.array[4].head.next,
		hashTable.array[4].head.next.next,
		hashTable.array[0].head,
		hashTable.array[3].head,
		hashTable.array[3].head.next,
	)

	hashTable.Search("WESLEY")
}
