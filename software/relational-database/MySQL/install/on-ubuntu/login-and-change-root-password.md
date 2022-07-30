
`sudo mysql`  

修改 root 本地登录密码  
`ALTER USER 'root'@'localhost' IDENTIFIED BY 'yourpassword';`   
`ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'yourpassword';`  // new  
