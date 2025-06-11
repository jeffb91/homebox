-- +goose Up

ALTER TABLE "maintenance_entries" ADD COLUMN "measurement" character varying NULL DEFAULT '';

