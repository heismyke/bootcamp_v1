
CREATE TYPE "user_role" AS ENUM (
  'user',
  'publisher'
);

CREATE TYPE "minimum_skill" AS ENUM (
  'beginner',
  'intermediate',
  'advanced'
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "role" user_role NOT NULL DEFAULT 'user',
  "password" varchar NOT NULL,
  "reset_password_token" varchar NOT NULL,
  "reset_password_expire" timestamp NOT NULL,
  "confirm_email_token" varchar NOT NULL,
  "is_email_confirmed" boolean NOT NULL DEFAULT false,
  "two_factor_code" varchar NOT NULL,
  "two_factor_code_expire" timestamp NOT NULL,
  "two_factor_enable" boolean NOT NULL DEFAULT false,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "bootcamps" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigserial,
  "name" varchar NOT NULL,
  "slug" varchar UNIQUE,
  "description" text NOT NULL,
  "website" varchar NOT NULL,
  "phone" varchar(20) NOT NULL,
  "email" varchar NOT NULL,
  "address" varchar NOT NULL,
  "latitude" numeric(9,6) NOT NULL,
  "longitude" numeric(9,6) NOT NULL,
  "location_details" jsonb NOT NULL,
  "careers" jsonb NOT NULL,
  "average_rating" numeric(2,1) NOT NULL DEFAULT 1,
  "average_cost" numeric NOT NULL,
  "photo" varchar NOT NULL DEFAULT 'no-photo.jpg',
  "housing" boolean NOT NULL DEFAULT false,
  "job_assistance" boolean NOT NULL DEFAULT false,
  "job_guarantee" boolean NOT NULL DEFAULT false,
  "accept_gi" boolean NOT NULL DEFAULT false,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "courses" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "description" text NOT NULL,
  "weeks" varchar NOT NULL,
  "tuition" numeric NOT NULL,
  "minimum_skill" minimum_skill NOT NULL,
  "scholarship_available" boolean NOT NULL DEFAULT false,
  "bootcamp_id" bigserial NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

COMMENT ON COLUMN "users"."email" IS 'Must be a valid email';

COMMENT ON COLUMN "users"."role" IS 'Allowed values: user, publisher';

COMMENT ON COLUMN "bootcamps"."website" IS 'Must start with http:// or https://';

COMMENT ON COLUMN "bootcamps"."phone" IS 'Use E.164 format (+1234567890)';

COMMENT ON COLUMN "bootcamps"."latitude" IS 'Latitude value (-90 to 90)';

COMMENT ON COLUMN "bootcamps"."longitude" IS 'Longitude value (-180 to 180)';

COMMENT ON COLUMN "bootcamps"."location_details" IS 'Stores city, state, country';

COMMENT ON COLUMN "bootcamps"."careers" IS 'List of career paths offered';

COMMENT ON COLUMN "bootcamps"."average_rating" IS 'Min: 1, Max: 10';

COMMENT ON COLUMN "courses"."minimum_skill" IS 'Allowed values: beginner, intermediate, advanced';

ALTER TABLE "bootcamps" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "courses" ADD FOREIGN KEY ("bootcamp_id") REFERENCES "bootcamps" ("id");
