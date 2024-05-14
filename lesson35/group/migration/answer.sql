-- Vazifa leetcodedan
-- https://leetcode.com/problems/largest-positive-integer-that-exists-with-its-negative/submissions/1247214138?envType=daily-question&envId=2024-05-02



-------------------- Vazifa 1 ---------------------
-- Bu vazifa tablelar yaratilgan vaqtda to'g'irlab yozildi

-------------------- Vazifa 2 ---------------------
WITH top_students AS (
SELECT
c.course_name,
s.name AS student_name,
g.grade,
RANK() OVER (PARTITION BY c.course_id ORDER BY g.grade DESC) AS rank
FROM
course c
JOIN student_course sc ON c.course_id = sc.course_id
JOIN grade g ON sc.student_course_id = g.student_course_id
JOIN student s ON sc.student_id = s.student_id
)

SELECT
course_name, student_name, grade
FROM
top_students
WHERE
rank = 1;

-------------------- Vazifa 3 ---------------------
WITH avagerage_grade_of_courses AS(
SELECT
AVG(g.grade) AS avgrade,
sc.course_id
FROM
student_course sc
JOIN grade g ON sc.student_course_id = g.student_course_id

GROUP BY
sc.course_id
)

SELECT
c.course_name, ROUND(avcg.avgrade::NUMERIC, 2)
FROM
course c
JOIN avagerage_grade_of_courses avcg ON avcg.course_id = c.course_id;

-------------------- Vazifa 4 ---------------------
WITH young_students AS (
SELECT
c.course_name,
s.name AS student_name,
s.age,
RANK() OVER (PARTITION BY c.course_id ORDER BY s.age) AS age_rank
FROM
course c
JOIN student_course sc ON c.course_id = sc.course_id

JOIN student s ON sc.student_id = s.student_id
)

SELECT
course_name, student_name, age
FROM
young_students
WHERE
age_rank = 1;

-------------------- Vazifa 5 ---------------------

WITH avagerage_grade_of_courses AS(
SELECT
AVG(g.grade) AS avgrade,
sc.course_id
FROM
student_course sc
JOIN grade g ON sc.student_course_id = g.student_course_id

GROUP BY
sc.course_id
)

SELECT
c.course_name,ROUND(avcg.avgrade::NUMERIC, 2) AS Eng_yaxshi_guruh
FROM
course c
JOIN avagerage_grade_of_courses avcg ON avcg.course_id = c.course_id

ORDER BY
avcg.avgrade DESC
LIMIT 1;