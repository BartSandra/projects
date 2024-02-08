WITH generate_series AS (
SELECT generate_series::date
FROM generate_series('2022-01-01', '2022-01-10', interval '1 day') AS generate_series)
SELECT generate_series::date AS missing_date
FROM (SELECT * FROM person_visits WHERE person_id IN (1, 2)) AS person_visits 
RIGHT JOIN generate_series AS generate_series ON person_visits.visit_date = generate_series
WHERE person_visits.id IS NULL
ORDER BY missing_date;