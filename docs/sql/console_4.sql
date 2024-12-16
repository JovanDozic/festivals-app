select * from festivals f
left join addresses a on a.id = f.address_id
left join cities c on c.id = a.city_id
join festival_organizers fo on fo.festival_id = f.id
left join users u on u.id = fo.user_id
left join user_profiles up on up.user_id = u.id;


select * from festival_organizers;

select fo.user_id, * from festivals
join festival_organizers fo on festivals.id = fo.festival_id
join addresses a on festivals.address_id = a.id
where fo.user_id = 51 and festivals.deleted_at is null;

select * from addresses where id = 12;

select id, name, address_id from festivals where deleted_at is null;

select * from organizers
join users on organizers.user_id = users.id;

select * from festival_images;

select * from festival_employees;

select * from images;


select url, name from festival_images
join images i on festival_images.image_id = i.id
join festivals f on festival_images.festival_id = f.id

select users.username, * from users
join user_profiles up on users.id = up.user_id
join employees e on users.id = e.user_id
join festival_employees on e.user_id = festival_employees.user_id
where festival_employees.festival_id = 1;

-- get employees per festival
select * from user_profiles up
join users u on up.user_id = u.id
join employees e on u.id = e.user_id
join festival_employees fe on e.user_id = fe.user_id
where fe.festival_id = 1;

-- get employees that don't work on one festival
select username, * from user_profiles up
join users u on up.user_id = u.id
join employees e on u.id = e.user_id
where e.user_id not in (select user_id from festival_employees where festival_id = 20);

-- get all employees in system
select username, * from user_profiles up
join users u on up.user_id = u.id
where u.role = 'EMPLOYEE';

select * from festival_employees;

select * from employees;

select * from users;