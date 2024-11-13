insert into momentum(user_id, steps, calories, distance, max_heart_rate, min_heart_rate, date, created_at) 
values('uids893njkdf89', 10000, 2000, 10000, 200, 100, '2024-11-05', '2024-11-05 12:34:56'),
('uids893njkdf89', 5000, 1000, 5000, 180, 90, '2024-11-07', '2024-11-07 12:55:00'),
('uids893njkdf89', 8000, 1500, 8000, 190, 95, '2024-11-10', '2024-11-10 15:21:42');


insert into sleep(user_id, hours, started_at, ended_at, deep_sleep, light_sleep, rem_sleep, wake, date, created_at)
values('uids893njkdf89', 8, '2024-11-05 00:00:00', '2024-11-05 08:00:00', 2, 3, 2, 1, '2024-11-05', '2024-11-05 08:00:00'),
('uids893njkdf89', 7, '2024-11-07 00:00:00', '2024-11-07 07:00:00', 1, 3, 2, 1, '2024-11-07', '2024-11-07 07:00:00'),
('uids893njkdf89', 6, '2024-11-10 00:00:00', '2024-11-10 06:00:00', 1, 2, 2, 1, '2024-11-10', '2024-11-10 06:00:00');

insert into meal(user_id, meal_name, calorie_per_100g, date, created_at)
values('uids893njkdf89', 'rice', 100, '2024-11-05', '2024-11-05 12:34:56'),
('uids893njkdf89', 'bread', 200, '2024-11-07', '2024-11-07 3:32:02'),
('uids893njkdf89', 'noodle', 300, '2024-11-10', '2024-11-10 22:04:14');