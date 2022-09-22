-- Useful commands
--SELECT * FROM pg_database;
--SELECT rolname FROM pg_roles;

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);

SET default_tablespace = '';
SET default_table_access_method = heap;


DO
$do$
BEGIN
   IF EXISTS (
      SELECT FROM pg_catalog.pg_roles
      WHERE  rolname = 'react') THEN

      RAISE NOTICE 'Role "react" already exists. Skipping.';
   ELSE
      CREATE ROLE react 
      WITH LOGIN 
      CREATEDB
      PASSWORD 'frontend';
   END IF;
END
$do$;

CREATE SCHEMA IF NOT EXISTS fe  AUTHORIZATION react;
CREATE TABLE IF NOT EXISTS fe.users (
    id              VARCHAR(255) PRIMARY KEY,
    username        VARCHAR(255) NOT NULL,
    email           VARCHAR(255) NOT NULL,
    password        VARCHAR(255) NOT NULL,
    created_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS fe.accounts (
    id              VARCHAR(255) PRIMARY KEY,
    username        VARCHAR(255) NOT NULL,
    email           VARCHAR(255) NOT NULL,
    password        VARCHAR(255) NOT NULL,
    created_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

DO
$do$
BEGIN
   IF EXISTS (
      SELECT FROM pg_catalog.pg_roles
      WHERE  rolname = 'go') THEN
      RAISE NOTICE 'Role "go" already exists. Skipping.';
   ELSE
      CREATE ROLE react 
      WITH LOGIN 
      CREATEDB
      PASSWORD 'backend';
   END IF;
END
$do$;

CREATE SCHEMA IF NOT EXISTS be AUTHORIZATION go;
CREATE TABLE IF NOT EXISTS be.subscription (
    id              VARCHAR(255) PRIMARY KEY,
    username        VARCHAR(255) NOT NULL,
    email           VARCHAR(255) NOT NULL,
    password        VARCHAR(255) NOT NULL,
    created_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS be.payments (
    id              VARCHAR(255) PRIMARY KEY,
    username        VARCHAR(255) NOT NULL,
    email           VARCHAR(255) NOT NULL,
    password        VARCHAR(255) NOT NULL,
    created_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

