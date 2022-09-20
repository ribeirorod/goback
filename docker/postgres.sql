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

CREATE ROLE IF NOT EXISTS "react"
WITH LOGIN
CREATEDB
NOSUPERUSER
PASSWORD 'frontend';

CREATE SCHEMA IF NOT EXISTS frontend.fe  
    CREATE TABLE users (
        id              VARCHAR(255) PRIMARY KEY,
        username        VARCHAR(255) NOT NULL,
        email           VARCHAR(255) NOT NULL,
        password        VARCHAR(255) NOT NULL,
        created_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    )
AUTHORIZATION react;
