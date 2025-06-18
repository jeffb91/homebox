-- +goose Up
-- Add polymorphic fields to attachments table
ALTER TABLE "attachments" ADD COLUMN "related_type" TEXT NOT NULL DEFAULT 'items';
ALTER TABLE "attachments" ADD COLUMN "related_id" TEXT;

-- Update existing records to link to items via the current foreign key
UPDATE "attachments" 
SET "related_id" = "item_attachments"
WHERE "item_attachments" IS NOT NULL;

-- Create indexes for performance
CREATE INDEX IF NOT EXISTS "idx_attachments_related" ON "attachments"("related_type", "related_id");
CREATE INDEX IF NOT EXISTS "idx_attachments_type" ON "attachments"("related_type");

-- +goose Down
-- Remove the added columns and indexes
DROP INDEX IF EXISTS "idx_attachments_related";
DROP INDEX IF EXISTS "idx_attachments_type";
-- Note: SQLite doesn't support DROP COLUMN, so Down migration is limited