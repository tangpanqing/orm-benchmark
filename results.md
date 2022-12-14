# Results

- orm-benchmark (with no flags)

```
   200 times - Insert
  raw_stmt:     0.16s       780198 ns/op     288 B/op     10 allocs/op
 gorm_stmt:     0.16s       806995 ns/op    2469 B/op     43 allocs/op
      xorm:     0.17s       861111 ns/op    1715 B/op     48 allocs/op
      aorm:     0.17s       863662 ns/op    1757 B/op     72 allocs/op
      gorm:     0.18s       879347 ns/op    2612 B/op     44 allocs/op
       raw:     0.18s       880385 ns/op     542 B/op     12 allocs/op
     beego:     0.18s       901736 ns/op     913 B/op     36 allocs/op

   200 times - MultiInsert 100 row
  raw_stmt:     0.33s      1662320 ns/op   66330 B/op    514 allocs/op
 gorm_stmt:     0.34s      1697680 ns/op   80838 B/op   1137 allocs/op
     beego:     0.35s      1754010 ns/op   91311 B/op   1435 allocs/op
       raw:     0.38s      1892128 ns/op   88600 B/op    517 allocs/op
      gorm:     0.38s      1895111 ns/op  101305 B/op   1139 allocs/op
      xorm:     0.39s      1973717 ns/op  155282 B/op   2850 allocs/op
      aorm:     doesn't work

   200 times - Update
  raw_stmt:     0.01s        29708 ns/op     304 B/op     10 allocs/op
     beego:     0.01s        34786 ns/op     884 B/op     36 allocs/op
 gorm_stmt:     0.01s        47762 ns/op    2657 B/op     55 allocs/op
       raw:     0.01s        56908 ns/op     536 B/op     12 allocs/op
      xorm:     0.01s        61562 ns/op    1671 B/op     85 allocs/op
      aorm:     0.01s        65404 ns/op    1936 B/op     85 allocs/op
      gorm:     0.01s        70750 ns/op    2888 B/op     57 allocs/op

   200 times - Read
  raw_stmt:     0.01s        30073 ns/op     852 B/op     32 allocs/op
     beego:     0.01s        35550 ns/op    1617 B/op     88 allocs/op
       raw:     0.01s        49370 ns/op     916 B/op     35 allocs/op
 gorm_stmt:     0.01s        66536 ns/op    2299 B/op     66 allocs/op
      aorm:     0.01s        68598 ns/op    2930 B/op    140 allocs/op
      gorm:     0.01s        70856 ns/op    2353 B/op     69 allocs/op
      xorm:     0.02s        92196 ns/op    5926 B/op    226 allocs/op

   200 times - MultiRead limit 100
  raw_stmt:     0.03s       163728 ns/op   21152 B/op   1020 allocs/op
       raw:     0.03s       167865 ns/op   26736 B/op   1509 allocs/op
     beego:     0.06s       277612 ns/op   51725 B/op   3783 allocs/op
      gorm:     0.06s       308905 ns/op   31667 B/op   2062 allocs/op
 gorm_stmt:     0.08s       388122 ns/op   31602 B/op   2059 allocs/op
      aorm:     0.13s       668074 ns/op  135105 B/op   8744 allocs/op
      xorm:     doesn't work
```

- orm-benchmark -multi=5

```
  1000 times - Insert
  raw_stmt:     0.79s       786417 ns/op     288 B/op     10 allocs/op
 gorm_stmt:     0.80s       801062 ns/op    2230 B/op     42 allocs/op
       raw:     0.81s       806618 ns/op     511 B/op     12 allocs/op
     beego:     0.83s       834552 ns/op     906 B/op     36 allocs/op
      xorm:     0.84s       840072 ns/op    1674 B/op     48 allocs/op
      aorm:     0.85s       848826 ns/op    1727 B/op     72 allocs/op
      gorm:     0.88s       877997 ns/op    2452 B/op     45 allocs/op

  1000 times - MultiInsert 100 row
  raw_stmt:     1.47s      1472303 ns/op   65957 B/op    511 allocs/op
 gorm_stmt:     1.53s      1531803 ns/op   80829 B/op   1138 allocs/op
       raw:     1.69s      1687143 ns/op   88236 B/op    514 allocs/op
     beego:     1.75s      1745269 ns/op   91298 B/op   1435 allocs/op
      gorm:     1.81s      1813447 ns/op  101322 B/op   1140 allocs/op
      xorm:     1.94s      1941548 ns/op  155196 B/op   2850 allocs/op
      aorm:     doesn't work

  1000 times - Update
  raw_stmt:     0.03s        27099 ns/op     304 B/op     10 allocs/op
     beego:     0.03s        34336 ns/op     880 B/op     36 allocs/op
 gorm_stmt:     0.04s        38450 ns/op    2656 B/op     55 allocs/op
       raw:     0.05s        52460 ns/op     536 B/op     12 allocs/op
      xorm:     0.06s        56579 ns/op    1666 B/op     85 allocs/op
      gorm:     0.06s        62307 ns/op    2888 B/op     57 allocs/op
      aorm:     0.07s        68353 ns/op    1929 B/op     85 allocs/op

  1000 times - Read
  raw_stmt:     0.03s        27386 ns/op     836 B/op     32 allocs/op
     beego:     0.03s        31627 ns/op    1617 B/op     88 allocs/op
 gorm_stmt:     0.04s        40873 ns/op    2262 B/op     66 allocs/op
       raw:     0.05s        52692 ns/op     900 B/op     35 allocs/op
      gorm:     0.05s        54847 ns/op    2327 B/op     69 allocs/op
      aorm:     0.07s        72048 ns/op    2922 B/op    140 allocs/op
      xorm:     0.09s        93443 ns/op    5904 B/op    226 allocs/op

  1000 times - MultiRead limit 100
  raw_stmt:     0.17s       165182 ns/op   21152 B/op   1020 allocs/op
       raw:     0.18s       181178 ns/op   26736 B/op   1509 allocs/op
     beego:     0.30s       299378 ns/op   51734 B/op   3783 allocs/op
 gorm_stmt:     0.32s       321313 ns/op   31604 B/op   2059 allocs/op
      gorm:     0.37s       371398 ns/op   31662 B/op   2062 allocs/op
      aorm:     0.68s       676583 ns/op  135108 B/op   8744 allocs/op
      xorm:     doesn't work
```

- orm-benchmark -multi=20

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
