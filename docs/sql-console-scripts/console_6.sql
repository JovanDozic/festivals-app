
select * from price_lists;

select * from package_addons;

select id, name, type, * from items;

select  * from price_list_items pli
join items i on pli.item_id = i.id
join package_addons pa on i.id = pa.item_id
where pa.category = 'TRANSPORT'
    and i.type = 'PACKAGE_ADDON';

select c1.name, c2.name, ta.transport_type, ta.item_id, i.name, * from transport_addons ta
join items i on ta.item_id = i.id
join cities c1 on ta.departure_city_id = c1.id
join cities c2 on ta.arrival_city_id = c2.id;

delete from transport_addons;



select id from items
join package_addons on items.id = package_addons.item_id
          where package_addons.category = 'GENERAL';

delete from price_list_items where item_id in (62, 63, 64, 65, 66);
delete from package_addons where item_id in (62, 63, 64, 65, 66);
delete from items where id in (62, 63, 64, 65, 66);

select * from package_addon_images
join images on package_addon_images.image_id = images.id;


-- get all transport addons
select * from price_list_items pli
join items i on pli.item_id = i.id
join package_addons pa on i.id = pa.item_id
join transport_addons ta on pa.item_id = ta.item_id
join cities cd on ta.departure_city_id = cd.id
join countries ccd on cd.country_id = ccd.id
join cities ca on ta.arrival_city_id = ca.id
join countries cca on ca.country_id = cca.id
where i.festival_id = 16;

-- get all addons (ovo iskoristi za COUNT)
select i.name, i.id, * from price_list_items pli
join items i on pli.item_id = i.id
join package_addons pa on i.id = pa.item_id
where i.festival_id = 16
order by i.id;

-- get all general addons
select i.name, i.id, * from price_list_items pli
join items i on pli.item_id = i.id
join package_addons pa on i.id = pa.item_id
where i.festival_id = 16 and pa.category = 'GENERAL'
order by i.id;

-- get all camp addons
select
pli.id as price_list_item_id,
			pli.price_list_id,
			pli.item_id,
			i.name as item_name,
			i.description as item_description,
			i.type as item_type,
			i.available_number as item_available_number,
			i.remaining_number as item_remaining_number,
			pli.date_from as date_from,
			pli.date_to as date_to,
			pli.is_fixed as is_fixed,
			pli.price as price,
			pa.category as package_addon_category,
			ca.camp_name as camp_name,
			im.url as image_url,
			STRING_AGG(ce.name, ', ') as equipment_names
from price_list_items pli
JOIN items i ON pli.item_id = i.id
JOIN package_addons pa ON i.id = pa.item_id
JOIN camp_addons ca ON pa.item_id = ca.item_id
JOIN package_addon_images pai ON pai.item_id = i.id
JOIN images im ON pai.image_id = im.id
LEFT JOIN camp_equipments ce ON ce.item_id = i.id
where i.festival_id = 16
group by pli.id, pli.price_list_id, pli.item_id, i.name, i.description, i.type, i.available_number, i.remaining_number, pli.date_from, pli.date_to, pli.is_fixed, pli.price, pa.category, ca.camp_name, im.url
