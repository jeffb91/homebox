-- +goose Up
-- Add polymorphic fields to attachments table
ALTER TABLE "attachments" 
ADD COLUMN "related_type" character varying(50) NOT NULL DEFAULT 'items',
ADD COLUMN "related_id" uuid;

-- Update existing records to link to items via the current foreign key
UPDATE "attachments" 
SET "related_id" = "item_attachments"
WHERE "item_attachments" IS NOT NULL;

-- Create indexes for performance
CREATE INDEX IF NOT EXISTS "idx_attachments_related" ON "attachments"("related_type", "related_id");
CREATE INDEX IF NOT EXISTS "idx_attachments_type" ON "attachments"("related_type");

-- Add constraint for valid related_type values
ALTER TABLE "attachments" 
ADD CONSTRAINT "chk_attachments_related_type" 
CHECK ("related_type" IN ('items', 'maintenance_entries', 'incidents', 'reports'));

-- +goose Down
-- Remove the added columns and constraints
ALTER TABLE "attachments" DROP CONSTRAINT IF EXISTS "chk_attachments_related_type";
DROP INDEX IF EXISTS "idx_attachments_related";
DROP INDEX IF EXISTS "idx_attachments_type";
ALTER TABLE "attachments" DROP COLUMN IF EXISTS "related_id";
ALTER TABLE "attachments" DROP COLUMN IF EXISTS "related_type";