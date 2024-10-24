
CREATE TABLE "user" (
  "user_id" uuid PRIMARY KEY NOT NULL,
  "username" varchar(255) UNIQUE NOT NULL,
  "email" varchar(255) UNIQUE NOT NULL,
  "password" varchar(255) NOT NULL,
  "role" varchar(50) NOT NULL,
  "created_at" timestamp NOT NULL
);

CREATE TABLE "user_profile" (
  "user_profile_id" uuid PRIMARY KEY NOT NULL,
  "first_name" varchar(255) NOT NULL,
  "last_name" varchar(255) NOT NULL,
  "date_of_birth" date NOT NULL,
  "phone_number" varchar(20) UNIQUE,
  "user_id" uuid UNIQUE NOT NULL,
  "address_id" uuid,
  "image_id" uuid
);

CREATE TABLE "attendee" (
  "user_id" uuid PRIMARY KEY NOT NULL
);

CREATE TABLE "employee" (
  "user_id" uuid PRIMARY KEY NOT NULL
);

CREATE TABLE "organizer" (
  "user_id" uuid PRIMARY KEY NOT NULL
);

CREATE TABLE "image" (
  "image_id" uuid PRIMARY KEY NOT NULL,
  "url" varchar(255) NOT NULL
);

CREATE TABLE "log" (
  "log_id" uuid PRIMARY KEY NOT NULL,
  "type" varchar(100) NOT NULL,
  "timestamp" timestamp NOT NULL,
  "description" text,
  "data" jsonb,
  "user_id" uuid
);

CREATE TABLE "country" (
  "country_id" uuid PRIMARY KEY NOT NULL,
  "name" varchar(100) NOT NULL,
  "iso_code_2" char(2) UNIQUE NOT NULL
);

CREATE TABLE "city" (
  "city_id" uuid PRIMARY KEY NOT NULL,
  "name" varchar(100) NOT NULL,
  "postal_code" varchar(20) NOT NULL,
  "country_id" uuid NOT NULL
);

CREATE TABLE "address" (
  "address_id" uuid PRIMARY KEY NOT NULL,
  "street" varchar(255) NOT NULL,
  "number" varchar(20) NOT NULL,
  "apartment_suite" varchar(100),
  "city_id" uuid NOT NULL
);

CREATE TABLE "festival" (
  "festival_id" uuid PRIMARY KEY NOT NULL,
  "name" varchar(255) NOT NULL,
  "description" text NOT NULL,
  "start_date" date NOT NULL,
  "end_date" date NOT NULL,
  "capacity" int NOT NULL,
  "status" varchar(50) NOT NULL,
  "store_status" varchar(50) NOT NULL,
  "address_id" uuid NOT NULL
);

CREATE TABLE "festival_image" (
  "festival_id" uuid NOT NULL,
  "image_id" uuid NOT NULL,
  PRIMARY KEY ("festival_id", "image_id")
);

CREATE TABLE "price_list" (
  "price_list_id" uuid PRIMARY KEY NOT NULL,
  "festival_id" uuid NOT NULL
);

CREATE TABLE "price_list_item" (
  "price_list_item_id" uuid PRIMARY KEY NOT NULL,
  "date_from" date NOT NULL,
  "date_to" date NOT NULL,
  "is_fixed" boolean NOT NULL,
  "price_list_id" uuid NOT NULL,
  "item_id" uuid NOT NULL
);

CREATE TABLE "item" (
  "item_id" uuid PRIMARY KEY NOT NULL,
  "name" varchar(255) NOT NULL,
  "description" text NOT NULL,
  "available_number" int NOT NULL,
  "remaining_number" int NOT NULL,
  "type" varchar(100) NOT NULL,
  "festival_id" uuid NOT NULL
);

CREATE TABLE "ticket_type" (
  "item_id" uuid PRIMARY KEY NOT NULL
);

CREATE TABLE "package_addon" (
  "item_id" uuid PRIMARY KEY NOT NULL,
  "category" varchar(100) NOT NULL
);

CREATE TABLE "package_addon_image" (
  "item_id" uuid NOT NULL,
  "image_id" uuid NOT NULL,
  PRIMARY KEY ("item_id", "image_id")
);

CREATE TABLE "camp_addon" (
  "item_id" uuid PRIMARY KEY NOT NULL,
  "camp_name" varchar(255) NOT NULL
);

