-- 001_init.sql
-- AuthKit initial schema (Postgres)

-- Optional, aber sehr clean: eigenes Schema
-- Wenn du lieber "public" willst, kommentier die nächsten 2 Zeilen aus.
create schema if not exists authkit;
set search_path to authkit;

-- Extensions
-- gen_random_uuid() kommt aus pgcrypto
create extension if not exists pgcrypto;

-- Roles
create table if not exists userroles (
  id uuid primary key default gen_random_uuid(),
  name text not null unique
);

-- Users
create table if not exists users (
  id uuid primary key default gen_random_uuid(),
  firstname text,
  lastname text,
  username text not null,
  email text not null,
  password_hash text not null,
  password_changed_at timestamptz,
  roleid uuid not null references userroles(id),
  created_at timestamptz not null default now()
);

-- Case-insensitive unique constraints
create unique index if not exists users_email_unique_ci
  on users (lower(email));

create unique index if not exists users_username_unique_ci
  on users (lower(username));

-- Sessions
-- id ist exakt der Cookie-Sessionwert (base64url random)
create table if not exists usersessions (
  id text primary key,
  userid uuid not null references users(id) on delete cascade,
  created_at timestamptz not null default now(),
  expires_at timestamptz not null,
  last_seen_at timestamptz,
  revoked_at timestamptz,
  ip inet,
  user_agent text
);

create index if not exists usersessions_userid_idx
  on usersessions(userid);

create index if not exists usersessions_expires_idx
  on usersessions(expires_at);

-- Password reset tokens (store SHA-256(token_bytes) as BYTEA)
create table if not exists password_reset_tokens (
  id uuid primary key default gen_random_uuid(),
  userid uuid not null references users(id) on delete cascade,
  token_hash bytea not null,
  created_at timestamptz not null default now(),
  expires_at timestamptz not null,
  used_at timestamptz
);

create unique index if not exists prt_token_hash_unique
  on password_reset_tokens(token_hash);

create index if not exists prt_userid_idx
  on password_reset_tokens(userid);

create index if not exists prt_expires_idx
  on password_reset_tokens(expires_at);

-- Audit log
create table if not exists audit_log (
  id uuid primary key default gen_random_uuid(),
  userid uuid references users(id) on delete set null,
  event_type text not null,
  created_at timestamptz not null default now(),
  ip inet,
  user_agent text,
  metadata jsonb
);

create index if not exists audit_created_idx
  on audit_log(created_at);

create index if not exists audit_event_idx
  on audit_log(event_type);

-- Seed roles (idempotent)
insert into userroles (name) values ('admin')
on conflict (name) do nothing;
