create table product
(
    id       uuid primary key,
    name     varchar,
    category varchar,
    cost     int
);

SELECT
    tablename,
    indexname,
    indexdef
FROM
    pg_indexes
WHERE
    schemaname = 'public'
ORDER BY
    tablename,
    indexname;
-- single index
create index product_id_idx on product (id);

-- multi index
create index product_id_idx_cat on product(name, category);

drop index product_id_idx;

create unique index product_id_idx on product (id, cost);

explain (analyse )
select * from product where name  = 'Luis' and category='冯' -- id = 'b2457480-75c7-481b-8b76-c40795ec8ff0';

