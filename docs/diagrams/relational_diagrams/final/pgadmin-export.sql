CREATE TABLE IF NOT EXISTS
  activation_help_requests (
    id bigserial NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    user_entered_pin TEXT,
    user_entered_barcode TEXT,
    issue_description TEXT,
    proof_image_id BIGINT,
    status TEXT,
    bracelet_id BIGINT,
    attendee_id BIGINT,
    employee_id BIGINT,
    CONSTRAINT activation_help_requests_pkey PRIMARY KEY (id)
  );

CREATE TABLE IF NOT EXISTS
  addresses (
    id bigserial NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    street TEXT,
    "number" TEXT,
    apartment_suite TEXT,
    city_id BIGINT,
    CONSTRAINT addresses_pkey PRIMARY KEY (id)
  );

CREATE TABLE IF NOT EXISTS
  administrators (user_id bigserial NOT NULL, CONSTRAINT administrators_pkey PRIMARY KEY (user_id));

CREATE TABLE IF NOT EXISTS
  attendees (user_id bigserial NOT NULL, CONSTRAINT attendees_pkey PRIMARY KEY (user_id));

CREATE TABLE IF NOT EXISTS
  bracelets (
    id bigserial NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    pin TEXT,
    barcode_number TEXT,
    balance NUMERIC,
    status TEXT,
    festival_ticket_id BIGINT,
    attendee_id BIGINT,
    employee_id BIGINT,
    CONSTRAINT bracelets_pkey PRIMARY KEY (id)
  );

CREATE TABLE IF NOT EXISTS
  camp_addons (item_id bigserial NOT NULL, camp_name TEXT, CONSTRAINT camp_addons_pkey PRIMARY KEY (item_id));

CREATE TABLE IF NOT EXISTS
  camp_equipments (
    id bigserial NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    item_id BIGINT,
    NAME TEXT,
    CONSTRAINT camp_equipments_pkey PRIMARY KEY (id)
  );

CREATE TABLE IF NOT EXISTS
  cities (
    id bigserial NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    NAME TEXT,
    postal_code TEXT,
    country_id BIGINT,
    CONSTRAINT cities_pkey PRIMARY KEY (id)
  );

CREATE TABLE IF NOT EXISTS
  countries (
    id bigserial NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    NAME TEXT,
    nice_name TEXT,
    iso TEXT,
    iso3 TEXT,
    num_code BIGINT,
    phone_code BIGINT,
    CONSTRAINT countries_pkey PRIMARY KEY (id)
  );

CREATE TABLE IF NOT EXISTS
  employees (user_id bigserial NOT NULL, CONSTRAINT employees_pkey PRIMARY KEY (user_id));

CREATE TABLE IF NOT EXISTS
  festival_employees (festival_id BIGINT, user_id BIGINT);

CREATE TABLE IF NOT EXISTS
  festival_images (festival_id BIGINT, image_id BIGINT);

CREATE TABLE IF NOT EXISTS
  festival_organizers (festival_id BIGINT, user_id BIGINT);

CREATE TABLE IF NOT EXISTS
  festival_package_addons (festival_package_id BIGINT, item_id BIGINT);

CREATE TABLE IF NOT EXISTS
  festival_packages (
    id bigserial NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    CONSTRAINT festival_packages_pkey PRIMARY KEY (id)
  );

CREATE TABLE IF NOT EXISTS
  festival_tickets (
    id bigserial NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    item_id BIGINT,
    CONSTRAINT festival_tickets_pkey PRIMARY KEY (id)
  );

CREATE TABLE IF NOT EXISTS
  festivals (
    id bigserial NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    NAME TEXT,
    description TEXT,
    start_date TIMESTAMP WITH TIME ZONE,
    end_date TIMESTAMP WITH TIME ZONE,
    capacity BIGINT,
    status TEXT,
    store_status TEXT,
    address_id BIGINT,
    CONSTRAINT festivals_pkey PRIMARY KEY (id)
  );

