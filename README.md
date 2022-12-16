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
- [aorm v1.0.11](https://github.com/tangpanqing/aorm)

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
      xorm:     3.25s       812484 ns/op    1666 B/op     48 allocs/op
 gorm_stmt:     3.35s       837466 ns/op    2195 B/op     42 allocs/op
     beego:     3.35s       838498 ns/op     905 B/op     36 allocs/op
  raw_stmt:     3.46s       866242 ns/op     288 B/op     10 allocs/op
      gorm:     3.54s       883939 ns/op    2413 B/op     45 allocs/op
      aorm:     3.55s       886341 ns/op    1721 B/op     72 allocs/op
       raw:     3.58s       895178 ns/op     505 B/op     12 allocs/op

  4000 times - MultiInsert 100 row
  raw_stmt:     5.85s      1462246 ns/op   65887 B/op    511 allocs/op
 gorm_stmt:     6.15s      1537828 ns/op   80831 B/op   1138 allocs/op
     beego:     6.53s      1632810 ns/op   91343 B/op   1435 allocs/op
       raw:     6.70s      1675391 ns/op   88166 B/op    514 allocs/op
      gorm:     6.98s      1744611 ns/op  101320 B/op   1140 allocs/op
      xorm:     7.28s      1820212 ns/op  155224 B/op   2850 allocs/op
      aorm:     doesn't work

  4000 times - Update
  raw_stmt:     0.10s        25554 ns/op     304 B/op     10 allocs/op
     beego:     0.14s        35880 ns/op     880 B/op     36 allocs/op
 gorm_stmt:     0.17s        41283 ns/op    2656 B/op     55 allocs/op
       raw:     0.21s        52570 ns/op     536 B/op     12 allocs/op
      xorm:     0.25s        62027 ns/op    1665 B/op     85 allocs/op
      gorm:     0.25s        63400 ns/op    2888 B/op     57 allocs/op
      aorm:     0.27s        67571 ns/op    1929 B/op     85 allocs/op

  4000 times - Read
  raw_stmt:     0.11s        27649 ns/op     833 B/op     32 allocs/op
     beego:     0.13s        31841 ns/op    1616 B/op     88 allocs/op
 gorm_stmt:     0.14s        34574 ns/op    2261 B/op     66 allocs/op
       raw:     0.22s        54177 ns/op     897 B/op     35 allocs/op
      aorm:     0.23s        57623 ns/op    1929 B/op     86 allocs/op
      gorm:     0.24s        59979 ns/op    2325 B/op     69 allocs/op
      xorm:     0.33s        83659 ns/op    5901 B/op    226 allocs/op

  4000 times - MultiRead limit 100
  raw_stmt:     0.66s       166103 ns/op   21152 B/op   1020 allocs/op
       raw:     0.72s       180008 ns/op   26736 B/op   1509 allocs/op
     beego:     1.16s       291017 ns/op   51738 B/op   3783 allocs/op
 gorm_stmt:     1.20s       299639 ns/op   31598 B/op   2059 allocs/op
      aorm:     1.21s       303194 ns/op   61705 B/op   3057 allocs/op
      gorm:     1.34s       335818 ns/op   31661 B/op   2062 allocs/op
      xorm:     doesn't work

```
