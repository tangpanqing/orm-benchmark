# ORM Benchmark

_(forked from https://github.com/yusaer/orm-benchmark)_

### Environment

- go version go1.18 windows

### Mysql

- Mysql 8 for Windows on x86_64

### ORMs

All package run in no-cache mode.

- [beego/orm v1.12.3](https://github.com/astaxie/beego/tree/master/orm)
- [gorm v1.24.2](https://github.com/go-gorm/gorm)
- [xorm v1.3.2](https://github.com/xormplus/xorm)
- [aorm v1.0.9](https://github.com/tangpanqing/aorm)

See [`go.mod`](https://github.com/frederikhors/orm-benchmark/blob/master/go.mod) for their latest versions.

### Run

```shell
go get github.com/tangpanqing/orm-benchmark
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
      gorm:     3.25s       812676 ns/op    2413 B/op     45 allocs/op
     beego:     3.31s       827127 ns/op     904 B/op     36 allocs/op
  raw_stmt:     3.52s       879867 ns/op     288 B/op     10 allocs/op
       raw:     3.55s       886312 ns/op     505 B/op     12 allocs/op
      xorm:    18.31s      4578624 ns/op    1666 B/op     48 allocs/op
      aorm:    18.43s      4608312 ns/op    1721 B/op     72 allocs/op
 gorm_stmt:    18.86s      4715925 ns/op    2195 B/op     42 allocs/op

  4000 times - MultiInsert 100 row
  raw_stmt:     5.79s      1448354 ns/op   65888 B/op    511 allocs/op
     beego:     6.57s      1642322 ns/op   91320 B/op   1435 allocs/op
       raw:     6.80s      1699080 ns/op   88166 B/op    514 allocs/op
      gorm:     7.36s      1840108 ns/op  101320 B/op   1140 allocs/op
      xorm:    22.79s      5698691 ns/op  155228 B/op   2850 allocs/op
 gorm_stmt:    23.01s      5751913 ns/op   80831 B/op   1138 allocs/op
      aorm:     doesn't work

  4000 times - Update
  raw_stmt:     0.11s        27475 ns/op     304 B/op     10 allocs/op
     beego:     0.14s        33995 ns/op     880 B/op     36 allocs/op
 gorm_stmt:     0.14s        35409 ns/op    2656 B/op     55 allocs/op
      xorm:     0.23s        56513 ns/op    1665 B/op     85 allocs/op
       raw:     0.23s        57128 ns/op     536 B/op     12 allocs/op
      gorm:     0.27s        66415 ns/op    2888 B/op     57 allocs/op
      aorm:     0.27s        68278 ns/op    1929 B/op     85 allocs/op

  4000 times - Read
  raw_stmt:     0.12s        28939 ns/op     833 B/op     32 allocs/op
     beego:     0.13s        32930 ns/op    1616 B/op     88 allocs/op
 gorm_stmt:     0.14s        35487 ns/op    2261 B/op     66 allocs/op
       raw:     0.24s        59106 ns/op     897 B/op     35 allocs/op
      aorm:     0.24s        59886 ns/op    1929 B/op     86 allocs/op
      gorm:     0.25s        61906 ns/op    2325 B/op     69 allocs/op
      xorm:     0.35s        87040 ns/op    5900 B/op    226 allocs/op

  4000 times - MultiRead limit 100
  raw_stmt:     0.71s       177036 ns/op   21152 B/op   1020 allocs/op
       raw:     0.72s       179353 ns/op   26736 B/op   1509 allocs/op
 gorm_stmt:     1.20s       300717 ns/op   31599 B/op   2059 allocs/op
      aorm:     1.24s       309598 ns/op   61705 B/op   3057 allocs/op
     beego:     1.28s       320622 ns/op   51736 B/op   3783 allocs/op
      gorm:     1.34s       335147 ns/op   31662 B/op   2062 allocs/op
      xorm:     doesn't work
```
