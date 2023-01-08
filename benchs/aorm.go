package benchs

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tangpanqing/aorm"
	"github.com/tangpanqing/aorm/builder"
	"github.com/tangpanqing/aorm/null"
)

var ao *sql.DB

func initDB2Aorm() {
	sqls := []string{
		`DROP TABLE IF EXISTS aorm_model;`,
		"CREATE TABLE aorm_model (" +
			"id int NOT NULL AUTO_INCREMENT," +
			"name text NOT NULL," +
			"title text NOT NULL," +
			"fax text NOT NULL," +
			"web text NOT NULL," +
			"age int NOT NULL," +
			"right_val boolean NOT NULL," +
			"counter int NOT NULL," +
			"CONSTRAINT model_pkey PRIMARY KEY (`id`)" +
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
	Id       null.Int
	Name     null.String
	Title    null.String
	Fax      null.String
	Web      null.String
	Age      null.Int
	RightVal null.Bool
	Counter  null.Int
}

func NewAormModel() AormModel {
	m := AormModel{
		Name:     null.StringFrom("Orm Benchmark"),
		Title:    null.StringFrom("Just a Benchmark for fun"),
		Fax:      null.StringFrom("99909990"),
		Web:      null.StringFrom("http://blog.milkpod29.me"),
		Age:      null.IntFrom(100),
		RightVal: null.BoolFrom(true),
		Counter:  null.IntFrom(1000),
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
	var ms []AormModel
	wrapExecute(b, func() {
		initDB2Aorm()
		ms = make([]AormModel, 0, 100)
		for i := 0; i < 100; i++ {
			ms = append(ms, NewAormModel())
		}
	})
	for i := 0; i < b.N; i++ {
		if _, err := aorm.Use(ao).InsertBatch(&ms); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
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
		if _, err := aorm.Use(ao).Where(&AormModel{Id: null.IntFrom(lastId)}).Update(&m); err != nil {
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
		if err := aorm.Use(ao).Table("aorm_model").Where(&AormModel{Id: null.IntFrom(lastId)}).GetOne(&mm); err != nil {
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
		var where []builder.WhereItem
		where = append(where, builder.WhereItem{Field: "id", Opt: builder.Gt, Val: 0})
		if err := aorm.Use(ao).Table("aorm_model").WhereArr(where).Limit(0, 100).GetMany(&models); err != nil {
			fmt.Println(err)
			b.FailNow()
		}

		//for i := 0; i < len(models); i++ {
		//	fmt.Println(models[i])
		//}
	}
}
