# Results

- orm-benchmark (with no flags)

```
   200 times - Insert
  raw_stmt:     0.15s       770719 ns/op     288 B/op     10 allocs/op
      aorm:     0.16s       797375 ns/op    1756 B/op     72 allocs/op
      gorm:     0.16s       804619 ns/op    2662 B/op     45 allocs/op
     beego:     0.16s       805161 ns/op     915 B/op     36 allocs/op
      xorm:     0.17s       828912 ns/op    1714 B/op     48 allocs/op
 gorm_stmt:     0.17s       849400 ns/op    2419 B/op     43 allocs/op
       raw:     0.17s       862548 ns/op     542 B/op     12 allocs/op

   200 times - MultiInsert 100 row
  raw_stmt:     0.32s      1579945 ns/op   66330 B/op    514 allocs/op
 gorm_stmt:     0.32s      1585980 ns/op   80823 B/op   1137 allocs/op
     beego:     0.34s      1695663 ns/op   91261 B/op   1435 allocs/op
      gorm:     0.36s      1824202 ns/op  101319 B/op   1139 allocs/op
       raw:     0.37s      1858088 ns/op   88610 B/op    517 allocs/op
      xorm:     0.38s      1899470 ns/op  155194 B/op   2850 allocs/op
      aorm:     doesn't work

   200 times - Update
     beego:     0.01s        34322 ns/op     887 B/op     36 allocs/op
  raw_stmt:     0.01s        35044 ns/op     304 B/op     10 allocs/op
 gorm_stmt:     0.01s        48088 ns/op    2657 B/op     55 allocs/op
       raw:     0.01s        52153 ns/op     536 B/op     12 allocs/op
      aorm:     0.01s        67453 ns/op    1936 B/op     85 allocs/op
      gorm:     0.01s        69032 ns/op    2888 B/op     57 allocs/op
      xorm:     0.02s        75831 ns/op    1674 B/op     85 allocs/op

   200 times - Read
     beego:     0.01s        31911 ns/op    1621 B/op     88 allocs/op
  raw_stmt:     0.01s        34350 ns/op     852 B/op     32 allocs/op
 gorm_stmt:     0.01s        35708 ns/op    2289 B/op     66 allocs/op
       raw:     0.01s        54604 ns/op     916 B/op     35 allocs/op
      aorm:     0.01s        60493 ns/op    1936 B/op     86 allocs/op
      gorm:     0.01s        62742 ns/op    2359 B/op     69 allocs/op
      xorm:     0.02s        86120 ns/op    5925 B/op    226 allocs/op

   200 times - MultiRead limit 100
  raw_stmt:     0.03s       166087 ns/op   21152 B/op   1020 allocs/op
       raw:     0.04s       175572 ns/op   26736 B/op   1509 allocs/op
      aorm:     0.06s       283215 ns/op   61708 B/op   3057 allocs/op
 gorm_stmt:     0.06s       288360 ns/op   31602 B/op   2059 allocs/op
     beego:     0.06s       295255 ns/op   51725 B/op   3783 allocs/op
      gorm:     0.06s       302646 ns/op   31667 B/op   2062 allocs/op
      xorm:     doesn't work
```

- orm-benchmark -multi=5

```
  1000 times - Insert
  raw_stmt:     0.77s       770909 ns/op     288 B/op     10 allocs/op
 gorm_stmt:     0.78s       777428 ns/op    2240 B/op     43 allocs/op
     beego:     0.79s       787603 ns/op     906 B/op     36 allocs/op
       raw:     0.79s       791816 ns/op     511 B/op     12 allocs/op
      gorm:     0.81s       810114 ns/op    2441 B/op     44 allocs/op
      xorm:     0.82s       819090 ns/op    1674 B/op     48 allocs/op
      aorm:     0.83s       834877 ns/op    1727 B/op     72 allocs/op

  1000 times - MultiInsert 100 row
  raw_stmt:     1.45s      1454755 ns/op   65957 B/op    511 allocs/op
 gorm_stmt:     1.53s      1525804 ns/op   80829 B/op   1138 allocs/op
     beego:     1.63s      1628128 ns/op   91331 B/op   1435 allocs/op
       raw:     1.67s      1667932 ns/op   88235 B/op    514 allocs/op
      gorm:     1.79s      1786172 ns/op  101317 B/op   1140 allocs/op
      xorm:     1.80s      1797460 ns/op  155201 B/op   2850 allocs/op
      aorm:     doesn't work

  1000 times - Update
  raw_stmt:     0.03s        27949 ns/op     304 B/op     10 allocs/op
     beego:     0.03s        29683 ns/op     880 B/op     36 allocs/op
 gorm_stmt:     0.04s        36137 ns/op    2656 B/op     55 allocs/op
       raw:     0.05s        52740 ns/op     536 B/op     12 allocs/op
      xorm:     0.06s        58062 ns/op    1666 B/op     85 allocs/op
      gorm:     0.06s        58711 ns/op    2888 B/op     57 allocs/op
      aorm:     0.06s        61989 ns/op    1929 B/op     85 allocs/op

  1000 times - Read
  raw_stmt:     0.03s        28948 ns/op     836 B/op     32 allocs/op
     beego:     0.03s        30875 ns/op    1617 B/op     88 allocs/op
 gorm_stmt:     0.03s        31367 ns/op    2264 B/op     66 allocs/op
       raw:     0.05s        51320 ns/op     900 B/op     35 allocs/op
      gorm:     0.06s        55667 ns/op    2326 B/op     69 allocs/op
      aorm:     0.06s        56921 ns/op    1929 B/op     86 allocs/op
      xorm:     0.08s        84816 ns/op    5904 B/op    226 allocs/op

  1000 times - MultiRead limit 100
  raw_stmt:     0.18s       175270 ns/op   21152 B/op   1020 allocs/op
       raw:     0.19s       187905 ns/op   26736 B/op   1509 allocs/op
 gorm_stmt:     0.29s       294646 ns/op   31598 B/op   2059 allocs/op
     beego:     0.31s       307227 ns/op   51739 B/op   3783 allocs/op
      aorm:     0.32s       315347 ns/op   61704 B/op   3057 allocs/op
      gorm:     0.33s       334611 ns/op   31667 B/op   2062 allocs/op
      xorm:     doesn't work
```

- orm-benchmark -multi=20

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
