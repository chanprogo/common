

数据库的版本号：   

PostgreSQL 10.4 on x86_64-pc-linux-gnu, compiled by gcc (GCC) 4.8.5 20150623 (Red Hat 4.8.5-39), 64-bit  


系统： CentOS release 6.8 (Final)  

`cat /etc/redhat-release`  
`cat /etc/issue`  
`uname -r`  












```
sudo yum install -y https://download.postgresql.org/pub/repos/yum/reporpms/EL-6-x86_64/pgdg-redhat-repo-latest.noarch.rpm
sudo yum install -y postgresql10-server
sudo service postgresql-10 initdb
sudo chkconfig postgresql-10 on
sudo service postgresql-10 start
```
