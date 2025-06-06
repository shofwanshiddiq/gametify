use gametify;

INSERT INTO packages (room_id, package_type, price, open_hrs_start, open_hrs_end) VALUES
-- Room 1
(1, 'Midnight', 15200, '2025-04-13 00:00:00', '2025-04-13 05:00:00'),
(1, 'Regular',  9800,  '2025-04-13 08:00:00', '2025-04-13 20:00:00'),
(1, 'Night',    14800, '2025-04-13 20:00:00', '2025-04-13 00:00:00'),

-- Room 2
(2, 'Midnight', 15100, '2025-04-13 00:00:00', '2025-04-13 05:00:00'),
(2, 'Regular',  10200, '2025-04-13 05:00:00', '2025-04-13 20:00:00'),
(2, 'Night',    14900, '2025-04-13 20:00:00', '2025-04-13 00:00:00'),

-- Room 3
(3, 'Midnight', 14700, '2025-04-13 00:00:00', '2025-04-13 04:00:00'),
(3, 'Regular',  9900,  '2025-04-13 04:00:00', '2025-04-13 20:00:00'),
(3, 'Night',    15300, '2025-04-13 20:00:00', '2025-04-13 00:00:00'),

-- Room 4
(4, 'Midnight', 15000, '2025-04-13 00:00:00', '2025-04-13 05:00:00'),
(4, 'Regular',  9700,  '2025-04-13 05:00:00', '2025-04-13 19:00:00'),
(4, 'Night',    15100, '2025-04-13 19:00:00', '2025-04-13 00:00:00'),

-- Room 5
(5, 'Midnight', 14900, '2025-04-13 00:00:00', '2025-04-13 07:00:00'),
(5, 'Regular',  10100, '2025-04-13 07:00:00', '2025-04-13 20:00:00'),
(5, 'Night',    14700, '2025-04-13 20:00:00', '2025-04-13 00:00:00'),

-- Room 6
(6, 'Midnight', 15200, '2025-04-13 00:00:00', '2025-04-13 08:00:00'),
(6, 'Regular',  9900,  '2025-04-13 08:00:00', '2025-04-13 20:00:00'),
(6, 'Night',    15000, '2025-04-13 20:00:00', '2025-04-13 00:00:00'),

-- Room 7
(7, 'Midnight', 15300, '2025-04-13 00:00:00', '2025-04-13 02:00:00'),
(7, 'Regular',  10000, '2025-04-13 09:00:00', '2025-04-13 22:00:00'),
(7, 'Night',    14900, '2025-04-13 22:00:00', '2025-04-13 00:00:00'),

-- Room 8
(8, 'Midnight', 15100, '2025-04-13 00:00:00', '2025-04-13 04:00:00'),
(8, 'Regular',  9700,  '2025-04-13 12:00:00', '2025-04-13 19:00:00'),
(8, 'Night',    15000, '2025-04-13 19:00:00', '2025-04-13 00:00:00'),

-- Room 9
(9, 'Midnight', 14800, '2025-04-13 00:00:00', '2025-04-13 02:00:00'),
(9, 'Regular',  10100, '2025-04-13 08:00:00', '2025-04-13 21:00:00'),
(9, 'Night',    14700, '2025-04-13 21:00:00', '2025-04-13 00:00:00'),

-- Room 10
(10, 'Midnight', 15000, '2025-04-13 09:00:00', '2025-04-13 12:00:00'),
(10, 'Regular',  10000, '2025-04-13 12:00:00', '2025-04-13 17:00:00'),
(10, 'Night',    15000, '2025-04-13 17:00:00', '2025-04-13 23:00:00');
