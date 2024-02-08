DROP FUNCTION IF EXISTS fnc_persons_male;
DROP FUNCTION IF EXISTS fnc_persons_female;

CREATE OR REPLACE FUNCTION fnc_persons(IN pgender VARCHAR DEFAULT 'female')
RETURNS TABLE(id BIGINT, name VARCHAR, age INTEGER, gender VARCHAR, address VARCHAR)
	AS $$
		SELECT * FROM person
		WHERE person.gender = pgender;
	$$ LANGUAGE SQL;

SELECT *
FROM fnc_persons(pgender := 'male');
SELECT *
FROM fnc_persons();