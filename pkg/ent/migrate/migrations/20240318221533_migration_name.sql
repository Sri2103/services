-- Modify "users" table
ALTER TABLE "users" ALTER COLUMN "password" TYPE character varying, ADD COLUMN "email" character varying NOT NULL;
-- Create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX "users_email_key" ON "users" ("email");
