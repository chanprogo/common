
默认情况下，Postgres 使用称为“角色”的概念 来处理 身份验证和授权。  
这些 在某些方面 类似于 普通的 Unix 风格 的账户，但是 Postgres 并没有 区分用户和组，而是倾向于 更灵活的术语“角色”。

安装后，Postgres 被设置为 使用 ident 身份验证，这意味着 它将 Postgres 角色 
与匹配的 Unix / Linux 系统帐户 相关联。  
如果 Postgres 中存在一个角色，则具有相同名称的 Unix / Linux 用户名 可以 作为 该角色登录。

安装过程 创建了一个名为 postgres 的用户帐户， 它与默认的 Postgres 角色 相关联。  
为了使用 Postgres， 您可以登录到 该帐户。



`sudo -i -u postgres`  +  `psql`    ==    `sudo -u postgres psql`  



`sudo -u postgres createuser --interactive`  
`sudo -u postgres createdb sammy`  


`sudo adduser sammy`   



`\conninfo`   
`\q`  




查看配置文件路径  
`select name,setting from pg_settings where category='File Locations';`  
