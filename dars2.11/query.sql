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




--new


create  index product_id_3 on product (category, name, cost);

DROP INDEX product_id_3;

create unique index product_id_2 on product ( name, category, cost);

create unique index product_id_1 on product (cost, category, name);

create unique index product_2 on product (address, phoneNumber);

create unique index product_2 on product (password, name);

create index persons_indx on product using hash (id);


explain (analyse )
select * from product where category = 'Rippin' and name  = 'Allen' ;

explain (analyse )
select * from product where name  = 'Dante ' and cost = 4234 ;

explain (analyse )
select * from product where id = 'b4beab58-19f5-4663-911b-8c1fcc07b9c1' 

