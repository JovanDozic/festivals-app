select * from users order by id;

select * from user_profiles;

select u.username, up.first_name, a.street, a.number, c.name, cc.name from user_profiles up
join users u on u.id = up.user_id
join addresses a on up.address_id = a.id
join cities c on a.city_id = c.id
join countries cc on c.country_id = cc.id;

select * from countries;

select * from addresses;

select * from cities;

select * from cities;

select a.street, a.number, c.name, c.postal_code, cc.name
from addresses a
join cities c on a.city_id = c.id
join countries cc on c.country_id = cc.id;

delete from addresses;
delete from cities;

select u.username, u.email, up.first_name, up.last_name, up.date_of_birth, up.phone_number, a.street, c.name, cc.name from user_profiles up
join users u on u.id = up.user_id
left join addresses a on up.address_id = a.id
left join cities c on a.city_id = c.id
left join countries cc on cc.id = c.country_id
left join images i on i.id = up.image_id;

select * from attendees;
select * from organizers;
select * from administrators;

delete from user_profiles;
delete from addresses;
delete from attendees;
delete from users;


-- This way we can get all users with <ROLE> without doing something like users.role = 'ATTENDEE'

select * from users
right join administrators on administrators.user_id = users.id;

select * from users
right join attendees on attendees.user_id = users.id;

delete from attendees where user_id > 5;
delete from users where id > 5;
delete from addresses where id > 7;

select * from organizers
left join users on organizers.user_id = users.id;

select * from price_lists;

select * from festival_organizers;
select * from festival_employees;
select * from festival_images;
drop table festival_employees;

select * from ticket_types;
select * from package_addons;
select * from package_addon_images;

select * from camp_addons;

select * from camp_equipments;

select * from transport_addons;

select * from orders;

select * from festival_tickets;

select * from bracelets;

select * from activation_help_requests;