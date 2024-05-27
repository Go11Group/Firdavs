-- UPDATE bu jadvaldagi bir yoki bir nechta ma'lumotlarni yangilaash uchun ishlatiladi. U bilan birga Set degangan kamandaham ishlatiladi, u komanda qaysi ustunni olishni bildiradi.
UPDATE courses
SET published_date = '2020-08-01' 
WHERE course_id = 3;

--DELETE ma'lumotlar bazasidan mavjud yozuvlarni olib tashlashga yordam beradi. 
DELETE FROM courses WHERE course_id = 3;

--GROUP BY - bir xil ma'lumotlarni guruhlarga ajratuvchi SQL bandidir. Ko'pincha natijalar to'plamini bir yoki bir nechta ustunlar bo'yicha guruhlash uchun agregat funktsiyalari (COUNT, MAX, MIN, SUM, AVG) bilan ishlatiladi.
SELECT column1, column2, ..., aggregate_function(column_name)
FROM table_name
WHERE condition
GROUP BY column1, column2, ...;

--ORDER BY bu jadvallarni saralash uchun ishlatiladi. uning ASC yoki DESC qilish mumkun.

SELECT column1, column2, ...
FROM table_name
ORDER BY column1, column2, ... ASC & DESC;
