CREATE TABLE users (
    user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    birthday TIMESTAMP NOT NULL,
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

INSERT INTO users ( name, email, birthday, password) 
VALUES ( 'Firdavs', 'firdavs@example.com', '2002-07-15', 'yourpassword');



CREATE TABLE courses (
    course_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE lessons (
    lesson_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    course_id UUID NOT NULL,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0,
    FOREIGN KEY (course_id) REFERENCES courses(course_id)
);

CREATE TABLE enrollments (
    enrollment_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    course_id UUID NOT NULL,
    enrollment_date TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (course_id) REFERENCES courses(course_id)
);


-- Foydalanuvchilar jadvaliga ma'lumot kiritish
INSERT INTO users (name, email, birthday, password) 
VALUES 
('Firdavs', 'firdavs@example.com', '2002-07-15', 'yourpassword'),
('Ali', 'ali@example.com', '1995-03-22', 'password123'),
('Zaynab', 'zaynab@example.com', '1998-11-05', 'securepass'),
('Bobur', 'bobur@example.com', '1990-06-18', 'mypassword');

-- Kurslar jadvaliga ma'lumot kiritish
INSERT INTO courses (title, description) 
VALUES 
('Python Programming', 'Learn the basics of Python programming.'),
('Web Development', 'Build modern web applications with HTML, CSS, and JavaScript.'),
('Data Science', 'Introduction to data science with Python and R.'),
('Machine Learning', 'Fundamentals of machine learning and AI.');

-- Darslar jadvaliga ma'lumot kiritish
INSERT INTO lessons (course_id, title, content) 
VALUES 
((SELECT course_id FROM courses WHERE title = 'Python Programming'), 'Introduction to Python', 'Welcome to Python Programming. This lesson will cover the basics of Python.'),
((SELECT course_id FROM courses WHERE title = 'Python Programming'), 'Data Types and Variables', 'In this lesson, we will discuss data types and variables in Python.'),
((SELECT course_id FROM courses WHERE title = 'Web Development'), 'HTML Basics', 'This lesson introduces the basics of HTML.'),
((SELECT course_id FROM courses WHERE title = 'Web Development'), 'CSS Basics', 'This lesson covers the basics of CSS for styling web pages.'),
((SELECT course_id FROM courses WHERE title = 'Data Science'), 'Data Analysis with Python', 'Learn how to perform data analysis using Python.'),
((SELECT course_id FROM courses WHERE title = 'Data Science'), 'Data Visualization', 'This lesson covers data visualization techniques.'),
((SELECT course_id FROM courses WHERE title = 'Machine Learning'), 'Introduction to Machine Learning', 'This lesson introduces the fundamentals of machine learning.'),
((SELECT course_id FROM courses WHERE title = 'Machine Learning'), 'Supervised Learning', 'In this lesson, we will explore supervised learning techniques.');


-- Ro'yxatdan o'tish jadvaliga ma'lumot kiritish
INSERT INTO enrollments (user_id, course_id, enrollment_date) 
VALUES 
((SELECT user_id FROM users WHERE name = 'Firdavs' LIMIT 1), (SELECT course_id FROM courses WHERE title = 'Python Programming' LIMIT 1), CURRENT_TIMESTAMP),
((SELECT user_id FROM users WHERE name = 'Ali' LIMIT 1), (SELECT course_id FROM courses WHERE title = 'Web Development' LIMIT 1), CURRENT_TIMESTAMP),
((SELECT user_id FROM users WHERE name = 'Zaynab' LIMIT 1), (SELECT course_id FROM courses WHERE title = 'Data Science' LIMIT 1), CURRENT_TIMESTAMP),
((SELECT user_id FROM users WHERE name = 'Bobur' LIMIT 1), (SELECT course_id FROM courses WHERE title = 'Machine Learning' LIMIT 1), CURRENT_TIMESTAMP);
