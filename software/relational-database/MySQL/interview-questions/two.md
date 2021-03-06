




## 说说 mysql 主从同步 怎么做的吧？

mysql 主从同步 的原理

1. master 提交完事务后，写入 binlog
2. slave 连接到 master，获取 binlog
3. master 创建 dump 线程，推送 binglog 到 slave
4. slave 启动一个 IO 线程    读取  同步过来的 master 的 binlog，记录到 relay log 中继 日志中
5. slave 再开启一个 sql 线程 读取 relay log 事件   并在 slave 执行，完成同步
6. slave 记录 自己的 binglog



由于 mysql 默认的 复制方式 是异步的，主库 把日志发送给 从库后 不关心从库是否 已经处理，
这样会产生一个问题 就是假设 主库挂了，从库 处理失败了，这时候从库升为 主库后，日志 就丢失了。
由此产生两个概念。  

全同步复制：  
主库写入 binlog 后 强制 同步 日志 到从库，所有的从库 都执行完成后 才返回给客户端，
但是 很显然这个方式的话   性能会受到严重影响。

半同步复制： 
和 全同步不同的是，半同步复制的逻辑 是这样，
从库 写入日志 成功后 返回 ACK 确认
给 主库，主库 收到 至少一个从库的 确认 就认为 写操作完成。  

