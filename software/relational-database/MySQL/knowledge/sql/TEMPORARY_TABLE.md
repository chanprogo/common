```sql
CREATE TEMPORARY TABLE `collect_task_tmp` (
	`id` INT (11) NOT NULL,
	`parent_id` INT (11) NOT NULL,
	`start_time` datetime DEFAULT NULL,
	`end_time` datetime DEFAULT NULL
);
```