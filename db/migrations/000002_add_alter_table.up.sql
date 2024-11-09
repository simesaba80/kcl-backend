ALTER TABLE momentum
ADD CONSTRAINT fk_momentum_user
FOREIGN KEY (user_id) REFERENCES users(uid);

ALTER TABLE sleep
ADD CONSTRAINT fk_sleep_user
FOREIGN KEY (user_id) REFERENCES users(uid);

ALTER TABLE meal
ADD CONSTRAINT fk_meal_user
FOREIGN KEY (user_id) REFERENCES users(uid);