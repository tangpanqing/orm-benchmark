package benchs

import (
	"fmt"

	"database/sql"

	"xorm.io/xorm"
)

var xo *xorm.Session

func initDB2() {

	sqls := []string{
		`DROP TABLE IF EXISTS xorm_model;`,
		`CREATE TABLE xorm_model (
		id integer NOT NULL,
		name text NOT NULL,
		title text NOT NULL,
		fax text NOT NULL,
		web text NOT NULL,
		age integer NOT NULL,
		right_val boolean NOT NULL,
		counter bigint NOT NULL
		);`,
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

type XormModel struct {
	Id       int
	Name     string
	Title    string
	Fax      string
	Web      string
	Age      int
	RightVal bool
	Counter  int64
}

func NewXormModel() *XormModel {
	m := new(XormModel)
	m.Name = "Orm Benchmark"
	m.Title = "Just a Benchmark for fun"
	m.Fax = "99909990"
	m.Web = "http://blog.milkpod29.me"
	m.Age = 100
	m.RightVal = true
	m.Counter = 1000

	return m
}

func init() {
	st := NewSuite("xorm")
	st.InitF = func() {
		st.AddBenchmark("Insert", 200*OrmMulti, XormInsert)
		st.AddBenchmark("MultiInsert 100 row", 200*OrmMulti, XormInsertMulti)
		st.AddBenchmark("Update", 200*OrmMulti, XormUpdate)
		st.AddBenchmark("Read", 200*OrmMulti, XormRead)
		st.AddBenchmark("MultiRead limit 100", 200*OrmMulti, XormReadSlice)

		engine, err := xorm.NewEngine("mysql", OrmSource)
		if err != nil {
			fmt.Print(err)
		}

		// engine.SetMaxIdleConns(ORM_MAX_IDLE)
		// engine.SetMaxOpenConns(ORM_MAX_CONN)

		xo = engine.NewSession()
	}
}

func XormInsert(b *B) {
	var m *XormModel
	wrapExecute(b, func() {
		initDB2()
		m = NewXormModel()
	})

	for i := 0; i < b.N; i++ {
		m.Id = 0
		if _, err := xo.Insert(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func XormInsertMulti(b *B) {
	var ms []*XormModel
	wrapExecute(b, func() {
		initDB2()
		ms = make([]*XormModel, 0, 100)
		for i := 0; i < 100; i++ {
			ms = append(ms, NewXormModel())
		}
	})
	for i := 0; i < b.N; i++ {
		for _, m := range ms {
			m.Id = 0
		}
		if _, err := xo.InsertMulti(&ms); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func XormUpdate(b *B) {
	var m *XormModel
	wrapExecute(b, func() {
		initDB2()
		m = NewXormModel()
		if _, err := xo.Insert(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		if _, err := xo.Update(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func XormRead(b *B) {
	var m *XormModel
	wrapExecute(b, func() {
		initDB2()
		m = NewXormModel()
		if _, err := xo.Insert(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		if _, err := xo.Get(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func XormReadSlice(_ *B) {
	panic(fmt.Errorf("doesn't work"))
	//var m *XormModel
	//wrapExecute(b, func() {
	//	initDB2()
	//	m = NewXormModel()
	//	for i := 0; i < 100; i++ {
	//		m.Id = 0
	//		if _, err := xo.Insert(m); err != nil {
	//			fmt.Println(err)
	//			b.FailNow()
	//		}
	//	}
	//})
	//
	//for i := 0; i < b.N; i++ {
	//	panic(fmt.Errorf("doesn't work"))
	//	var models []XormModel
	//	if err := xo.Table("xorm_model").Where("id > ?", 0).Limit(100).Find(&models); err != nil {
	//		fmt.Println(err)
	//		b.FailNow()
	//	}
	//}
}
