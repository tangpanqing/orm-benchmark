package benchs

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tangpanqing/aorm"
)

var ao *sql.DB

func initDB2Aorm() {
	sqls := []string{
		"DROP TABLE IF EXISTS aorm_model;",
		"CREATE TABLE aorm_model (" +
			"id int NOT NULL AUTO_INCREMENT," +
			"name text NOT NULL," +
			"title text NOT NULL," +
			"fax text NOT NULL," +
			"web text NOT NULL," +
			"age int NOT NULL," +
			"right_val boolean NOT NULL," +
			"counter int NOT NULL," +
			"PRIMARY KEY (`id`)" +
			");",
	}

	DB, err := sql.Open("mysql", OrmSource)
	checkErr(err)
	defer func() {
		err := DB.Close()
		checkErr(err)
	}()

	err = DB.Ping()
	checkErr(err)

	for _, line := range sqls {
		_, err = DB.Exec(line)
		checkErr(err)
	}
}

type AormModel struct {
	Id       aorm.Int
	Name     aorm.String
	Title    aorm.String
	Fax      aorm.String
	Web      aorm.String
	Age      aorm.Int
	RightVal aorm.Bool
	Counter  aorm.Int
}

func NewAormModel() AormModel {
	m := AormModel{
		Name:     aorm.StringFrom("Orm Benchmark"),
		Title:    aorm.StringFrom("Just a Benchmark for fun"),
		Fax:      aorm.StringFrom("99909990"),
		Web:      aorm.StringFrom("http://blog.milkpod29.me"),
		Age:      aorm.IntFrom(100),
		RightVal: aorm.BoolFrom(true),
		Counter:  aorm.IntFrom(1000),
	}

	return m
}

func init() {
	st := NewSuite("aorm")
	st.InitF = func() {
		st.AddBenchmark("Insert", 200*OrmMulti, AormInsert)
		st.AddBenchmark("MultiInsert 100 row", 200*OrmMulti, AormInsertMulti)
		st.AddBenchmark("Update", 200*OrmMulti, AormUpdate)
		st.AddBenchmark("Read", 200*OrmMulti, AormRead)
		st.AddBenchmark("MultiRead limit 100", 200*OrmMulti, AormReadSlice)

		engine, err := sql.Open("mysql", OrmSource)
		if err != nil {
			fmt.Print(err)
		}

		// engine.SetMaxIdleConns(ORM_MAX_IDLE)
		// engine.SetMaxOpenConns(ORM_MAX_CONN)

		ao = engine
	}
}

func AormInsert(b *B) {
	var m AormModel
	wrapExecute(b, func() {
		initDB2Aorm()
		m = NewAormModel()
	})

	for i := 0; i < b.N; i++ {
		if _, err := aorm.Use(ao).Insert(&m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func AormInsertMulti(b *B) {
	panic(fmt.Errorf("doesn't work"))
	//var ms []*AormModel
	//wrapExecute(b, func() {
	//	initDB2Aorm()
	//	ms = make([]*AormModel, 0, 100)
	//	for i := 0; i < 100; i++ {
	//		ms = append(ms, NewAormModel())
	//	}
	//})
	//for i := 0; i < b.N; i++ {
	//	for _, m := range ms {
	//		m.Id = 0
	//	}
	//	if _, err := ao.InsertMulti(&ms); err != nil {
	//		fmt.Println(err)
	//		b.FailNow()
	//	}
	//}
}

func AormUpdate(b *B) {
	var m AormModel
	var lastId int64
	wrapExecute(b, func() {
		initDB2Aorm()
		m = NewAormModel()
		id, err := aorm.Use(ao).Insert(&m)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
		lastId = id
	})

	for i := 0; i < b.N; i++ {
		if _, err := aorm.Use(ao).Where(&AormModel{Id: aorm.IntFrom(lastId)}).Update(&m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func AormRead(b *B) {
	var lastId int64
	wrapExecute(b, func() {
		initDB2Aorm()
		m := NewAormModel()
		id, err := aorm.Use(ao).Insert(&m)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
		lastId = id
	})

	var mm AormModel
	for i := 0; i < b.N; i++ {
		if err := aorm.Use(ao).Where(&AormModel{Id: aorm.IntFrom(lastId)}).GetOne(&mm); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func AormReadSlice(b *B) {
	var m AormModel
	wrapExecute(b, func() {
		initDB2Aorm()
		m = NewAormModel()
		for i := 0; i < 100; i++ {
			if _, err := aorm.Use(ao).Insert(&m); err != nil {
				fmt.Println(err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		var models []AormModel
		var where []aorm.WhereItem
		where = append(where, aorm.WhereItem{Field: "id", Opt: aorm.Gt, Val: 0})
		if err := aorm.Use(ao).Table("aorm_model").WhereArr(where).Limit(0, 100).GetManyNew(&models); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}
