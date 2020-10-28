# dbf-statistics-column


# example 
> dbf-statistics-column -f path/example.dbf -k name -l 3
>
> 统计在 path/example.pdf 所有列中 name 出现的频次，top 3。
>


## output
``` shell
  -f string
    	需要统计的dbf 文件信息 (default "example.dbf")
  -k string
    	按照该字段统计的列名称
  -l int
    	展示的行数 (default 10)

   0.  8      -> 735       
   1.  9      -> 720       
   2.  10     -> 708       
```


