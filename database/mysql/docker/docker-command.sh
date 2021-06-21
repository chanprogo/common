#!/bin/bash

docker run -d --name calldb --net=bridge -p 3306:3306 -v /root/devware/database/mysql/docker/conf/my.cnf:/etc/my.cnf -v /root/devware/database/mysql/docker/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=toor -e MYSQL_USER=admin -e MYSQL_PASS=admin mysql/mysql-server