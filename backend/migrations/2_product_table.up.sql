CREATE EXTENSION IF NOT EXISTS pg_trgm;

CREATE TABLE IF NOT EXISTS products (
  id uuid NOT NULL DEFAULT gen_random_uuid(),

  name TEXT NOT NULL,
  description TEXT,
  category TEXT,
  brand TEXT,
  stock_quantity BIGINT NOT NULL DEFAULT 0,
  sku TEXT,

  created_at timestamp without time zone NOT NULL DEFAULT now(),
  updated_at timestamp without time zone NOT NULL DEFAULT now(),

  "name_search" TEXT GENERATED ALWAYS AS (regexp_replace(lower("name"), '[^a-z0-9]+', '', 'g')) STORED,

  CONSTRAINT product_id_pkey PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS product_name_search_trgm ON products USING gin ("name_search" gin_trgm_ops);
