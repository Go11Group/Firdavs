CREATE TABLE employees (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  age INTEGER NOT NULL,
  department VARCHAR(255) NOT NULL
);



INSERT INTO employees (name, age, department) VALUES
('John Doe', 30, 'IT'),
('Jane Smith', 25, 'Sales'),
('Michael Jones', 35, 'Marketing'),
('Sarah Miller', 28, 'IT'),
('Robert Brown', 40, 'Finance'),
('Jon Digle', 22, 'IT'),
('Mchle Brown', 27, 'Sales'),
('Robert Miller', 23, 'IT'),
('Anglena Smith', 28, 'Marketing'),
('Sarah Doe', 31, 'Finance');