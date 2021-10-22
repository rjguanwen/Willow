// @Title  请填写文件名称（需要改）
// @Description  请填写文件描述（需要改）
// @Author  zhengbin  2021/9/8 16:52
package sample_1

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/go-memdb"
)

// Create a sample struct
type Person struct {
	Email 		string	`json:"email"`
	Name  		string	`json:"name"`
	Age   		int		`json:"age"`
	Sex   		int		`json:"sex"`
	NickName	string	`json:"nickname"`
}

func OptMemdb() {
	// Create the DB schema
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"person": &memdb.TableSchema{
				Name: "person",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Email"},
					},
					"age": &memdb.IndexSchema{
						Name:    "age",
						Unique:  false,
						Indexer: &memdb.IntFieldIndex{Field: "Age"},
					},
					"nickname": &memdb.IndexSchema{
						Name:	"nickname",
						Unique: false,
						Indexer: &memdb.StringFieldIndex{Field: "NickName"},
						AllowMissing: true,
					},
				},
			},
		},
	}

	// Create a new data base
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}

	// Create a write transaction
	txn := db.Txn(true)

	// Insert some people
	people := []*Person{
		&Person{"joe@aol.com", "Joe", 30, 0, "j"},
		&Person{"lucy@aol.com", "Lucy", 35, 0, "l"},
		&Person{"tariq@aol.com", "Tariq", 21, 1, ""},
		&Person{"dorothy@aol.com", "Dorothy", 53, 1, "d"},
	}

	// 使用字符串生成对象
	zhangsan := `{"email": "zhangsan@163.com", "name": "zhangsan", "age": 28}`
	var p_zhangsan Person
	json.Unmarshal([]byte(zhangsan), &p_zhangsan)
	people = append(people, &p_zhangsan)

	for _, p := range people {
		if err := txn.Insert("person", p); err != nil {
			panic(err)
		}
	}

	// Commit the transaction
	txn.Commit()

	// Create read-only transaction
	txn = db.Txn(false)
	defer txn.Abort()

	// Lookup by email
	raw, err := txn.First("person", "id", "joe@aol.com")
	if err != nil {
		panic(err)
	}

	// Say hi!
	fmt.Printf("Hello %s!\n", raw.(*Person).Name)

	// List all the people
	it, err := txn.Get("person", "id")
	if err != nil {
		panic(err)
	}

	fmt.Println("All the people:")
	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(*Person)
		fmt.Printf("  %s: %d: %s\n", p.Name, p.Sex, p.NickName)
	}

	// Range scan over people with ages between 25 and 35 inclusive
	it, err = txn.LowerBound("person", "age", 32)
	if err != nil {
		panic(err)
	}

	fmt.Println("People aged 25 - 35:")
	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(*Person)
		if p.Age > 35 {
			break
		}
		fmt.Printf("  %s is aged %d\n", p.Name, p.Age)
	}

	fmt.Println("多条件查询...")
	it, err = txn.LowerBound("person", "age", 25)
	if err != nil {
		panic(err)
	}
}