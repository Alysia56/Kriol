--Filename: 000002_add_created_at_to_kriol_table.up.sql
ALTER TABLE entries
ADD created_at timestamp(0) WITH time zone NOT NULL DEFAULT NOW();