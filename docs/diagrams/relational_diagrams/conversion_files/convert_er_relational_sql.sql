CREATE TABLE user
(
  user_id INT NOT NULL,
  username INT NOT NULL,
  password INT NOT NULL,
  role INT NOT NULL,
  created_at INT NOT NULL,
  PRIMARY KEY (user_id),
  UNIQUE (username)
);

CREATE TABLE user_profile
(
  user_profile_id INT NOT NULL,
  first_name INT NOT NULL,
  last_name INT NOT NULL,
  email INT NOT NULL,
  date_of_birth INT NOT NULL,
  phone_number INT NOT NULL,
  user_id INT NOT NULL,
  PRIMARY KEY (user_id),
  FOREIGN KEY (user_id) REFERENCES user(user_id)
);

CREATE TABLE attendee
(
  user_id INT NOT NULL,
  PRIMARY KEY (user_id),
  FOREIGN KEY (user_id) REFERENCES user(user_id)
);

CREATE TABLE employee
(
  user_id INT NOT NULL,
  PRIMARY KEY (user_id),
  FOREIGN KEY (user_id) REFERENCES user(user_id)
);

CREATE TABLE organizer
(
  user_id INT NOT NULL,
  PRIMARY KEY (user_id),
  FOREIGN KEY (user_id) REFERENCES user(user_id)
);

CREATE TABLE log
(
  log_id INT NOT NULL,
  type INT NOT NULL,
  timestamp INT NOT NULL,
  description INT,
  data INT,
  user_id INT,
  PRIMARY KEY (log_id),
  FOREIGN KEY (user_id) REFERENCES user(user_id)
);

CREATE TABLE country
(
  country_id INT NOT NULL,
  name INT NOT NULL,
  iso_code_2 INT NOT NULL,
  PRIMARY KEY (country_id),
  UNIQUE (iso_code_2)
);

CREATE TABLE city
(
  city_id INT NOT NULL,
  name INT NOT NULL,
  postal_code INT NOT NULL,
  country_id INT NOT NULL,
  PRIMARY KEY (city_id),
  FOREIGN KEY (country_id) REFERENCES country(country_id)
);

CREATE TABLE address
(
  address_id INT NOT NULL,
  street INT NOT NULL,
  number INT NOT NULL,
  apartment_suite INT NOT NULL,
  user_id INT,
  city_id INT NOT NULL,
  PRIMARY KEY (address_id),
  FOREIGN KEY (user_id) REFERENCES user_profile(user_id),
  FOREIGN KEY (city_id) REFERENCES city(city_id)
);

CREATE TABLE festival
(
  name INT NOT NULL,
  description INT NOT NULL,
  start_date INT NOT NULL,
  end_date INT NOT NULL,
  capacity INT NOT NULL,
  status INT NOT NULL,
  store_status INT NOT NULL,
  festival_id INT NOT NULL,
  address_id INT NOT NULL,
  PRIMARY KEY (festival_id),
  FOREIGN KEY (address_id) REFERENCES address(address_id)
);

CREATE TABLE pricelist
(
  pricelist_id INT NOT NULL,
  festival_id INT NOT NULL,
  PRIMARY KEY (pricelist_id),
  FOREIGN KEY (festival_id) REFERENCES festival(festival_id)
);

CREATE TABLE item
(
  name INT NOT NULL,
  description INT NOT NULL,
  available_number INT NOT NULL,
  remaining_number INT NOT NULL,
  item_id INT NOT NULL,
  type INT NOT NULL,
  festival_id INT NOT NULL,
  PRIMARY KEY (item_id),
  FOREIGN KEY (festival_id) REFERENCES festival(festival_id)
);

CREATE TABLE ticket_type
(
  item_id INT NOT NULL,
  PRIMARY KEY (item_id),
  FOREIGN KEY (item_id) REFERENCES item(item_id)
);

CREATE TABLE package_addon
(
  category INT NOT NULL,
  item_id INT NOT NULL,
  PRIMARY KEY (item_id),
  FOREIGN KEY (item_id) REFERENCES item(item_id)
);

CREATE TABLE camp_addon
(
  camp_name INT NOT NULL,
  item_id INT NOT NULL,
  PRIMARY KEY (item_id),
  FOREIGN KEY (item_id) REFERENCES package_addon(item_id)
);

CREATE TABLE transport_addon
(
  transport_type INT NOT NULL,
  departure_time INT NOT NULL,
  arrival_time INT NOT NULL,
  return_departure_time INT NOT NULL,
  return_arrival_time INT NOT NULL,
  item_id INT NOT NULL,
  city_id INT NOT NULL,
  departures_fromcity_id INT NOT NULL,
  PRIMARY KEY (item_id),
  FOREIGN KEY (item_id) REFERENCES package_addon(item_id),
  FOREIGN KEY (city_id) REFERENCES city(city_id),
  FOREIGN KEY (departures_fromcity_id) REFERENCES city(city_id)
);

