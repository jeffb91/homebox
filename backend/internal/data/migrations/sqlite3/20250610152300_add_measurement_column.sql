-- +goose Up

-- STAP 1: Add "measurement" column to "maintenance_entries"
ALTER TABLE maintenance_entries ADD COLUMN measurement TEXT;
