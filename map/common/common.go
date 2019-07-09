package hashmap

type MapI interface {
	Insert(key, val int)
	Get(key int)
	Remove(key int)
	// rehash()
	// getHash
}
