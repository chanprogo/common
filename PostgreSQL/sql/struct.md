

```sql
ALTER TABLE project ALTER COLUMN product_available 
SET DATA TYPE NUMERIC USING product_available :: "numeric" ( 32, 2 );
```

```sql
ALTER TABLE project ALTER COLUMN product_available TYPE NUMERIC ( 32, 2 );
```
