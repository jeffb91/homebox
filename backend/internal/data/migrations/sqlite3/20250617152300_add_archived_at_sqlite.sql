-- +goose Up
ALTER TABLE "items"
ADD COLUMN "archived_at" DATETIME;
