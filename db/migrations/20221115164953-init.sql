-- +migrate Up
CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "user_id" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "username" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "sex" varchar NOT NULL,
  "data_of_birth" varchar NOT NULL,
  "destination_family_name" varchar NOT NULL,
  "destination_first_name" varchar NOT NULL,
  "postcode" integer NOT NULL,
  "prefecture_code" varchar NOT NULL,
  "city" varchar NOT NULL,
  "street" varchar NOT NULL,
  "building" varchar NOT NULL,
  "phone" varchar UNIQUE NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL
);

CREATE TABLE "products" (
  "id" bigserial PRIMARY KEY,
  "product_name" varchar NOT NULL,
  "description" text,
  "img" text,
  "unit_price" integer NOT NULL,
  "discount" float8 NOT NULL DEFAULT 0,
  "stock" integer NOT NULL DEFAULT 0,
  "brand_id" bigserial NOT NULL,
  "category" bigserial NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL
);

CREATE TABLE "categories" (
  "id" bigserial PRIMARY KEY,
  "category" varchar UNIQUE NOT NULL
);

CREATE TABLE "brands" (
  "id" bigserial PRIMARY KEY,
  "brand_name" varchar UNIQUE NOT NULL
);

CREATE TABLE "carts" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigserial NOT NULL,
  "product_id" bigserial NOT NULL
);

CREATE TABLE "orders" (
  "id" bigserial PRIMARY KEY,
  "user_id" varchar NOT NULL,
  "product_id" bigserial NOT NULL,
  "quantity" integer NOT NULL,
  "postage" integer NOT NULL,
  "price" integer NOT NULL DEFAULT 0,
  "status" boolean NOT NULL DEFAULT false,
  "created_at" timestamp NOT NULL
);

CREATE TABLE "media" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "contents" text NOT NULL,
  "img" varchar NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL
);

CREATE TABLE "adminuser" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "password" varchar UNIQUE NOT NULL,
  "created_at" timestamp NOT NULL
);

ALTER TABLE "products" ADD FOREIGN KEY ("brand_id") REFERENCES "brands" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("category") REFERENCES "categories" ("id");

ALTER TABLE "carts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "carts" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "orders" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");


-- +migrate Down
DROP TABLE IF EXISTS "users" cascade;
DROP TABLE IF EXISTS "address" cascade;
DROP TABLE IF EXISTS "products" cascade;
DROP TABLE IF EXISTS "categories" cascade;
DROP TABLE IF EXISTS "brands" cascade;
DROP TABLE IF EXISTS "carts" cascade;
DROP TABLE IF EXISTS "orders" cascade;
DROP TABLE IF EXISTS "media" cascade;