CREATE TABLE "camp_equipment" (
  "camp_equipment_id" uuid PRIMARY KEY NOT NULL,
  "name" varchar(255) NOT NULL,
  "item_id" uuid NOT NULL
);

CREATE TABLE "transport_addon" (
  "item_id" uuid PRIMARY KEY NOT NULL,
  "transport_type" varchar(100) NOT NULL,
  "departure_time" timestamp NOT NULL,
  "arrival_time" timestamp NOT NULL,
  "return_departure_time" timestamp NOT NULL,
  "return_arrival_time" timestamp NOT NULL,
  "departure_city_id" uuid NOT NULL,
  "arrive_city_id" uuid NOT NULL
);

CREATE TABLE "festival_ticket" (
  "festival_ticket_id" uuid PRIMARY KEY NOT NULL,
  "item_id" uuid NOT NULL
);

CREATE TABLE "bracelet" (
  "bracelet_id" uuid PRIMARY KEY NOT NULL,
  "pin" varchar(20) NOT NULL,
  "barcode_number" varchar(50) UNIQUE NOT NULL,
  "balance" decimal(10, 2) NOT NULL DEFAULT 0,
  "festival_ticket_id" uuid NOT NULL,
  "attendee_id" uuid NOT NULL,
  "employee_id" uuid
);

CREATE TABLE "activation_help_request" (
  "activation_help_request_id" uuid PRIMARY KEY NOT NULL,
  "user_entered_pin" varchar(20) NOT NULL,
  "issue_description" text NOT NULL,
  "timestamp" timestamp NOT NULL,
  "proof_image_id" uuid NOT NULL,
  "status" varchar(50) NOT NULL,
  "bracelet_id" uuid NOT NULL,
  "attendee_id" uuid NOT NULL,
  "employee_id" uuid
);

CREATE TABLE "festival_employee" (
  "user_id" uuid NOT NULL,
  "festival_id" uuid NOT NULL,
  PRIMARY KEY ("user_id", "festival_id")
);

CREATE TABLE "festival_organizer" (
  "user_id" uuid NOT NULL,
  "festival_id" uuid NOT NULL,
  PRIMARY KEY ("user_id", "festival_id")
);

CREATE TABLE "order" (
  "order_id" uuid PRIMARY KEY NOT NULL,
  "timestamp" timestamp NOT NULL,
  "total_amount" decimal(10, 2) NOT NULL,
  "user_id" uuid NOT NULL,
  "festival_ticket_id" uuid NOT NULL,
  "festival_package_id" uuid
);

CREATE TABLE "festival_package" (
  "festival_package_id" uuid PRIMARY KEY NOT NULL
);

CREATE TABLE "festival_package_addon" (
  "festival_package_id" uuid NOT NULL,
  "item_id" uuid NOT NULL,
  PRIMARY KEY ("festival_package_id", "item_id")
);

ALTER TABLE "user_profile" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("user_id");

ALTER TABLE "user_profile" ADD FOREIGN KEY ("address_id") REFERENCES "address" ("address_id");

ALTER TABLE "user_profile" ADD FOREIGN KEY ("image_id") REFERENCES "image" ("image_id");

ALTER TABLE "attendee" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("user_id");

ALTER TABLE "employee" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("user_id");

ALTER TABLE "organizer" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("user_id");

ALTER TABLE "log" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("user_id");

ALTER TABLE "city" ADD FOREIGN KEY ("country_id") REFERENCES "country" ("country_id");

ALTER TABLE "address" ADD FOREIGN KEY ("city_id") REFERENCES "city" ("city_id");

ALTER TABLE "festival" ADD FOREIGN KEY ("address_id") REFERENCES "address" ("address_id");

ALTER TABLE "festival_image" ADD FOREIGN KEY ("festival_id") REFERENCES "festival" ("festival_id");

ALTER TABLE "festival_image" ADD FOREIGN KEY ("image_id") REFERENCES "image" ("image_id");

ALTER TABLE "price_list" ADD FOREIGN KEY ("festival_id") REFERENCES "festival" ("festival_id");

ALTER TABLE "price_list_item" ADD FOREIGN KEY ("price_list_id") REFERENCES "price_list" ("price_list_id");

ALTER TABLE "price_list_item" ADD FOREIGN KEY ("item_id") REFERENCES "item" ("item_id");

ALTER TABLE "item" ADD FOREIGN KEY ("festival_id") REFERENCES "festival" ("festival_id");

