BEGIN;
CREATE TABLE  phone(
    id INT PRIMARY KEY NOT NULL,
    brend VARCHAR(50),
    model VARCHAR(50),
);

INSERT INTO phone(id, brend, model) VALUES (1, 'iphone', '12ProMax'),(2, 'samsung', 'galaxy not10'),(3, 'redmi', 'not8');

UPDATE phone SET model = 'Iphone13ProMax' WHERE id=1;
UPDATE phone SET model = 'S24Ultra' WHERE id=2;
UPDATE phone SET model = '13c' WHERE id=3;
COMMIT;