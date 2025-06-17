-- +goose Up

-- 1. Voeg "measurement" column toe aan "maintenance_entries"
ALTER TABLE "maintenance_entries" ADD COLUMN "measurement" TEXT;

-- 2. Maak een tijdelijke tabel met de nieuwe polymorfe structuur
CREATE TABLE "attachments_tmp" (
    "id"              uuid        NOT NULL PRIMARY KEY,
    "created_at"      datetime    NOT NULL,
    "updated_at"      datetime    NOT NULL,
    "type"            text        NOT NULL DEFAULT 'attachment',
    "primary"       bool        NOT NULL DEFAULT false,
    "path"            text        NOT NULL,
    "title"           text        NOT NULL,
    "related_type"    text        NOT NULL,
    "related_id"      text        NOT NULL
);

-- 3. Kopieer bestaande data naar de tijdelijke tabel
--    Zet bestaande relaties om naar polymorf (voorbeeld: alles wat een item_attachments heeft wordt 'item')
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
-- UNION ALL
-- SELECT ..., 'document', document_attachments FROM attachments WHERE document_attachments IS NOT NULL

-- 4. Drop de oude attachments-tabel
DROP TABLE "attachments";

-- 5. Hernoem de tijdelijke tabel naar attachments
ALTER TABLE "attachments_tmp" RENAME TO "attachments";

-- 6. (Optioneel) Index voor snellere lookups
CREATE INDEX IF NOT EXISTS "idx_attachments_related" ON "attachments" ("related_type", "related_id");

