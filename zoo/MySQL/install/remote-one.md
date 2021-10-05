
`use mysql`   
`select user,host from user;`  
  

`create user 'root'@'%' identified by 'yourpassword';`   
​
`grant all on *.* to 'root'@'%' with grant option;`   
 
`flush privileges;`  


`select user,host from user;`  
`exit`  




​




`update user set authentication_string=password('yourpassword'), plugin='mysql_native_password' where user="root";`  或者  
`alter user ‘用户名’@'主机名' identified with mysql_native_password by 'yourpassword';`    
alter 对密码有安全规范要求，不能设置简单密码（大小写，数字，字符，无序）。update 不做限制。   


`show grants;`  ---查看当前用户的权限  
`show grants for 'abc'@'localhost';`  ---查看用户abc的权限  

`revoke all privileges on *.* from 'abc'@localhost';`   --回收用户abc的所有权限  
`revoke grant option on *.* from 'abc'@localhost';`   --回收权限的传递  
