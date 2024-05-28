CREATE TABLE IF NOT EXISTS Student(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid() not null,
    name VARCHAR NOT NULL,
    age int 
)