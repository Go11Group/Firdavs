create table department
(
    id   uuid primary key default gen_random_uuid() not null,
    name varchar                                    not null
);


create table if not exists employee
(
    id            uuid primary key default gen_random_uuid() not null,
    name          varchar not null,
    department_id uuid references department (id),
    salary        int     not null default 0
);


create table if not exists task
(
    id          uuid primary key default gen_random_uuid() not null,
    number      int unique                                 not null,
    description text             default ''
);

create table if not exists employee_task
(
    id          uuid primary key default gen_random_uuid() not null,
    employee_id uuid references employee (id),
    task_id     uuid references task (id)
);

insert into department(id, name) values('5939ed2c-4b4b-43ae-a0ab-695175c3a731','HR');
insert into department(id, name) values('558f6658-7793-4107-8930-cbf5dbd2b008','Programming');
insert into department(id, name) values('08894581-f43e-4713-bccd-4283b53fd08b','Marketing');

INSERT INTO employee (name, department_id, salary) VALUES
   ('John Doe', '5939ed2c-4b4b-43ae-a0ab-695175c3a731', 50000),
   ('Alice Smith', '558f6658-7793-4107-8930-cbf5dbd2b008', 60000),
   ('Bob Johnson', '558f6658-7793-4107-8930-cbf5dbd2b008', 55000),
   ('Emily Davis', '08894581-f43e-4713-bccd-4283b53fd08b', 52000),
   ('Michael Brown', '08894581-f43e-4713-bccd-4283b53fd08b', 62000),
   ('Jennifer Lee', '08894581-f43e-4713-bccd-4283b53fd08b', 58000),
   ('David Wilson', '5939ed2c-4b4b-43ae-a0ab-695175c3a731', 54000),
   ('Jessica Martinez', '5939ed2c-4b4b-43ae-a0ab-695175c3a731', 67000);

INSERT INTO task (id, number, description) VALUES
   ('1d7ab001-df4e-4ac6-adf2-70cab7568c42', 101, 'Task 101 description'),
   ('fbb3cd9f-0567-4496-b694-96de4aad5c8d', 102, 'Task 102 description'),
   ('488fb943-37c3-47bb-b3af-1cdd5b9dbfbb', 103, 'Task 103 description');

INSERT INTO employee_task (employee_id, task_id) VALUES
     ('0a876d4f-937e-4f77-9383-825f260728b2','1d7ab001-df4e-4ac6-adf2-70cab7568c42'),
     ('45c2aabd-143d-40a5-b053-2142ab6521f9','1d7ab001-df4e-4ac6-adf2-70cab7568c42'),
     ('45c2aabd-143d-40a5-b053-2142ab6521f9','fbb3cd9f-0567-4496-b694-96de4aad5c8d'),
     ('0a876d4f-937e-4f77-9383-825f260728b2','488fb943-37c3-47bb-b3af-1cdd5b9dbfbb'),
     ('45c2aabd-143d-40a5-b053-2142ab6521f9','488fb943-37c3-47bb-b3af-1cdd5b9dbfbb');

select array_agg(e.name), t.description from employee e
 left join department d on e.department_id=d.id
 left join employee_task et on e.id = et.employee_id
 left join task t on t.id = et.task_id
where t.description is not null
group by t.description;

select array_agg(e.name), d.name, avg(salary) sum from employee e
    left join department d on d.id = e.department_id
    group by d.id
    order by sum
;

-- Student table
-- course table
-- student_course table
-- grade table

-- 1. grade tabledagi student_id va course_id o'rniga student_course tabledagi
-- idni olishimiz darkor;

-- 2. guruhdagi eng yaxshi o'qiydigan studentlarni har guruh bo'yicha chiqaring. Agarda
-- eng yaxshi baholar bir nechta kishida bo'lsa, hammasi chiqsin.

-- 3. guruhning o'rtacha bahosini har bir guruh bo'yicha chiqaring;

-- 4. eng yosh o'quvchi guruh bo'yicha chiqarilsin chiqarilsin;
-- yani gar bir guruhdan eng yosh o'quvchi(bir nechta bo'lsa har birini)

-- 5. eng yaxshi o'qiydigan guruh chiqarilsin
-- ya'ni, har bir guruhning o'rtacha bahosining eng katta(yaxshi) bo'lgani