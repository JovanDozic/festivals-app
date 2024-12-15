-- Insert attendees into the database
DO $$
DECLARE
    usernames TEXT[] := ARRAY['matija', 'milena', 'isidora', 'dana', 'jovica', 'tamara', 'dusica', 'veronika', 'andjela', 'nikola', 'andrija', 'nemanja'];
    username TEXT;
    i INTEGER := 101;
    hashed_password TEXT;
BEGIN
    FOREACH username IN ARRAY usernames LOOP

        hashed_password := MD5(username);

        -- Insert user
        INSERT INTO users (id, username, email, password, role, created_at, updated_at)
        VALUES (i, username, username || '@mock.com', hashed_password, 'attendee', NOW(), NOW());

        -- Insert user profile
        INSERT INTO user_profiles (id, first_name, last_name, date_of_birth, phone_number, user_id, address_id)
        VALUES (
                i,
            initcap(username), 
            initcap(username) || 'ović', 
            '2000-01-01', 
            '+381601234567', 
            i, 
            NULL
        );

        -- Insert address for the user (Novi Sad has city_id = 1)
        INSERT INTO addresses (id, street, number, city_id, created_at, updated_at)
        VALUES (i, 'Vojvođanskih Brigada', '7', 1, NOW(), NOW());

        -- Update the user's profile with address ID
        UPDATE user_profiles SET address_id = i WHERE user_id = i;

        -- Insert into attendees table
        INSERT INTO attendees (user_id)
        VALUES (i);

        i := i + 1;
    END LOOP;
END $$;