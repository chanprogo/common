
一、【操作规范】
1.  
2. 如无备注，则数值类型 的字段 请使用 UNSIGNED 属性；




3. 如无备注，排序字段 order_id 在程序中 默认 使用 降序排列；

4. 如无备注，所有字段 都设置 NOT NULL，并设置 默认值；


5. 如无备注，所有的布尔值字段，如 is_hot、is_deleted，都必须 设置 一个默认值，并设为 0；


6. 所有的数字类型 字段，都必须设置一个默认值，并设为 0；



7. 针对 varchar 类型字段的 程序处理，请验证 用户输入，不要超出 其预设的长度；  

8. 建表时 将数据字典中 的 字段 中文名 和属性备注 写入数据表的备注中(“PK、自动增长”不用写)；








### 数据库命名规范

````
1.所有的数据库对象名称(包括库名、表名、列名等等)必须以小写字母命名，每个单词之间用下划线分割
3.数据库对象的命名要能做到见名知意，并且不要超过32个字符
6.在不同的库或表中，要保证所有存储相同数据的列名 和列类型 必须一致
````

### 数据库基本设计规范

````
1.数据库和表的字符集统一使用UTF8字符集，避免由于字符集的转换产生乱码。
2.所有的表和字段都需要添加注释。使用comment从句添加表和列的备注，从一开始就把数据字典维护好

4.可以采用历史数据归档(常见于日志表)和分库分表的方式控制单表数据的大小。
5.谨慎使用MySQL分区表，避免跨分区查询，否则效率很低。
6.分区表在逻辑上表现为一个表，但是在物理上将数据存储在多个文件。最好能将分区表的不同分区文件存储在不同的磁盘阵列上。
7.表中的列不要太多，尽量做到冷热数据分离，减小表的宽度。
8.减少表的宽度，可以让一页内存中容纳更多的行，进而减少磁盘IO，更有效的利用缓存。
9.经常一起使用的列尽量放到一个表中，避免过多的关联操作。
10.修改列的类型会所锁表。修改一个字段类型的成本要高于增加一个字段。
11.禁止在数据库中存储图片、文件等
````



### 数据库SQL开发规范

````
1.在程序中，建议使用预编译语句进行数据库操作。预编译只编译一次，多次使用，比SQL效率高。
2.避免数据类型的隐式转换。隐式转换会导致索引失效。
3.避免使用双%25号或前置%25号的查询条件，这样无法利用到索引。
4.不提倡在查询中使用select *
5.避免使用子查询，子查询会产生临时表，临时表没有任何索引，数据量大时严重影响效率。建议把子查询转化成关联查询。
(select * from user where uid in (select uid from company_users);)
6.避免使用JOIN关联太多的表。
7.尽量减少同数据库的交互次数，数据库更适合处理批量操作。批量insert的条目最好不要超过1000。
8.使用IN代替OR
9.禁止在where从句中对列进行函数转换和计算，会导致索引无效。（select age from user where age+1=10???）
10.在不需要去重的情况下，要使用UNION
11.ALL代替UNION。
````

