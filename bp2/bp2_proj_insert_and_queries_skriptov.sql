INSERT INTO attendee (atd_id, atd_email, atd_name, atd_last_name, atd_phone_number) VALUES (1, 'alice.johnson@edmlover.com', 'Alice', 'Johnson', '+13354100282');
INSERT INTO attendee (atd_id, atd_email, atd_name, atd_last_name, atd_phone_number) VALUES (2, 'bob.smith@edmlover.com', 'Bob', 'Smith', '+11187824877');
INSERT INTO attendee (atd_id, atd_email, atd_name, atd_last_name, atd_phone_number) VALUES (3, 'charlie.brown@edmlover.com', 'Charlie', 'Brown', '+12402976366');
INSERT INTO attendee (atd_id, atd_email, atd_name, atd_last_name, atd_phone_number) VALUES (4, 'david.wilson@edmlover.com', 'David', 'Wilson', '+14342879928');
INSERT INTO attendee (atd_id, atd_email, atd_name, atd_last_name, atd_phone_number) VALUES (5, 'ella.martinez@edmlover.com', 'Ella', 'Martinez', '+13248747354');
INSERT INTO attendee (atd_id, atd_email, atd_name, atd_last_name, atd_phone_number) VALUES (6, 'fiona.garcia@edmlover.com', 'Fiona', 'Garcia', '+17012998760');
INSERT INTO attendee (atd_id, atd_email, atd_name, atd_last_name, atd_phone_number) VALUES (7, 'george.miller@edmlover.com', 'George', 'Miller', '+13100774108');
INSERT INTO attendee (atd_id, atd_email, atd_name, atd_last_name, atd_phone_number) VALUES (8, 'hannah.davis@edmlover.com', 'Hannah', 'Davis', '+17686057141');
INSERT INTO attendee (atd_id, atd_email, atd_name, atd_last_name, atd_phone_number) VALUES (9, 'irene.lopez@edmlover.com', 'Irene', 'Lopez', '+12759425002');
INSERT INTO attendee (atd_id, atd_email, atd_name, atd_last_name, atd_phone_number) VALUES (10, 'jack.jones@edmlover.com', 'Jack', 'Jones', '+18433400045');
INSERT INTO attendee (atd_id, atd_email, atd_name, atd_last_name, atd_phone_number) VALUES (11, 'kathy.perez@edmlover.com', 'Kathy', 'Perez', '+12855333464');


INSERT INTO event (evt_id, evt_name, evt_start_date, evt_end_date, evt_location, event_evt_id) VALUES (1, 'Tomorrowland', TO_DATE('2024-07-20', 'YYYY-MM-DD'), TO_DATE('2024-07-29', 'YYYY-MM-DD'), 'Boom, Belgium', NULL);
INSERT INTO event (evt_id, evt_name, evt_start_date, evt_end_date, evt_location, event_evt_id) VALUES (2, 'Ultra Miami', TO_DATE('2024-03-28', 'YYYY-MM-DD'), TO_DATE('2024-03-30', 'YYYY-MM-DD'), 'Miami, USA', NULL);
INSERT INTO event (evt_id, evt_name, evt_start_date, evt_end_date, evt_location, event_evt_id) VALUES (3, 'EDC Las Vegas', TO_DATE('2024-05-16', 'YYYY-MM-DD'), TO_DATE('2024-05-18', 'YYYY-MM-DD'), 'Las Vegas, USA', NULL);
INSERT INTO event (evt_id, evt_name, evt_start_date, evt_end_date, evt_location, event_evt_id) VALUES (4, 'Creamfields', TO_DATE('2024-08-26', 'YYYY-MM-DD'), TO_DATE('2024-08-29', 'YYYY-MM-DD'), 'Daresbury, UK', NULL);
INSERT INTO event (evt_id, evt_name, evt_start_date, evt_end_date, evt_location, event_evt_id) VALUES (5, 'Mysteryland', TO_DATE('2024-08-23', 'YYYY-MM-DD'), TO_DATE('2024-08-25', 'YYYY-MM-DD'), 'Haarlemmermeer, Netherlands', NULL);