CREATE TABLE IF NOT EXISTS
  images (
    id bigserial NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    url TEXT,
    CONSTRAINT images_pkey PRIMARY KEY (id)
  );

CREATE TABLE IF NOT EXISTS
  items (
    id bigserial NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    festival_id BIGINT,
    NAME TEXT,
    TYPE TEXT,
    description TEXT,
    available_number BIGINT,
    remaining_number BIGINT,
    CONSTRAINT items_pkey PRIMARY KEY (id)
  );

CREATE TABLE IF NOT EXISTS
  logs (
    id bigserial NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    TYPE TEXT,
    description TEXT,
    DATA jsonb,
    user_id BIGINT,
    CONSTRAINT logs_pkey PRIMARY KEY (id)
  );

CREATE TABLE IF NOT EXISTS
  orders (
    id bigserial NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    total_amount NUMERIC,
    user_id BIGINT,
    festival_ticket_id BIGINT,
    festival_package_id BIGINT,
    CONSTRAINT orders_pkey PRIMARY KEY (id)
  );

CREATE TABLE IF NOT EXISTS
  organizers (user_id bigserial NOT NULL, CONSTRAINT organizers_pkey PRIMARY KEY (user_id));

CREATE TABLE IF NOT EXISTS
  package_addon_images (item_id BIGINT, image_id BIGINT);

CREATE TABLE IF NOT EXISTS
  package_addons (item_id bigserial NOT NULL, category TEXT, CONSTRAINT package_addons_pkey PRIMARY KEY (item_id));

CREATE TABLE IF NOT EXISTS
  price_list_items (
    id bigserial NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    price_list_id BIGINT,
    item_id BIGINT,
    date_from TIMESTAMP WITH TIME ZONE,
    date_to TIMESTAMP WITH TIME ZONE,
    is_fixed BOOLEAN,
    price NUMERIC,
    CONSTRAINT price_list_items_pkey PRIMARY KEY (id)
  );

CREATE TABLE IF NOT EXISTS
  price_lists (
    id bigserial NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    festival_id BIGINT,
    CONSTRAINT price_lists_pkey PRIMARY KEY (id)
  );

CREATE TABLE IF NOT EXISTS
  ticket_types (item_id bigserial NOT NULL, CONSTRAINT ticket_types_pkey PRIMARY KEY (item_id));

CREATE TABLE IF NOT EXISTS
  transport_addons (
    item_id bigserial NOT NULL,
    transport_type TEXT,
    departure_time TIMESTAMP WITH TIME ZONE,
    arrival_time TIMESTAMP WITH TIME ZONE,
    return_departure_time TIMESTAMP WITH TIME ZONE,
    return_arrival_time TIMESTAMP WITH TIME ZONE,
    departure_city_id BIGINT,
    arrival_city_id BIGINT,
    CONSTRAINT transport_addons_pkey PRIMARY KEY (item_id)
  );

CREATE TABLE IF NOT EXISTS
  user_profiles (
    id bigserial NOT NULL,
    first_name TEXT,
    last_name TEXT,
    date_of_birth TIMESTAMP WITH TIME ZONE,
    phone_number TEXT,
    user_id BIGINT,
    address_id BIGINT,
    image_id BIGINT,
    CONSTRAINT user_profiles_pkey PRIMARY KEY (id),
    CONSTRAINT uni_user_profiles_user_id UNIQUE (user_id)
  );

CREATE TABLE IF NOT EXISTS
  users (
    id bigserial NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    username TEXT,
    email TEXT,
    PASSWORD TEXT,
    ROLE TEXT,
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT uni_users_email UNIQUE (email),
    CONSTRAINT uni_users_username UNIQUE (username)
  );

ALTER TABLE IF EXISTS activation_help_requests
ADD CONSTRAINT fk_activation_help_requests_attendee FOREIGN KEY (attendee_id) REFERENCES attendees (user_id);

ALTER TABLE IF EXISTS activation_help_requests
ADD CONSTRAINT fk_activation_help_requests_bracelet FOREIGN KEY (bracelet_id) REFERENCES bracelets (id);

