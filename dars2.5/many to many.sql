CREATE TABLE cars (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    brend VARCHAR NOT NULL,
    model VARCHAR,
    year INT,
    price INT
);


INSERT INTO cars (brend, model, year, price) VALUES
('Toyota', 'Camry', 2020, 24000),
('BMW', 'X5', 2019, 45000),
('Mercedes-Benz', 'C-Class', 2021, 41000);

CREATE TABLE user_car (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    age INT NOT NULL,
    gender CHAR(1) NOT NULL
);

INSERT INTO user_car (name, age, gender ) VALUES
('John Doe', 30, 'M'),
('Jane Smith', 25, 'F'),
('Alice Johnson', 22, 'F'),
('Bob Brown', 28, 'M'),
('Charlie Davis', 35, 'M');




CREATE TABLE user_car_relationship (
    user_id UUID NOT NULL,
    car_id UUID NOT NULL,
    PRIMARY KEY (user_id, car_id),
    FOREIGN KEY (user_id) REFERENCES user_car(id),
    FOREIGN KEY (car_id) REFERENCES cars(id)
);



INSERT INTO user_car_relationship (user_id, car_id) VALUES
('ec457f8d-8c83-46bd-9e5b-9db93bd202bf', 'bfcdc260-acc9-45c4-84e9-9a5fa91cba18'), -- John Doe -> Toyota
('b23847a2-abea-4df1-9ecd-5f9ac98ec3c8', '072a89d8-0760-4e32-b5ac-96bc164530a9'), -- Jane Smith -> Mercedes-Benz
('b62994b4-1b42-4f6d-90bd-4a3e853738c3', 'bfcdc260-acc9-45c4-84e9-9a5fa91cba18'), -- Alice Johnson -> Toyota
('c8318d83-605f-4524-aa99-ae3278296572', '072a89d8-0760-4e32-b5ac-96bc164530a9'), -- Bob Brown -> Mercedes-Benz
('59153310-559e-44ad-879b-31d7d0ff3edc', 'd4f48a7d-44d5-47a5-98ac-e907b156d380'); -- Charlie Davis -> BMW


SELECT 
    u.id AS user_id,
    u.name,
    u.age,
    u.gender,
    c.id AS car_id,
    c.brend,
    c.model,
    c.year,
    c.price
FROM 
    user_car u
JOIN 
    user_car_relationship uc ON u.id = uc.user_id
JOIN 
    cars c ON uc.car_id = c.id;


SELECT
    u.name,
    u.age,
    u.gender,
    c.brend,
    c.model,
    c.year,
    c.price
FROM 
    user_car u
JOIN 
    user_car_relationship uc ON u.id = uc.user_id
JOIN 
    cars c ON uc.car_id = c.id;

