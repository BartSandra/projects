-- Session #1 then session #2
BEGIN;

SET TRANSACTION ISOLATION LEVEL SERIALIZABLE;

-- Session #1
SELECT * FROM pizzeria WHERE name = 'Pizza Hut';

-- Session #2 
UPDATE pizzeria SET rating = 3 WHERE name = 'Pizza Hut';

COMMIT;

-- Session #1
SELECT * FROM pizzeria WHERE name = 'Pizza Hut';

COMMIT;

-- Session #1 then session #2
SELECT * FROM pizzeria WHERE name = 'Pizza Hut';