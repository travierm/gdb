package main

type MapDatabase struct {
	// name of the database
	Name string

	// matches each key to the correct index in the store
	Index map[string]int

	// holds the databases information
	Store map[int]interface{}

	// total number of keys in use
	IndexCount int
}

func (mdb *MapDatabase) Add(key string, value interface{}) {
	mdb.initDB()

	mdb.IndexCount++
	mdb.Index[key] = mdb.IndexCount
	mdb.Store[mdb.IndexCount] = value
}

func (mdb *MapDatabase) Get(key string) interface{} {
	mdb.initDB()

	storeIndex := mdb.getIndexByKey(key)

	return mdb.Store[storeIndex]
}

func (mdb *MapDatabase) initDB() {
	if mdb.Index == nil {
		mdb.Index = make(map[string]int)
	}

	if mdb.Store == nil {
		mdb.Store = make(map[int]interface{})
	}
}

func (mdb *MapDatabase) getIndexByKey(key string) int {
	return mdb.Index[key]
}
