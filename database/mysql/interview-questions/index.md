
假设我们有张表，结构如下：  

```
create table user(
  id int(11) not null,
  age int(11) not null,
  primary key(id),
  key(age)
);

```




## 聚簇 和 非聚簇 索引


索引 按照 数据结构 来说 主要包含 B+ 树 和 Hash 索引。  

B+ 树是 左小右大 的 顺序 存储结构， 节点 只包含 id 索引列， 而 叶子节点 包含 索引列 和 数据， 这种 数据和索引 在一起存储的 索引方式 叫做 聚簇索引，  

一张表 只能有 一个 聚簇索引。假设 没有 定义主键，InnoDB 会选择 一个 唯一的 非空索引 代替，如果没有的话 则会隐式定义 一个主键 作为 聚簇索引。  



非聚簇索引(二级索引) ：    保存的是 主键 id 值，  这一点和 myisam 保存的是 数据地址 是不同的。  









## 什么是 覆盖索引 和 回表

覆盖索引 指的是 在一次查询中，如果一个索引  包含 或者说 覆盖  所有 需要查询的字段 的值，我们就称之为 覆盖索引，而不再需要 回表查询。

而要确定一个查询  是否是  覆盖索引，我们只需要 explain sql 语句看 Extra 的结果是否是 “Using index” 即可。

以上面的 user 表来举例，我们再增加一个 name 字段，然后做一些查询试试。
`explain select * from user where age=1;` // 查询的 name 无法从索引数据获取 
`explain select id,age from user where age=1;` // 可以直接从索引获取
      





### 数据库索引设计规范

每个 Innodb 表 都必须有 一个主键，且不使用 更新频繁的列 作为主键，不使用 多列主键。
不使用 UUID、MD5、字符串列 作为主键。
最好选择 值的顺序 是连续增长的列 作为主键，所以建议选择使用自增 ID 列作为主键  

5.建议在下面的列上建立索引：
- 在SELECT，UPDATE，DELETE语句的WHERE从句上的列。
- 在ORDERBY，GROUP BY，DISTINCT上的列。多表 JOIN 的 关联列。

- 区分度 最高 的列 放在联合索引的最左侧。使用频繁的列 放在联合索引 的最左侧。

- 避免冗余的索引，如：primary key(id)，index(id)，unique index(id)

