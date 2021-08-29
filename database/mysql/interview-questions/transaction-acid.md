


假设我们有张表，结构如下：  
```
create table user(
  id int(11) not null,
  age int(11) not null,
  primary key(id),
  key(age)
);

```



## 隔离级别

隔离性有 4 个隔离级别：

read uncommit 读未提交，可能会 读到其他事务 未提交的数据，也叫做脏读。   

read commit 读已提交，两次读取结果 不一致，叫做 不可重复读。不可重复读 解决了 脏读的问题，他只会读取 已经提交 的事务。  

repeatable read 可重复复读，这是 mysql 的默认级别，就是 每次 读取结果 都一样，但是 有可能产生 幻读。  

serializable 串行，一般是不会使用的，他会给每一行 读取的数据 加锁，会导致 大量超时 和 锁竞争 的问题。  







## 幻读，MVCC  

MVCC 叫做 多版本 并发控制，实际上 就是 保存了 数据 在某个时间节点 的快照。

我们每行数 实际上隐藏了两列，   创建时间 版本号， 过期(删除)时间 版本号， 每开始一个 新的事务，版本号 都会自动递增。

还是拿上面的 user 表 举例子，假设我们插入 两条数据，他们实际上 应该长这样。
id   name    create_version   delete_version

这时候假设小明去执行查询，此时 current_version=3  `select * from user where id<=3; `  

同时，小红 在这时候 开启事务 去修改 id=1 的记录， current_version=4  `update user set name=‘张三三’ where id=1;`  

如果这时候 还有小黑在删除 id=2 的数据，current_version=5
  

由于 MVCC 的原理是 查找   
创建版本 小于或等于 当前事务版本，
删除版本为空 或者大于 当前事务版本，

小明的真实的查询应该是这样   
`select * from user where id<=3 and create_version<=3          and          (delete_version>3 or delete_version is null);`  
所以小明最后查询到的 id=1 的名字还是’张三’，并且 id=2 的记录也能查询到。

这样做是为了保证事务 读取的数据 是在事务开始前 就已经存在的，要么是事务 自己插入 或者 修改的。





明白 MVCC 原理，我们来说什么是幻读 就简单多了。
举一个常见的场景，用户注册时，我们先查询用户名是否存在，不存在就插入，假定用户名是唯一索引。   

小明开启事务 current_version=6 查询名字为’王五’的记录，发现不存在。   
小红开启事务 current_version=7 插入一条数据  

小明执行插入名字’王五’的记录，发现唯一索引冲突，无法插入，这就是幻读。    









## 间隙锁

间隙锁 是 可重复读 级别下 才会有的锁，
结合 MVCC 和 间隙锁 可以解决幻读的问题。


当我们执行：

`select * from user where age=20 for update;` 

`insert into user(age) values(10);` #成功 

`insert into user(age) values(11);` #失败 
`insert into user(age) values(20);` #失败 
`insert into user(age) values(21);` #失败 
`insert into user(age) values(30);` #失败

只有 10 可以插入成功，那么因为表的间隙 mysql 自动帮我们生成了区间(左开右闭)

(negative infinity，10],(10,20],(20,30],(30,positive infinity)

由于 20 存在记录，所以(10,20]，(20,30] 区间 都被锁定了 无法插入、删除。

如果查询 21 呢？就会根据 21 定位到(20,30)的区间(都是开区间)。

需要注意的是唯一索引 是不会 有间隙索引的。  