insert into momentum(user_id, steps, calories, distance, max_heart_rate, min_heart_rate, date) 
values('uids893njkdf89', 10000, 2000, 10000, 200, 100, '2021-01-01'),
('uids893njkdf90', 5000, 1000, 5000, 180, 90, '2021-01-01'),
('uids893njkdf91', 8000, 1500, 8000, 190, 95, '2021-01-01');

insert into sleep(user_id, hours, started_at, ended_at, deep_sleep, light_sleep, rem_sleep, wake, date)
values('uids893njkdf89', 8, '2021-01-01 00:00:00', '2021-01-01 08:00:00', 2, 3, 2, 1, '2021-01-01'),
('uids893njkdf90', 7, '2021-01-01 00:00:00', '2021-01-01 07:00:00', 1, 3, 2, 1, '2021-01-01'),
('uids893njkdf91', 6, '2021-01-01 00:00:00', '2021-01-01 06:00:00', 1, 2, 2, 1, '2021-01-01');

insert into meal(user_id, meal_name, calorie_per_100g, date)
values('uids893njkdf89', 'rice', 100, '2021-01-01'),
('uids893njkdf90', 'bread', 200, '2021-01-01'),
('uids893njkdf91', 'noodle', 300, '2021-01-01');