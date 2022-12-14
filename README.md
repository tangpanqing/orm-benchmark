# ORM Benchmark

_(forked from https://github.com/yusaer/orm-benchmark)_

### Environment

- go version go1.18 windows

### Mysql

- Mysql 8 for Windows on x86_64

### ORMs

All package run in no-cache mode.

- [beego/orm](https://github.com/astaxie/beego/tree/master/orm)
- [gorm](https://github.com/go-gorm/gorm)
- [xorm](https://github.com/xormplus/xorm)
- [aorm](https://github.com/tangpanqing/aorm)

See [`go.mod`](https://github.com/frederikhors/orm-benchmark/blob/master/go.mod) for their latest versions.

### Run

```shell
go get github.com/frederikhors/orm-benchmark
# build
go install
# all
orm-benchmark -multi=20 -orm=all
# portion
orm-benchmark -multi=20 -orm=gorm
orm-benchmark -multi=20 -orm=aorm
orm-benchmark -multi=20 -orm=raw
# ... and so on...
```

### Results

From [`results.md`](https://github.com/tangpanqing/orm-benchmark/tree/master/results.md):

**orm-benchmark -multi=20**

```
  4000 times - Insert
  raw_stmt:     3.20s       799738 ns/op     288 B/op     10 allocs/op
     beego:     3.21s       802146 ns/op     905 B/op     36 allocs/op
 gorm_stmt:     3.23s       806765 ns/op    2198 B/op     43 allocs/op
      gorm:     3.32s       830179 ns/op    2410 B/op     44 allocs/op
       raw:     3.36s       838917 ns/op     505 B/op     12 allocs/op
      aorm:     3.64s       908754 ns/op    1721 B/op     72 allocs/op
      xorm:     3.68s       919397 ns/op    1666 B/op     48 allocs/op

  4000 times - MultiInsert 100 row
  raw_stmt:     5.85s      1463672 ns/op   65887 B/op    511 allocs/op
 gorm_stmt:     6.12s      1529892 ns/op   80833 B/op   1138 allocs/op
       raw:     6.77s      1691606 ns/op   88167 B/op    514 allocs/op
     beego:     7.00s      1750787 ns/op   91316 B/op   1435 allocs/op
      gorm:     7.15s      1787393 ns/op  101319 B/op   1140 allocs/op
      xorm:     7.30s      1825168 ns/op  155230 B/op   2850 allocs/op
      aorm:     doesn't work

  4000 times - Update
  raw_stmt:     0.12s        28929 ns/op     304 B/op     10 allocs/op
     beego:     0.14s        34485 ns/op     881 B/op     36 allocs/op
 gorm_stmt:     0.14s        35906 ns/op    2656 B/op     55 allocs/op
       raw:     0.22s        55383 ns/op     536 B/op     12 allocs/op
      xorm:     0.25s        62790 ns/op    1664 B/op     85 allocs/op
      aorm:     0.26s        64613 ns/op    1929 B/op     85 allocs/op
      gorm:     0.28s        69908 ns/op    2888 B/op     57 allocs/op

  4000 times - Read
  raw_stmt:     0.12s        29174 ns/op     833 B/op     32 allocs/op
     beego:     0.12s        30973 ns/op    1616 B/op     88 allocs/op
 gorm_stmt:     0.15s        37358 ns/op    2261 B/op     66 allocs/op
       raw:     0.22s        55575 ns/op     897 B/op     35 allocs/op
      gorm:     0.27s        67097 ns/op    2325 B/op     69 allocs/op
      aorm:     0.27s        68227 ns/op    2922 B/op    140 allocs/op
      xorm:     0.37s        92023 ns/op    5900 B/op    226 allocs/op

  4000 times - MultiRead limit 100
       raw:     0.75s       186707 ns/op   26736 B/op   1509 allocs/op
  raw_stmt:     0.79s       198145 ns/op   21152 B/op   1020 allocs/op
     beego:     1.16s       289951 ns/op   51738 B/op   3783 allocs/op
 gorm_stmt:     1.20s       299052 ns/op   31599 B/op   2059 allocs/op
      gorm:     1.33s       332463 ns/op   31662 B/op   2062 allocs/op
      aorm:     2.75s       687571 ns/op  135109 B/op   8744 allocs/op
      xorm:     doesn't work
```