INSERT INTO tickettype (tty_id, tty_type, tty_quantity, event_evt_id, tickettype_type) VALUES (1, 'General Admission', 10000, 1, 'TicketType');
INSERT INTO tickettype (tty_id, tty_type, tty_quantity, event_evt_id, tickettype_type) VALUES (2, 'VIP', 2000, 1, 'VIPTicket');
INSERT INTO tickettype (tty_id, tty_type, tty_quantity, event_evt_id, tickettype_type) VALUES (3, 'GAPlus', 3000, 1, 'GAPlusTicket');
INSERT INTO tickettype (tty_id, tty_type, tty_quantity, event_evt_id, tickettype_type) VALUES (4, 'General Admission', 15000, 2, 'TicketType');
INSERT INTO tickettype (tty_id, tty_type, tty_quantity, event_evt_id, tickettype_type) VALUES (5, 'VIP', 2500, 2, 'VIPTicket');
INSERT INTO tickettype (tty_id, tty_type, tty_quantity, event_evt_id, tickettype_type) VALUES (6, 'GAPlus', 3500, 2, 'GAPlusTicket');
INSERT INTO tickettype (tty_id, tty_type, tty_quantity, event_evt_id, tickettype_type) VALUES (7, 'General Admission', 20000, 3, 'TicketType');
INSERT INTO tickettype (tty_id, tty_type, tty_quantity, event_evt_id, tickettype_type) VALUES (8, 'VIP', 3000, 3, 'VIPTicket');
INSERT INTO tickettype (tty_id, tty_type, tty_quantity, event_evt_id, tickettype_type) VALUES (9, 'GAPlus', 4000, 3, 'GAPlusTicket');
INSERT INTO tickettype (tty_id, tty_type, tty_quantity, event_evt_id, tickettype_type) VALUES (10, 'General Admission', 12000, 4, 'TicketType');
INSERT INTO tickettype (tty_id, tty_type, tty_quantity, event_evt_id, tickettype_type) VALUES (11, 'VIP', 1800, 4, 'VIPTicket');
INSERT INTO tickettype (tty_id, tty_type, tty_quantity, event_evt_id, tickettype_type) VALUES (12, 'GAPlus', 2500, 4, 'GAPlusTicket');
INSERT INTO tickettype (tty_id, tty_type, tty_quantity, event_evt_id, tickettype_type) VALUES (13, 'General Admission', 8000, 5, 'TicketType');
INSERT INTO tickettype (tty_id, tty_type, tty_quantity, event_evt_id, tickettype_type) VALUES (14, 'VIP', 1200, 5, 'VIPTicket');
INSERT INTO tickettype (tty_id, tty_type, tty_quantity, event_evt_id, tickettype_type) VALUES (15, 'GAPlus', 2000, 5, 'GAPlusTicket');


INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (1, 150.64, 1); -- 1-9 su evt_id 1
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (2, 177.51, 1);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (3, 196.91, 1);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (4, 196.44, 1);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (5, 157.26, 1);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (6, 173.15, 1);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (7, 175.22, 1);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (8, 196.31, 1);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (9, 189.8, 1);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (10, 195.16, 1);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (11, 234.09, 2); -- 11,12 su evt_id 1
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (12, 234.53, 2);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (13, 100.02, 3); -- 13-15 su evt_id 1
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (14, 118.29, 3);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (15, 101.36, 3);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (16, 181.34, 4); -- 16-30 su evt_id 2
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (17, 151.18, 4);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (18, 158.99, 4);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (19, 158.44, 4);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (20, 181.61, 4);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (21, 157.52, 4);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (22, 187.05, 4);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (23, 181.78, 4);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (24, 156.2, 4);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (25, 161.89, 4);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (26, 169.24, 4);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (27, 198.7, 4);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (28, 198.61, 4);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (29, 190.07, 4);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (30, 194.39, 4);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (31, 229.2, 5); -- 31-32 su evt_id 2
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (32, 235.65, 5);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (33, 134.08, 6); -- 33-35 su evt_id 2
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (34, 126.98, 6);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (35, 106.01, 6);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (36, 194.43, 7); -- 36-55 su evt_id 3
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (37, 189.18, 7);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (38, 196.32, 7);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (39, 161.24, 7);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (40, 195.17, 7);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (41, 159.25, 7);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (42, 167.15, 7);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (43, 158.52, 7);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (44, 184.64, 7);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (45, 170.97, 7);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (46, 194.61, 7);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (47, 168.24, 7);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (48, 157.77, 7);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (49, 172.44, 7);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (50, 155.17, 7);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (51, 167.86, 7);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (52, 178.2, 7);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (53, 187.72, 7);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (54, 194.17, 7);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (55, 185.78, 7);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (56, 224.86, 8); -- 56-58 su evt_id 3
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (57, 221.88, 8);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (58, 217.2, 8);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (59, 116.45, 9); -- 59-62 su evt_id 3
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (60, 116.1, 9);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (61, 131.67, 9);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (62, 110.44, 9);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (63, 186.99, 10); -- 63-74 su evt_id 4
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (64, 160.14, 10);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (65, 178.6, 10);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (66, 171.94, 10);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (67, 180.16, 10);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (68, 197.29, 10);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (69, 173.7, 10);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (70, 177.88, 10);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (71, 175.3, 10);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (72, 176.35, 10);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (73, 195.78, 10);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (74, 194.38, 10);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (75, 207.5, 11); -- 75 je evt_id 4
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (76, 103.63, 12); -- 76,77 su evt_id 4
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (77, 106.34, 12);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (78, 178.52, 13); -- 78-85 su evt_id 5
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (79, 151.49, 13);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (80, 165.52, 13);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (81, 159.67, 13);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (82, 197.97, 13);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (83, 187.91, 13);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (84, 191.05, 13);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (85, 168.55, 13);
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (86, 205.39, 14); -- 86 je evt_id 5
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (87, 115.48, 15); -- 87-88 su evt_id 5
INSERT INTO ticket (tck_id, tck_price, tickettype_tty_id) VALUES (88, 103.27, 15);


INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (1, 1, TO_DATE('2024-03-01', 'YYYY-MM-DD'), 1, 6);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (1, 2, TO_DATE('2024-04-18', 'YYYY-MM-DD'), 2, 10);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (1, 3, TO_DATE('2024-09-02', 'YYYY-MM-DD'), 3, 3);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (1, 4, TO_DATE('2024-04-02', 'YYYY-MM-DD'), 4, 9);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (1, 5, TO_DATE('2024-07-03', 'YYYY-MM-DD'), 5, 4);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (1, 6, TO_DATE('2024-08-12', 'YYYY-MM-DD'), 6, 6);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (1, 7, TO_DATE('2024-05-15', 'YYYY-MM-DD'), 7, 3);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (1, 8, TO_DATE('2024-06-11', 'YYYY-MM-DD'), 8, 2);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (1, 9, TO_DATE('2024-04-01', 'YYYY-MM-DD'), 9, 1);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (1, 10, TO_DATE('2024-08-22', 'YYYY-MM-DD'), 10, 3);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (1, 11, TO_DATE('2024-12-23', 'YYYY-MM-DD'), 11, 3);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (1, 12, TO_DATE('2024-05-19', 'YYYY-MM-DD'), 12, 2);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (1, 13, TO_DATE('2024-05-07', 'YYYY-MM-DD'), 13, 1);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (1, 14, TO_DATE('2024-09-20', 'YYYY-MM-DD'), 14, 9);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (1, 15, TO_DATE('2024-03-12', 'YYYY-MM-DD'), 15, 1);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (2, 16, TO_DATE('2024-11-02', 'YYYY-MM-DD'), 16, 2);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (2, 17, TO_DATE('2024-12-14', 'YYYY-MM-DD'), 17, 3);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (2, 18, TO_DATE('2024-07-06', 'YYYY-MM-DD'), 18, 1);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (2, 19, TO_DATE('2024-10-07', 'YYYY-MM-DD'), 19, 5);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (2, 20, TO_DATE('2024-12-05', 'YYYY-MM-DD'), 20, 10);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (2, 21, TO_DATE('2024-05-13', 'YYYY-MM-DD'), 21, 8);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (2, 22, TO_DATE('2024-01-15', 'YYYY-MM-DD'), 22, 9);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (2, 23, TO_DATE('2024-06-12', 'YYYY-MM-DD'), 23, 10);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (2, 24, TO_DATE('2024-06-25', 'YYYY-MM-DD'), 24, 2);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (2, 25, TO_DATE('2024-01-31', 'YYYY-MM-DD'), 25, 2);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (2, 26, TO_DATE('2024-01-12', 'YYYY-MM-DD'), 26, 9);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (2, 27, TO_DATE('2024-01-10', 'YYYY-MM-DD'), 27, 6);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (2, 28, TO_DATE('2024-09-14', 'YYYY-MM-DD'), 28, 2);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (2, 29, TO_DATE('2024-08-28', 'YYYY-MM-DD'), 29, 7);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (2, 30, TO_DATE('2024-10-29', 'YYYY-MM-DD'), 30, 7);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (2, 31, TO_DATE('2024-09-05', 'YYYY-MM-DD'), 31, 8);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (2, 32, TO_DATE('2024-11-10', 'YYYY-MM-DD'), 32, 10);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (2, 33, TO_DATE('2024-11-14', 'YYYY-MM-DD'), 33, 6);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (2, 34, TO_DATE('2024-04-19', 'YYYY-MM-DD'), 34, 1);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (2, 35, TO_DATE('2024-05-02', 'YYYY-MM-DD'), 35, 1);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (3, 36, TO_DATE('2024-07-07', 'YYYY-MM-DD'), 36, 4);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (3, 37, TO_DATE('2024-06-16', 'YYYY-MM-DD'), 37, 8);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (3, 38, TO_DATE('2024-05-27', 'YYYY-MM-DD'), 38, 1);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (3, 39, TO_DATE('2024-01-08', 'YYYY-MM-DD'), 39, 10);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (3, 40, TO_DATE('2024-07-13', 'YYYY-MM-DD'), 40, 7);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (3, 41, TO_DATE('2024-01-07', 'YYYY-MM-DD'), 41, 4);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (3, 42, TO_DATE('2024-04-06', 'YYYY-MM-DD'), 42, 6);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (3, 43, TO_DATE('2024-06-01', 'YYYY-MM-DD'), 43, 6);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (3, 44, TO_DATE('2024-11-08', 'YYYY-MM-DD'), 44, 10);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (3, 45, TO_DATE('2024-11-12', 'YYYY-MM-DD'), 45, 5);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (3, 46, TO_DATE('2024-09-05', 'YYYY-MM-DD'), 46, 5);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (3, 47, TO_DATE('2024-07-08', 'YYYY-MM-DD'), 47, 2);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (3, 48, TO_DATE('2024-12-23', 'YYYY-MM-DD'), 48, 3);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (3, 49, TO_DATE('2024-09-22', 'YYYY-MM-DD'), 49, 6);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (3, 50, TO_DATE('2024-08-17', 'YYYY-MM-DD'), 50, 10);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (3, 51, TO_DATE('2024-07-07', 'YYYY-MM-DD'), 51, 7);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (3, 52, TO_DATE('2024-11-15', 'YYYY-MM-DD'), 52, 3);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (3, 53, TO_DATE('2024-04-17', 'YYYY-MM-DD'), 53, 10);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (3, 54, TO_DATE('2024-12-03', 'YYYY-MM-DD'), 54, 6);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (3, 55, TO_DATE('2024-05-12', 'YYYY-MM-DD'), 55, 4);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (3, 56, TO_DATE('2024-07-20', 'YYYY-MM-DD'), 56, 4);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (3, 57, TO_DATE('2024-06-25', 'YYYY-MM-DD'), 57, 7);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (3, 58, TO_DATE('2024-08-19', 'YYYY-MM-DD'), 58, 5);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (3, 59, TO_DATE('2024-04-05', 'YYYY-MM-DD'), 59, 3);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (3, 60, TO_DATE('2024-07-05', 'YYYY-MM-DD'), 60, 9);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (3, 61, TO_DATE('2024-05-26', 'YYYY-MM-DD'), 61, 6);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (3, 62, TO_DATE('2024-11-07', 'YYYY-MM-DD'), 62, 3);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (4, 63, TO_DATE('2024-04-19', 'YYYY-MM-DD'), 63, 10);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (4, 64, TO_DATE('2024-04-17', 'YYYY-MM-DD'), 64, 5);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (4, 65, TO_DATE('2024-09-11', 'YYYY-MM-DD'), 65, 6);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (4, 66, TO_DATE('2024-01-06', 'YYYY-MM-DD'), 66, 2);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (4, 67, TO_DATE('2024-07-27', 'YYYY-MM-DD'), 67, 8);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (4, 68, TO_DATE('2024-04-27', 'YYYY-MM-DD'), 68, 5);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (4, 69, TO_DATE('2024-08-09', 'YYYY-MM-DD'), 69, 8);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (4, 70, TO_DATE('2024-01-09', 'YYYY-MM-DD'), 70, 2);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (4, 71, TO_DATE('2024-12-10', 'YYYY-MM-DD'), 71, 8);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (4, 72, TO_DATE('2024-04-14', 'YYYY-MM-DD'), 72, 6);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (4, 73, TO_DATE('2024-01-02', 'YYYY-MM-DD'), 73, 10);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (4, 74, TO_DATE('2024-12-03', 'YYYY-MM-DD'), 74, 8);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (4, 75, TO_DATE('2024-07-29', 'YYYY-MM-DD'), 75, 10);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (4, 76, TO_DATE('2024-10-19', 'YYYY-MM-DD'), 76, 9);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (4, 77, TO_DATE('2024-03-04', 'YYYY-MM-DD'), 77, 1);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (5, 78, TO_DATE('2024-06-27', 'YYYY-MM-DD'), 78, 3);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (5, 79, TO_DATE('2024-09-28', 'YYYY-MM-DD'), 79, 8);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (5, 80, TO_DATE('2024-01-12', 'YYYY-MM-DD'), 80, 1);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (5, 81, TO_DATE('2024-04-06', 'YYYY-MM-DD'), 81, 8);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (5, 82, TO_DATE('2024-01-11', 'YYYY-MM-DD'), 82, 7);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (5, 83, TO_DATE('2024-12-11', 'YYYY-MM-DD'), 83, 8);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (5, 84, TO_DATE('2024-05-27', 'YYYY-MM-DD'), 84, 8);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (5, 85, TO_DATE('2024-04-23', 'YYYY-MM-DD'), 85, 7);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (5, 86, TO_DATE('2024-08-23', 'YYYY-MM-DD'), 86, 2);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (5, 87, TO_DATE('2024-05-24', 'YYYY-MM-DD'), 87, 1);
INSERT INTO purchase (event_evt_id, pch_id, pch_date, ticket_tck_id, attendee_atd_id) VALUES (5, 88, TO_DATE('2024-08-28', 'YYYY-MM-DD'), 88, 2);



