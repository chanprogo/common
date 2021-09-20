
`cd /var/lib/pgsql/10/data`  



`sudo vim postgresql.conf`   
```
listen_addresses = '*'
```   



`sudo vim pg_hba.conf` 客户机连接文件    
```
host    all             all             0.0.0.0/0               md5
```    