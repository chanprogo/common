



## HTTPS 是 如何实现 安全传输数据的  

HTTPS 其实就是在 HTTP 跟 TCP 中间 加多了 一层加密层 TLS/SSL。  

SSL 是个加密套件，负责对 HTTP 的数据进行加密。TLS 是 SSL 的升级版。  

现在提到 HTTPS，加密套件 基本指的是 TLS。

原先是应用层 将数据直接给到 TCP 进行传输，现在改成应用层 将数据 给到 TLS/SSL，将数据加密后，再给到 TCP 进行传输。     







## HTTPS 传输过程?

客户端发起 HTTPS 请求,服务端 返回证书,客户端 对证书验证,验证通过后 本地生成 用于改造对称加密算法 的随机数，
通过证书中的公钥 对随机数 进行加密传输到服务端，服务端 接收后 通过私钥 解密得到 随机数，之后的 数据交互 通过对称加密算法 进行加解密。

 

## 为什么需要证书?
防止中间人攻击,验证服务器身份

　　　

## 怎么防止的 篡改?
证书是公开的，虽然中间人可以拿到证书，但私钥无法获取，公钥无法推断出私钥，
所以篡改后 不能用私钥加密，强行加密 客户也无法解密，强行修改内容，  会导致证书内容 与签名中的指纹不匹配  