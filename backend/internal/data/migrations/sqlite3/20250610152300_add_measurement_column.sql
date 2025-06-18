-- +goose Up

-- 1. Voeg "measurement" column toe aan "maintenance_entries"
ALTER TABLE maintenance_entries ADD COLUMN measurement TEXT;

