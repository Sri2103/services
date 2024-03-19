-- Create "roles" table
CREATE TABLE "roles" ("id" uuid NOT NULL, "role" character varying NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, PRIMARY KEY ("id"));
-- Modify "users" table
ALTER TABLE "users" ADD COLUMN "role_user" uuid NULL, ADD CONSTRAINT "users_roles_user" FOREIGN KEY ("role_user") REFERENCES "roles" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