ALTER TABLE IF EXISTS activation_help_requests
ADD CONSTRAINT fk_activation_help_requests_employee FOREIGN KEY (employee_id) REFERENCES employees (user_id);

ALTER TABLE IF EXISTS activation_help_requests
ADD CONSTRAINT fk_activation_help_requests_proof_image FOREIGN KEY (proof_image_id) REFERENCES images (id);

ALTER TABLE IF EXISTS addresses
ADD CONSTRAINT fk_addresses_city FOREIGN KEY (city_id) REFERENCES cities (id);

ALTER TABLE IF EXISTS administrators
ADD CONSTRAINT fk_administrators_user FOREIGN KEY (user_id) REFERENCES users (id);

-- CREATE INDEX IF NOT EXISTS administrators_pkey ON administrators (user_id);
ALTER TABLE IF EXISTS attendees
ADD CONSTRAINT fk_attendees_user FOREIGN KEY (user_id) REFERENCES users (id);

-- CREATE INDEX IF NOT EXISTS attendees_pkey ON attendees (user_id);
ALTER TABLE IF EXISTS bracelets
ADD CONSTRAINT fk_bracelets_attendee FOREIGN KEY (attendee_id) REFERENCES attendees (user_id);

ALTER TABLE IF EXISTS bracelets
ADD CONSTRAINT fk_bracelets_employee FOREIGN KEY (employee_id) REFERENCES employees (user_id);

ALTER TABLE IF EXISTS bracelets
ADD CONSTRAINT fk_bracelets_festival_ticket FOREIGN KEY (festival_ticket_id) REFERENCES festival_tickets (id);

ALTER TABLE IF EXISTS camp_addons
ADD CONSTRAINT fk_camp_addons_item FOREIGN KEY (item_id) REFERENCES package_addons (item_id);

-- CREATE INDEX IF NOT EXISTS camp_addons_pkey ON camp_addons (item_id);
ALTER TABLE IF EXISTS camp_equipments
ADD CONSTRAINT fk_camp_equipments_item FOREIGN KEY (item_id) REFERENCES camp_addons (item_id);

ALTER TABLE IF EXISTS cities
ADD CONSTRAINT fk_cities_country FOREIGN KEY (country_id) REFERENCES countries (id);

ALTER TABLE IF EXISTS employees
ADD CONSTRAINT fk_employees_user FOREIGN KEY (user_id) REFERENCES users (id);

-- CREATE INDEX IF NOT EXISTS employees_pkey ON employees (user_id);
ALTER TABLE IF EXISTS festival_employees
ADD CONSTRAINT fk_festival_employees_festival FOREIGN KEY (festival_id) REFERENCES festivals (id);

ALTER TABLE IF EXISTS festival_employees
ADD CONSTRAINT fk_festival_employees_user FOREIGN KEY (user_id) REFERENCES employees (user_id);

ALTER TABLE IF EXISTS festival_images
ADD CONSTRAINT fk_festival_images_festival FOREIGN KEY (festival_id) REFERENCES festivals (id);

ALTER TABLE IF EXISTS festival_images
ADD CONSTRAINT fk_festival_images_image FOREIGN KEY (image_id) REFERENCES images (id);

ALTER TABLE IF EXISTS festival_organizers
ADD CONSTRAINT fk_festival_organizers_festival FOREIGN KEY (festival_id) REFERENCES festivals (id);

ALTER TABLE IF EXISTS festival_organizers
ADD CONSTRAINT fk_festival_organizers_user FOREIGN KEY (user_id) REFERENCES organizers (user_id);

ALTER TABLE IF EXISTS festival_package_addons
ADD CONSTRAINT fk_festival_package_addons_festival_package FOREIGN KEY (festival_package_id) REFERENCES festival_packages (id);

ALTER TABLE IF EXISTS festival_package_addons
ADD CONSTRAINT fk_festival_package_addons_item FOREIGN KEY (item_id) REFERENCES package_addons (item_id);

