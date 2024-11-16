ALTER TABLE activities
DROP CONSTRAINT fk_activities_user;

ALTER TABLE sleep
DROP CONSTRAINT fk_sleep_user;

ALTER TABLE meal
DROP CONSTRAINT fk_meal_user;