ALTER TABLE "ticket_type" ADD FOREIGN KEY ("item_id") REFERENCES "item" ("item_id");

ALTER TABLE "package_addon" ADD FOREIGN KEY ("item_id") REFERENCES "item" ("item_id");

ALTER TABLE "package_addon_image" ADD FOREIGN KEY ("item_id") REFERENCES "package_addon" ("item_id");

ALTER TABLE "package_addon_image" ADD FOREIGN KEY ("image_id") REFERENCES "image" ("image_id");

ALTER TABLE "camp_addon" ADD FOREIGN KEY ("item_id") REFERENCES "package_addon" ("item_id");

ALTER TABLE "camp_equipment" ADD FOREIGN KEY ("item_id") REFERENCES "camp_addon" ("item_id");

ALTER TABLE "transport_addon" ADD FOREIGN KEY ("item_id") REFERENCES "package_addon" ("item_id");

ALTER TABLE "transport_addon" ADD FOREIGN KEY ("departure_city_id") REFERENCES "city" ("city_id");

ALTER TABLE "transport_addon" ADD FOREIGN KEY ("arrive_city_id") REFERENCES "city" ("city_id");

ALTER TABLE "festival_ticket" ADD FOREIGN KEY ("item_id") REFERENCES "item" ("item_id");

ALTER TABLE "bracelet" ADD FOREIGN KEY ("festival_ticket_id") REFERENCES "festival_ticket" ("festival_ticket_id");

ALTER TABLE "bracelet" ADD FOREIGN KEY ("attendee_id") REFERENCES "attendee" ("user_id");

ALTER TABLE "bracelet" ADD FOREIGN KEY ("employee_id") REFERENCES "employee" ("user_id");

ALTER TABLE "activation_help_request" ADD FOREIGN KEY ("proof_image_id") REFERENCES "image" ("image_id");

ALTER TABLE "activation_help_request" ADD FOREIGN KEY ("bracelet_id") REFERENCES "bracelet" ("bracelet_id");

ALTER TABLE "activation_help_request" ADD FOREIGN KEY ("attendee_id") REFERENCES "attendee" ("user_id");

ALTER TABLE "activation_help_request" ADD FOREIGN KEY ("employee_id") REFERENCES "employee" ("user_id");

ALTER TABLE "festival_employee" ADD FOREIGN KEY ("user_id") REFERENCES "employee" ("user_id");

ALTER TABLE "festival_employee" ADD FOREIGN KEY ("festival_id") REFERENCES "festival" ("festival_id");

ALTER TABLE "festival_organizer" ADD FOREIGN KEY ("user_id") REFERENCES "organizer" ("user_id");

ALTER TABLE "festival_organizer" ADD FOREIGN KEY ("festival_id") REFERENCES "festival" ("festival_id");

ALTER TABLE "order" ADD FOREIGN KEY ("user_id") REFERENCES "attendee" ("user_id");

ALTER TABLE "order" ADD FOREIGN KEY ("festival_ticket_id") REFERENCES "festival_ticket" ("festival_ticket_id");

ALTER TABLE "order" ADD FOREIGN KEY ("festival_package_id") REFERENCES "festival_package" ("festival_package_id");

ALTER TABLE "festival_package_addon" ADD FOREIGN KEY ("festival_package_id") REFERENCES "festival_package" ("festival_package_id");

ALTER TABLE "festival_package_addon" ADD FOREIGN KEY ("item_id") REFERENCES "package_addon" ("item_id");


COMMENT ON COLUMN "festival"."status" IS 'active, inactive, private...';

COMMENT ON COLUMN "festival"."store_status" IS 'open, closed';

COMMENT ON COLUMN "item"."type" IS 'ticket, addon...';

COMMENT ON COLUMN "package_addon"."category" IS 'camp, transport, general';

COMMENT ON COLUMN "transport_addon"."transport_type" IS 'bus, train, plane...';

COMMENT ON COLUMN "transport_addon"."departure_city_id" IS 'departure city';

COMMENT ON COLUMN "transport_addon"."arrive_city_id" IS 'destination city';

COMMENT ON COLUMN "bracelet"."attendee_id" IS 'belongs to the Attendee';

COMMENT ON COLUMN "bracelet"."employee_id" IS 'Employee issued the Bracelet';

COMMENT ON COLUMN "activation_help_request"."status" IS 'open, accepted, rejected';