ALTER TABLE IF EXISTS festival_tickets
ADD CONSTRAINT fk_festival_tickets_item FOREIGN KEY (item_id) REFERENCES ticket_types (item_id);

ALTER TABLE IF EXISTS festivals
ADD CONSTRAINT fk_festivals_address FOREIGN KEY (address_id) REFERENCES addresses (id);

ALTER TABLE IF EXISTS items
ADD CONSTRAINT fk_items_festival FOREIGN KEY (festival_id) REFERENCES festivals (id);

ALTER TABLE IF EXISTS logs
ADD CONSTRAINT fk_logs_user FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE IF EXISTS orders
ADD CONSTRAINT fk_orders_festival_package FOREIGN KEY (festival_package_id) REFERENCES festival_packages (id);

ALTER TABLE IF EXISTS orders
ADD CONSTRAINT fk_orders_festival_ticket FOREIGN KEY (festival_ticket_id) REFERENCES festival_tickets (id);

ALTER TABLE IF EXISTS orders
ADD CONSTRAINT fk_orders_user FOREIGN KEY (user_id) REFERENCES attendees (user_id);

ALTER TABLE IF EXISTS organizers
ADD CONSTRAINT fk_organizers_user FOREIGN KEY (user_id) REFERENCES users (id);

-- CREATE INDEX IF NOT EXISTS organizers_pkey ON organizers (user_id);
ALTER TABLE IF EXISTS package_addon_images
ADD CONSTRAINT fk_package_addon_images_image FOREIGN KEY (image_id) REFERENCES images (id);

ALTER TABLE IF EXISTS package_addon_images
ADD CONSTRAINT fk_package_addon_images_item FOREIGN KEY (item_id) REFERENCES package_addons (item_id);

ALTER TABLE IF EXISTS package_addons
ADD CONSTRAINT fk_package_addons_item FOREIGN KEY (item_id) REFERENCES items (id);

-- CREATE INDEX IF NOT EXISTS package_addons_pkey ON package_addons (item_id);
ALTER TABLE IF EXISTS price_list_items
ADD CONSTRAINT fk_price_list_items_item FOREIGN KEY (item_id) REFERENCES items (id);

ALTER TABLE IF EXISTS price_list_items
ADD CONSTRAINT fk_price_list_items_price_list FOREIGN KEY (price_list_id) REFERENCES price_lists (id);

ALTER TABLE IF EXISTS price_lists
ADD CONSTRAINT fk_price_lists_festival FOREIGN KEY (festival_id) REFERENCES festivals (id);

ALTER TABLE IF EXISTS ticket_types
ADD CONSTRAINT fk_ticket_types_item FOREIGN KEY (item_id) REFERENCES items (id);

-- CREATE INDEX IF NOT EXISTS ticket_types_pkey ON ticket_types (item_id);
ALTER TABLE IF EXISTS transport_addons
ADD CONSTRAINT fk_transport_addons_arrival_city FOREIGN KEY (arrival_city_id) REFERENCES cities (id);

ALTER TABLE IF EXISTS transport_addons
ADD CONSTRAINT fk_transport_addons_departure_city FOREIGN KEY (departure_city_id) REFERENCES cities (id);

ALTER TABLE IF EXISTS transport_addons
ADD CONSTRAINT fk_transport_addons_item FOREIGN KEY (item_id) REFERENCES package_addons (item_id);

-- CREATE INDEX IF NOT EXISTS transport_addons_pkey ON transport_addons (item_id);
ALTER TABLE IF EXISTS user_profiles
ADD CONSTRAINT fk_user_profiles_address FOREIGN KEY (address_id) REFERENCES addresses (id);

ALTER TABLE IF EXISTS user_profiles
ADD CONSTRAINT fk_user_profiles_image FOREIGN KEY (image_id) REFERENCES images (id);

ALTER TABLE IF EXISTS user_profiles
ADD CONSTRAINT fk_user_profiles_user FOREIGN KEY (user_id) REFERENCES users (id);

-- CREATE INDEX IF NOT EXISTS uni_user_profiles_user_id ON user_profiles (user_id);