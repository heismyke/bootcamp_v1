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
  "careers" jsonb NOT NULL,
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
  "user_id" bigserial NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

COMMENT ON COLUMN "users"."email" IS 'Must be a valid email';

COMMENT ON COLUMN "users"."role" IS 'Allowed values: user, publisher';

COMMENT ON COLUMN "bootcamps"."website" IS 'Must start with http:// or https://';

COMMENT ON COLUMN "bootcamps"."phone" IS 'Use E.164 format (+1234567890)';

COMMENT ON COLUMN "bootcamps"."careers" IS 'List of career paths offered';

COMMENT ON COLUMN "courses"."minimum_skill" IS 'Allowed values: beginner, intermediate, advanced';

ALTER TABLE "bootcamps" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "courses" ADD FOREIGN KEY ("bootcamp_id") REFERENCES "bootcamps" ("id");

ALTER TABLE "courses" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
