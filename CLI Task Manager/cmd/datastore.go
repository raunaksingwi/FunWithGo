package cmd

import (
	"encoding/binary"
	"fmt"

	bolt "go.etcd.io/bbolt"
)

const dbPath = "./Resources/todo.db"

type datastore interface {
	addTask(task string)
	rmTask(taskID int)
	doTask(taskID int)

	listPendingTasks(taskID int) map[int]string
	listCompletedTasks(taskID int) map[int]string
}

type database struct {
	db *bolt.DB
}

func getBboltDatabase() database {
	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		panic(err)
	}
	return database{
		db: db,
	}
}

func (d *database) addTask(task string) {

	err := d.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("todo"))
		return err
	})

	if err != nil {
		panic(err)
	}

	err = d.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todo"))

		id, _ := b.NextSequence()
		return b.Put(itob(id), []byte(task))
	})

	if err != nil {
		fmt.Println("Paniking in add task")
		panic(err)
	}

}

func (d *database) rmTask(taskID int) {

	key := getKeyForTaskID(taskID, d.db)
	err := d.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todo"))
		return b.Delete(key)
	})

	if err != nil {
		panic(err)
	}
}

func (d *database) doTask(taskID int) {

	err := d.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("completed"))
		return err
	})

	if err != nil {
		panic(err)
	}

	key := getKeyForTaskID(taskID, d.db)

	if key == nil {
		fmt.Println("Given taskID does not exist")
		return
	}
	var value []byte
	err = d.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todo"))
		value = b.Get(key)
		return nil
	})

	d.rmTask(taskID)

	err = d.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("completed"))
		b.Put(key, value)
		return nil
	})

}

func (d *database) listPendingTasks() {
	listTask("todo", d.db)
}

func (d *database) listCompletedTasks() {
	listTask("completed", d.db)
}

func listTask(bucketName string, db *bolt.DB) {
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		c := b.Cursor()
		i := 1
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("%d. %s\n", i, v)
			i++
		}
		return nil
	})

	if err != nil {
		panic(err)
	}
}

func getKeyForTaskID(taskID int, db *bolt.DB) []byte {
	var key []byte
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todo"))
		c := b.Cursor()
		i := 1
		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			if i == taskID {
				key = k
				break
			}
			i++
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	return key
}

func itob(i uint64) []byte {
	ret := make([]byte, 8)
	binary.BigEndian.PutUint64(ret, i)
	return ret
}
