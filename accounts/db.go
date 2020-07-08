package accounts

import "database/sql"

const (
	// Drop all tables, keys and sequences, then create fresh ones.
	tablesql = `DROP TABLE IF EXISTS public.users CASCADE;
	DROP TABLE IF EXISTS public.groups CASCADE;
	DROP TABLE IF EXISTS public.sites CASCADE;
	DROP TABLE IF EXISTS public.memberships CASCADE;

	CREATE SEQUENCE public.groups_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

	CREATE SEQUENCE public.sites_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
	CACHE 1;

	CREATE SEQUENCE public.users_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

	CREATE TABLE public.users
(
    id bigint NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    name character varying(100) COLLATE pg_catalog."default" NOT NULL,
    email character varying(200) COLLATE pg_catalog."default" NOT NULL,
    first character varying(200) COLLATE pg_catalog."default",
    last character varying(200) COLLATE pg_catalog."default",
    admin boolean NOT NULL DEFAULT false,
    CONSTRAINT users_pkey PRIMARY KEY (id)
)

	CREATE TABLE public.groups
(
    id bigint NOT NULL DEFAULT nextval('groups_id_seq'::regclass),
    name character varying(100) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT groups_pkey PRIMARY KEY (id)
)

	CREATE TABLE public.sites
(
    id bigint NOT NULL DEFAULT nextval('sites_id_seq'::regclass),
    owner bigint NOT NULL,
    domain character varying(200) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT sites_pkey PRIMARY KEY (id),
    CONSTRAINT fkey_sites_owner FOREIGN KEY (owner)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)

	CREATE TABLE public.memberships
(
    userid bigint NOT NULL,
    groupid bigint NOT NULL,
    CONSTRAINT memberships_pkey PRIMARY KEY (userid)
        INCLUDE(groupid),
    CONSTRAINT fkey_member_groupid FOREIGN KEY (groupid)
        REFERENCES public.groups (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT fkey_member_userid FOREIGN KEY (userid)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)
`

	// Get all tables if the user has access to them.
	checktablesql = `SELECT to_regclass('public.users'),to_regclass('public.groups'),to_regclass('public.sites'),to_regclass('public.memberships');`
)

// CheckTables returns true if the required tables exist.
func (m *Manager) CheckTables() bool {
	var users, groups, sites, memberships sql.NullBool
	err := m.QueryRow(checktablesql).Scan(&users, &groups, &sites, &memberships)
	if err != nil {
		return false
	}

	return users.Bool && groups.Bool && sites.Bool && memberships.Bool
}

// CreateTables creates the tables, keys and sequences.
// The database must already exist, and the connecting user must have full read/write access to it.
// This allows the account system to share a database with other subsytems.
func (m *Manager) CreateTables() error {
	return nil
}
