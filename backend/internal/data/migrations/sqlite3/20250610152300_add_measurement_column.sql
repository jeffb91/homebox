-- +goose Up

-- STAP 1: Add "measurement" column to "maintenance_entries"
ALTER TABLE maintenance_entries ADD COLUMN measurement TEXT;

-- STAP 2: Create "attachments" table
CREATE TABLE IF NOT EXISTS "maintenance_entry_attachments" (
  "id" TEXT NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "type" TEXT NOT NULL DEFAULT 'attachment',
  "primary" BOOLEAN NOT NULL DEFAULT false,
  "maintenance_entry_id" TEXT NOT NULL,
  "filename" TEXT NOT NULL,
  "filepath" TEXT NOT NULL,
  PRIMARY KEY ("id"),
  FOREIGN KEY ("maintenance_entry_id") REFERENCES "maintenance_entries" ("id") ON DELETE CASCADE
);