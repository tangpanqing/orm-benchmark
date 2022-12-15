package benchs

import (
	"fmt"

	"database/sql"

	"github.com/astaxie/beego/orm"
)

var bo orm.Ormer

func initDB3() {

	sqls := []string{
		`DROP TABLE IF EXISTS beego_model;`,
		"CREATE TABLE beego_model (" +
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

type BeegoModel struct {
	Id       int `orm:"auto"`
	Name     string
	Title    string
	Fax      string
	Web      string
	Age      int
	RightVal bool
	Counter  int64
}

func NewBeegoModel() *BeegoModel {
	m := new(BeegoModel)
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
	st := NewSuite("beego")
	st.InitF = func() {
		st.AddBenchmark("Insert", 200*OrmMulti, BeegoOrmInsert)
		st.AddBenchmark("MultiInsert 100 row", 200*OrmMulti, BeegoOrmInsertMulti)
		st.AddBenchmark("Update", 200*OrmMulti, BeegoOrmUpdate)
		st.AddBenchmark("Read", 200*OrmMulti, BeegoOrmRead)
		st.AddBenchmark("MultiRead limit 100", 200*OrmMulti, BeegoOrmReadSlice)

		err := orm.RegisterDataBase("default", "mysql", OrmSource)
		checkErr(err)
		orm.RegisterModel(new(BeegoModel))

		bo = orm.NewOrm()
	}
}

func BeegoOrmInsert(b *B) {
	var m *BeegoModel
	wrapExecute(b, func() {
		initDB3()
		m = NewBeegoModel()
	})

	for i := 0; i < b.N; i++ {
		m.Id = 0
		if _, err := bo.Insert(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func BeegoOrmInsertMulti(b *B) {
	var ms []*BeegoModel
	wrapExecute(b, func() {
		initDB3()
		ms = make([]*BeegoModel, 0, 100)
		for i := 0; i < 100; i++ {
			ms = append(ms, NewBeegoModel())
		}
	})

	for i := 0; i < b.N; i++ {
		if _, err := bo.InsertMulti(100, ms); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func BeegoOrmUpdate(b *B) {
	var m *BeegoModel
	wrapExecute(b, func() {
		initDB3()
		m = NewBeegoModel()
		if _, err := bo.Insert(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		if _, err := bo.Update(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func BeegoOrmRead(b *B) {
	var m *BeegoModel
	wrapExecute(b, func() {
		initDB3()
		m = NewBeegoModel()
		if _, err := bo.Insert(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		if err := bo.Read(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func BeegoOrmReadSlice(b *B) {
	var m *BeegoModel
	wrapExecute(b, func() {
		initDB3()
		m = NewBeegoModel()
		for i := 0; i < 100; i++ {
			m.Id = 0
			if _, err := bo.Insert(m); err != nil {
				fmt.Println(err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		var models []*BeegoModel
		if _, err := bo.QueryTable("beego_model").Filter("id__gt", 0).Limit(100).All(&models); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}
