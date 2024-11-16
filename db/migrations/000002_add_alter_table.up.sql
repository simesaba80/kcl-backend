ALTER TABLE activities
ADD CONSTRAINT fk_activities_user
FOREIGN KEY (user_id) REFERENCES users(uid);

ALTER TABLE sleep
ADD CONSTRAINT fk_sleep_user
FOREIGN KEY (user_id) REFERENCES users(uid);

ALTER TABLE meal
ADD CONSTRAINT fk_meal_user
FOREIGN KEY (user_id) REFERENCES users(uid);