CREATE TABLE camp_equipment
(
  camp_equipment_id INT NOT NULL,
  name INT NOT NULL,
  item_id INT NOT NULL,
  PRIMARY KEY (camp_equipment_id),
  FOREIGN KEY (item_id) REFERENCES camp_addon(item_id)
);

CREATE TABLE festival_ticket
(
  festival_ticket_id INT NOT NULL,
  item_id INT NOT NULL,
  PRIMARY KEY (festival_ticket_id),
  FOREIGN KEY (item_id) REFERENCES ticket_type(item_id)
);

CREATE TABLE bracelet
(
  bracelet_id INT NOT NULL,
  pin INT NOT NULL,
  barcode_number INT NOT NULL,
  balance INT NOT NULL,
  festival_ticket_id INT NOT NULL,
  attendee_id INT NOT NULL,
  employee_id INT,
  PRIMARY KEY (bracelet_id),
  FOREIGN KEY (festival_ticket_id) REFERENCES festival_ticket(festival_ticket_id),
  FOREIGN KEY (attendee_id) REFERENCES attendee(user_id),
  FOREIGN KEY (employee_id) REFERENCES employee(user_id)
);

CREATE TABLE activation_help_request
(
  activation_help_request_id INT NOT NULL,
  user_entered_pin INT NOT NULL,
  issue_description INT NOT NULL,
  timestamp INT NOT NULL,
  status INT NOT NULL,
  bracelet_id INT NOT NULL,
  attendee_id INT NOT NULL,
  employee_id INT,
  PRIMARY KEY (activation_help_request_id),
  FOREIGN KEY (bracelet_id) REFERENCES bracelet(bracelet_id),
  FOREIGN KEY (attendee_id) REFERENCES attendee(user_id),
  FOREIGN KEY (employee_id) REFERENCES employee(user_id)
);

CREATE TABLE works_at
(
  user_id INT NOT NULL,
  festival_id INT NOT NULL,
  PRIMARY KEY (user_id, festival_id),
  FOREIGN KEY (user_id) REFERENCES employee(user_id),
  FOREIGN KEY (festival_id) REFERENCES festival(festival_id)
);

CREATE TABLE organizes
(
  user_id INT NOT NULL,
  festival_id INT NOT NULL,
  PRIMARY KEY (user_id, festival_id),
  FOREIGN KEY (user_id) REFERENCES organizer(user_id),
  FOREIGN KEY (festival_id) REFERENCES festival(festival_id)
);

CREATE TABLE image
(
  image_id INT NOT NULL,
  url INT NOT NULL,
  user_id INT,
  activation_help_request_id INT,
  PRIMARY KEY (image_id),
  FOREIGN KEY (user_id) REFERENCES user_profile(user_id),
  FOREIGN KEY (activation_help_request_id) REFERENCES activation_help_request(activation_help_request_id)
);

CREATE TABLE pricelist_item
(
  date_from INT NOT NULL,
  date_to INT NOT NULL,
  is_fixed INT NOT NULL,
  pricelist_item_id INT NOT NULL,
  pricelist_id INT NOT NULL,
  item_id INT NOT NULL,
  PRIMARY KEY (pricelist_item_id),
  FOREIGN KEY (pricelist_id) REFERENCES pricelist(pricelist_id),
  FOREIGN KEY (item_id) REFERENCES item(item_id)
);

CREATE TABLE order
(
  order_id INT NOT NULL,
  timestamp INT NOT NULL,
  total_amount INT NOT NULL,
  user_id INT NOT NULL,
  festival_ticket_id INT NOT NULL,
  PRIMARY KEY (order_id),
  FOREIGN KEY (user_id) REFERENCES attendee(user_id),
  FOREIGN KEY (festival_ticket_id) REFERENCES festival_ticket(festival_ticket_id)
);

CREATE TABLE festival_package
(
  festival_package_id INT NOT NULL,
  order_id INT,
  PRIMARY KEY (festival_package_id),
  FOREIGN KEY (order_id) REFERENCES order(order_id)
);

CREATE TABLE has_images
(
  festival_id INT NOT NULL,
  image_id INT NOT NULL,
  PRIMARY KEY (festival_id, image_id),
  FOREIGN KEY (festival_id) REFERENCES festival(festival_id),
  FOREIGN KEY (image_id) REFERENCES image(image_id)
);

CREATE TABLE illustrated_by
(
  item_id INT NOT NULL,
  image_id INT NOT NULL,
  PRIMARY KEY (item_id, image_id),
  FOREIGN KEY (item_id) REFERENCES package_addon(item_id),
  FOREIGN KEY (image_id) REFERENCES image(image_id)
);

CREATE TABLE bundles
(
  festival_package_id INT NOT NULL,
  item_id INT NOT NULL,
  PRIMARY KEY (festival_package_id, item_id),
  FOREIGN KEY (festival_package_id) REFERENCES festival_package(festival_package_id),
  FOREIGN KEY (item_id) REFERENCES package_addon(item_id)
);