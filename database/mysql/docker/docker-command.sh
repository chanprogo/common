#!/bin/bash

docker run -d --name calldb --net=bridge -p 3306:3306 -v /deploy/ccusercenter/mysql-conf/my.cnf:/etc/my.cnf -v /deploy/ccusercenter/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=rootroot -e MYSQL_USER=admin -e MYSQL_PASS=admin mysql/mysql-server