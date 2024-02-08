CREATE INDEX idx_person_name
ON person(UPPER(name));

SET ENABLE_SEQSCAN = OFF;

EXPLAIN ANALYZE SELECT *
FROM person
WHERE UPPER(name) = 'Kate';