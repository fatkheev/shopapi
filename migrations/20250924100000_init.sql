-- +goose Up
create extension if not exists pgcrypto;

create table if not exists address (
  id uuid primary key default gen_random_uuid(),
  country text not null,
  city text not null,
  street text not null,
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now()
);

create table if not exists client (
  id uuid primary key default gen_random_uuid(),
  client_name    text not null,
  client_surname text not null,
  birthday       date not null,
  gender smallint not null check (gender in (0,1,2)),
  registration_date timestamptz not null default now(),
  address_id uuid not null references address(id) on delete restrict,
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now()
);
create index if not exists idx_client_name_surname on client (client_surname, client_name);
create index if not exists idx_client_address_id on client (address_id);

create table if not exists supplier (
  id uuid primary key default gen_random_uuid(),
  name text not null,
  address_id uuid not null references address(id) on delete restrict,
  phone_number text not null check (phone_number ~ '^\+?[0-9\s\-\(\)]{7,20}$'),
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now()
);
create unique index if not exists uq_supplier_name_lower on supplier (lower(name));
create index if not exists idx_supplier_address_id on supplier (address_id);

create table if not exists images (
  id uuid primary key default gen_random_uuid(),
  image bytea not null,
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now()
);

create table if not exists product (
  id uuid primary key default gen_random_uuid(),
  name     text not null,
  category text not null,
  price numeric(12,2) not null check (price >= 0),
  available_stock integer not null default 0 check (available_stock >= 0),
  last_update_date date not null,
  supplier_id uuid not null references supplier(id) on delete restrict,
  image_id uuid null references images(id) on delete set null,
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now()
);
create index if not exists idx_product_supplier_id on product (supplier_id);
create index if not exists idx_product_category on product (category);
create index if not exists idx_product_available_stock on product (available_stock);

-- +goose Down
drop table if exists product;
drop table if exists images;
drop table if exists supplier;
drop table if exists client;
drop table if exists address;
