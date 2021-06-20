
使用客户端链接到 zookeeper： `sh zkCli.sh`

`/usr/local/services/zookeeper/zookeeper-3.4.9/bin/zkCli.sh -server 127.0.0.1:2181`  

`quit`  


创建配置节点（这个后续会提供页面来处理）：  

`create /demo demo`  
`create /demo/confs confs`  
`create /demo/confs/conf1 111111111111111111111`  
`create /demo/confs/conf2 222222222222222222222`  
`create /demo/confs/conf3 333333333333333333333`  







`ls` 查看 当前节点数据  
`ls2`  查看 当前节点数据 并能 看到 更新次数 等 数据  
`create` 创建 一个节点  
`get` 得到 一个节点，包含 数据 和 更新次数 等 数据  
`set` 修改节点  
`delete` 删除一个节点  