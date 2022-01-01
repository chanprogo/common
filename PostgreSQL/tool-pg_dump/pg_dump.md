

`cp /usr/bin/pg_dump /usr/bin/pg_dump_bak`  
`cp pg_dump /usr/bin/pg_dump`   


`pg_dump -h 10.190.24.252 -U admin creative_project> creative_project_bak_20211231.sql`
`psql -h 10.190.24.252 -U admin creative_project < creative_project_20211231.sql`   
 