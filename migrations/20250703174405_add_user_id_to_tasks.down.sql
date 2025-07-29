DO $$ BEGIN
ALTER TABLE tasks DROP CONSTRAINT IF EXISTS fk_tasks_user;
EXCEPTION WHEN others THEN
    RAISE NOTICE 'constraint not found, skipping';
END $$;

ALTER TABLE tasks DROP COLUMN IF EXISTS user_id;
