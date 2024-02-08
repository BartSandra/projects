CREATE FUNCTION func_minimum(VARIADIC arr numeric []) RETURNS numeric AS $$
SELECT MIN(i)
FROM UNNEST(arr) g(i);
$$ LANGUAGE SQL;

SELECT func_minimum(VARIADIC arr => ARRAY [10.0, -1.0, 5.0, 4.4]);