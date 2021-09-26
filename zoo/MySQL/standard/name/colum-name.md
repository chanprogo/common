

1. 数据库 字段命名 与 表名 命名 类似： 使用小写英文单词，如果有多个单词使用下划线隔开； 使用简单单词，避免生僻词；
2. 禁止使用数据库关键字，如：name，time ，datetime password 等。
3. 字段名称 一般采用名词 或 动宾短语。  
4. 是别的表 的外键 均使用 xxx_id 的方式 来表明；  表名_id  
5. 表的主键 一般 都约定成为 id，自增类型；





三、【常用列名约定】  
2. cid – 特殊的编号，带有元数据，方便 关联查询，你可以把它 理解成 类别(层次)编号。  
举个例子，产品在分类时，往往需要 将其归类到 子分类下，相应的字段中 也一般 只记录 子分类的 id，  
这时若需要知道  该产品 属于哪个 主分类，就需要通过子分类信息 再查询到主分类信息，这是比较麻烦的，cid 字段就是要解决这个问题。一般的站点几十个分类肯定是够用了，  
所以这里假设 某一主分类的 cid 为 11，则子分类的 cid 从 1101 开始编号，处理时 只需截取前两位数值 便可知道该产品 属于哪一个主分类了。  

3. add_time – 添加时间、上架时间等
4. last_time – 最后操作时间，如登录、修改记录
5. expire_time – 过期时间
6. name – 商品名称、商家名称等，不要跟title混用，title只用于文章标题、职称等
7. price – 价格
8. thumb – 只要是列表页面中的窗口图，一律用此命名
9. image_src – 相册中的图片地址一律用此命名，不要出现各种img,image,img_url,thumb_url等
10. head_thumb – 用户头像， 虽然有点长，一定要遵守。不要出现上述情况
11. image_alt – 相册中图片的alt属性
12. desc – 描述、简介，比如goods_desc，不要出现goods_txt这种
13. details – 详情、文章内容等
14. order_id – 排序
15. telephone – 座机号码
16. mobile – 手机号码
17. phone – 当不区分手机和座机时，请用phone命名
18. address – 地址，单独出现不要用addr缩写，组合出现时需用缩写，比如mac地址，mac_addr
19. zipcode – 邮编
20. region – 地区，大的区域，比如记录杭州市、温州市等
21. area – 区域，小的，比如上城区，江干区等
22. avg_cost – 人均消费
23. 待续