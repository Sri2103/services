 -- Create "users" table
CREATE TABLE "users" ("_id" uuid NOT NULL, "name" character varying NOT NULL, "username" character varying NOT NULL, "password" bytea NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, PRIMARY KEY ("_id"));
-- Create "addresses" table
CREATE TABLE "addresses" ("id" uuid NOT NULL, "street_address" character varying NOT NULL, "city" character varying NOT NULL, "state" character varying NOT NULL, "country" character varying NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "user_addresses" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "addresses_users_addresses" FOREIGN KEY ("user_addresses") REFERENCES "users" ("_id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create "carts" table
CREATE TABLE "carts" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "user_carts" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "carts_users_carts" FOREIGN KEY ("user_carts") REFERENCES "users" ("_id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create "cart_items" table
CREATE TABLE "cart_items" ("id" uuid NOT NULL, "quantity" bigint NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "cart_items" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "cart_items_carts_items" FOREIGN KEY ("cart_items") REFERENCES "carts" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create "categories" table
CREATE TABLE "categories" ("id" uuid NOT NULL, "name" character varying NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, PRIMARY KEY ("id"));
-- Create index "categories_name_key" to table: "categories"
CREATE UNIQUE INDEX "categories_name_key" ON "categories" ("name");
-- Create "orders" table
CREATE TABLE "orders" ("id" uuid NOT NULL, "total_amount" double precision NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "user_orders" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "orders_users_orders" FOREIGN KEY ("user_orders") REFERENCES "users" ("_id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create "order_items" table
CREATE TABLE "order_items" ("id" uuid NOT NULL, "quantity" bigint NOT NULL, "price" double precision NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "order_items" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "order_items_orders_items" FOREIGN KEY ("order_items") REFERENCES "orders" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create "products" table
CREATE TABLE "products" ("id" uuid NOT NULL, "name" character varying NOT NULL, "description" character varying NOT NULL, "price" double precision NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "cart_item_product" uuid NULL, "order_item_product" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "products_cart_items_product" FOREIGN KEY ("cart_item_product") REFERENCES "cart_items" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, CONSTRAINT "products_order_items_product" FOREIGN KEY ("order_item_product") REFERENCES "order_items" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create "category_products" table
CREATE TABLE "category_products" ("category_id" uuid NOT NULL, "product_id" uuid NOT NULL, PRIMARY KEY ("category_id", "product_id"), CONSTRAINT "category_products_category_id" FOREIGN KEY ("category_id") REFERENCES "categories" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "category_products_product_id" FOREIGN KEY ("product_id") REFERENCES "products" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
