
增：   
```sql
ALTER TABLE theme ADD COLUMN current_workflow VARCHAR ( 20 ) NOT NULL DEFAULT '';

COMMENT ON COLUMN theme.current_workflow IS '流程状态码';  
```






改：   
```sql
ALTER TABLE project ALTER COLUMN product_available 
SET DATA TYPE NUMERIC USING product_available :: "numeric" ( 32, 2 );
```

```sql
ALTER TABLE project ALTER COLUMN product_available TYPE NUMERIC ( 32, 2 );
```


删：  
```sql
ALTER TABLE products DROP COLUMN description;
```
