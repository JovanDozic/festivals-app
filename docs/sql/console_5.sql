select * from price_lists;

select id, name, * from items where festival_id = 16;

select * from ticket_types;

select item_id, * from price_list_items where item_id= 48;

select id, name, festival_id, * from items
join ticket_types tt on items.id = tt.item_id
where festival_id=3;

select * from items where festival_id = 3 and type='TICKET_TYPE';

select id, name, * from festivals;

delete from price_list_items where item_id = 49;
delete from ticket_types where item_id = 49;
delete from items where id = 49;

delete from price_list_items where item_id in (
select id from items where festival_id = 1
);
delete from ticket_types where item_id in (
    select id from items where festival_id = 1
    );
delete from package_addons where item_id in (
    select id from items where festival_id = 1
    );
delete from price_lists where festival_id = 1;
delete from items where festival_id = 1;

-- get all items

SELECT
    i.id,
    pli.id,
    i.name,
    pli.price,
    pli.is_fixed,
    pli.date_from,
    pli.date_to,
    i.available_number,
    i.remaining_number
FROM
    ticket_types tt
JOIN
    items i ON tt.item_id = i.id
JOIN
    price_list_items pli ON i.id = pli.item_id
JOIN
    price_lists pl ON pli.price_list_id = pl.id
WHERE
    pl.festival_id = 3
    AND (
        pli.is_fixed = TRUE
        OR (
            pli.is_fixed = FALSE
            AND CURRENT_DATE BETWEEN pli.date_from AND pli.date_to
        )
    );

