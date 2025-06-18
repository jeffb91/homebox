-- +goose Up
--STAP 1: Add "measurement" column to "maintenance_entries
ALTER TABLE "maintenance_entries" ADD COLUMN "measurement" character varying NULL DEFAULT '';

-- Create "attachments" table
--CREATE TABLE IF NOT EXISTS "maintenance_entry_attachments" (
--  "id" TEXT NOT NULL,
--  "uploaded_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
--  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
--  "type" TEXT NOT NULL DEFAULT 'attachment',
--  "primary" BOOLEAN NOT NULL DEFAULT false,
--  "maintenance_entry_id" TEXT NOT NULL,
--  "filename" TEXT NOT NULL,
--  "filepath" TEXT NOT NULL,
--  "content_type" TEXT NULL
--  PRIMARY KEY ("id"),
--  FOREIGN KEY ("maintenance_entry_id") REFERENCES "maintenance_entries" ("id") ON DELETE CASCADE
--); 

-- 1. Voeg "measurement" column toe aan "maintenance_entries"
ALTER TABLE "maintenance_entries" ADD COLUMN IF NOT EXISTS "measurement" character varying;

-- 2. Maak een tijdelijke tabel met de nieuwe polymorfe structuur
CREATE TABLE "attachments_tmp" (
    "id"              uuid        NOT NULL PRIMARY KEY,
    "created_at"     timestamptz NOT NULL,
    "updated_at"      timestamptz NOT NULL,
    "type"            character varying NOT NULL DEFAULT 'attachment',
    "primary"       boolean     NOT NULL DEFAULT false,
    "path"            character varying NOT NULL,
    "title"          character varying NOT NULL,
    "related_type"    character varying NOT NULL,
    "related_id"      uuid        NOT NULL
);

-- 3. Kopieer bestaande data naar de tijdelijke tabel
INSERT INTO "attachments_tmp" ("id", "created_at", "updated_at", "type", "primary", "path", "title", "related_type", "related_id")
SELECT
    "id",
    "created_at",
    "updated_at",
    "type",
    "primary",
    "path",
    "title",
    'item' as "related_type",
    "item_attachments" as "related_id"
FROM attachments;

-- (Optioneel) Voeg hier extra INSERTs toe als je andere types (zoals document_attachments) wilt migreren:
-- Bijvoorbeeld:
-- INSERT INTO attachments_tmp (...) SELECT ..., 'document', document_attachments FROM attachments WHERE document_attachments IS NOT NULL;

-- 4. Drop de oude attachments-tabel
DROP TABLE "attachments";

-- 5. Hernoem de tijdelijke tabel naar attachments
ALTER TABLE "attachments_tmp" RENAME TO "attachments";

-- 6. (Optioneel) Index voor snellere lookups
CREATE INDEX IF NOT EXISTS "idx_attachments_related" ON "attachments" ("related_type", "related_id");