-- Kompleksni upiti:

-- Vraca ukupan broj prodatih tiketa po fetivalu sa njihovim prosecnim cenama.
CREATE OR REPLACE VIEW event_ticket_summary AS
SELECT 
    e.evt_id,
    e.evt_name,
    e.evt_start_date,
    e.evt_end_date,
    COUNT(tt.tty_id) AS total_sold_tickets,
    AVG(t.tck_price) AS average_ticket_price
FROM 
    event e
LEFT JOIN tickettype tt ON e.evt_id = tt.event_evt_id
LEFT JOIN ticket t ON tt.tty_id = t.tickettype_tty_id
GROUP BY 
    e.evt_id, e.evt_name, e.evt_start_date, e.evt_end_date
ORDER BY 
    e.evt_start_date;
    
SELECT * FROM event_ticket_summary;

-- Prikazuje sve korisnike koje su kupili tikete, i koji su to tiketi.
CREATE OR REPLACE VIEW attendee_ticket_details AS
SELECT 
    a.atd_id,
    a.atd_name,
    a.atd_last_name,
    e.evt_name,
    tt.tty_type,
    t.tck_price
FROM 
    attendee a
LEFT JOIN purchase p ON a.atd_id = p.attendee_atd_id
LEFT JOIN ticket t ON p.ticket_tck_id = t.tck_id
LEFT JOIN tickettype tt ON t.tickettype_tty_id = tt.tty_id
LEFT JOIN event e ON tt.event_evt_id = e.evt_id
WHERE 
    t.tck_price IS NOT NULL
ORDER BY 
    a.atd_name, a.atd_last_name, e.evt_name;

SELECT * FROM attendee_ticket_details;

-- Racuna ukupan broj mogucih prodatih tiketa, i potencijalne prihode ukoliko svi budu rasprodati.
CREATE OR REPLACE VIEW event_total_ticket_sales_potential AS
SELECT 
    e.evt_id,
    e.evt_name,
    SUM(tt.tty_quantity) AS total_tickets_available,
    SUM(tt.tty_quantity * t.tck_price) AS total_potential_revenue
FROM 
    event e
LEFT JOIN tickettype tt ON e.evt_id = tt.event_evt_id
LEFT JOIN ticket t ON tt.tty_id = t.tickettype_tty_id
GROUP BY 
    e.evt_id, e.evt_name
ORDER BY 
    e.evt_id;
