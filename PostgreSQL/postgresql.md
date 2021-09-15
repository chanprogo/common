
资产表 逐笔 更新记录 
更新前后资产

per_change_finacial_record   (root)                  before, later 为 json 格式数据








```
select per_change_finacial_record.before::json->>'use_amount' as before_use_amount,
       per_change_finacial_record.later::json->>'use_amount' as later_use_amount 
from per_change_finacial_record limit 1
```    


查询 一条 更新 操作 前后     用户的 可用资产
