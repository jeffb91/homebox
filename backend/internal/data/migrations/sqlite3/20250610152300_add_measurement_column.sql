-- +goose Up
ALTER TABLE maintenance_entries ADD COLUMN measurement TEXT;
