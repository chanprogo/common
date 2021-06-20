
使用客户端链接到 zookeeper： `sh zkCli.sh`

`/usr/local/services/zookeeper/zookeeper-3.4.9/bin/zkCli.sh -server 127.0.0.1:2181`  




创建配置节点（这个后续会提供页面来处理）：  

`create /demo demo`  
`create /demo/confs confs`  
`create /demo/confs/conf1 111111111111111111111`  
`create /demo/confs/conf2 222222222222222222222`  
`create /demo/confs/conf3 333333333333333333333`  
