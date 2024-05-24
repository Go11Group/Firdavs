-- UPDATE bu jadvaldagi bir yoki bir nechta ma'lumotlarni yangilaash uchun ishlatiladi. U bilan birga Set degangan kamandaham ishlatiladi, u komanda qaysi ustunni olishni bildiradi.
UPDATE courses
SET published_date = '2020-08-01' 
WHERE course_id = 3;