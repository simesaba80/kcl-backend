insert into momentum(user_id, steps, calories, distance, max_heart_rate, min_heart_rate, date, created_at) 
values('uids893njkdf89', 10000, 2000, 10000, 200, 100, '2024-11-05', '2024-11-05 12:34:56'),
('uids893njkdf89', 5000, 1000, 5000, 180, 90, '2024-11-07', '2024-11-07 12:55:00'),
('uids893njkdf89', 8000, 1500, 8000, 190, 95, '2024-11-10', '2024-11-10 15:21:42');


insert into sleep(user_id, hours, started_at, ended_at, deep_sleep, light_sleep, rem_sleep, wake, date, created_at)
values('uids893njkdf89', 8, '2024-11-05 00:00:00', '2024-11-05 08:00:00', 2, 3, 2, 1, '2024-11-05', '2024-11-05 08:00:00'),
('uids893njkdf89', 7, '2024-11-07 00:00:00', '2024-11-07 07:00:00', 1, 3, 2, 1, '2024-11-07', '2024-11-07 07:00:00'),
('uids893njkdf89', 6, '2024-11-10 00:00:00', '2024-11-10 06:00:00', 1, 2, 2, 1, '2024-11-10', '2024-11-10 06:00:00');

insert into meal(user_id, meal_name, calories, protein, fat, carbohydrates, salt, calcium, date, created_at)
values('uids893njkdf89', 'rice', 364, 3.2, 5.6, 53.1, 0.2, 2.1, '2024-11-05', '2024-11-05 12:34:56'),
('uids893njkdf89', 'bread', 233, 5.9, 9.9, 29.8, 1.7, 1.2, '2024-11-07', '2024-11-07 3:32:02'),
('uids893njkdf89', 'noodle', 244, 15.1, 12.1, 16.6, 1.1, 3.2, '2024-11-10', '2024-11-10 22:04:14');