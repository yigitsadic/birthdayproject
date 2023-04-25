--
-- PostgreSQL database dump
--

-- Dumped from database version 15.2
-- Dumped by pg_dump version 15.2

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: companies; Type: TABLE; Schema: public; Owner: birthdayproject
--

CREATE TABLE public.companies (
    id integer NOT NULL,
    name character varying(50) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.companies OWNER TO birthdayproject;

--
-- Name: companies_id_seq; Type: SEQUENCE; Schema: public; Owner: birthdayproject
--

CREATE SEQUENCE public.companies_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.companies_id_seq OWNER TO birthdayproject;

--
-- Name: companies_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: birthdayproject
--

ALTER SEQUENCE public.companies_id_seq OWNED BY public.companies.id;


--
-- Name: employees; Type: TABLE; Schema: public; Owner: birthdayproject
--

CREATE TABLE public.employees (
    id integer NOT NULL,
    first_name character varying(50) NOT NULL,
    last_name character varying(50) NOT NULL,
    email character varying(70) NOT NULL,
    birth_day integer NOT NULL,
    birth_month integer NOT NULL,
    company_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.employees OWNER TO birthdayproject;

--
-- Name: employees_id_seq; Type: SEQUENCE; Schema: public; Owner: birthdayproject
--

CREATE SEQUENCE public.employees_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.employees_id_seq OWNER TO birthdayproject;

--
-- Name: employees_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: birthdayproject
--

ALTER SEQUENCE public.employees_id_seq OWNED BY public.employees.id;


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: birthdayproject
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO birthdayproject;

--
-- Name: users; Type: TABLE; Schema: public; Owner: birthdayproject
--

CREATE TABLE public.users (
    id integer NOT NULL,
    first_name character varying(50) NOT NULL,
    last_name character varying(50) NOT NULL,
    email character varying(70) NOT NULL,
    password_hash character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    company_id integer NOT NULL
);


ALTER TABLE public.users OWNER TO birthdayproject;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: birthdayproject
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO birthdayproject;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: birthdayproject
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: companies id; Type: DEFAULT; Schema: public; Owner: birthdayproject
--

ALTER TABLE ONLY public.companies ALTER COLUMN id SET DEFAULT nextval('public.companies_id_seq'::regclass);


--
-- Name: employees id; Type: DEFAULT; Schema: public; Owner: birthdayproject
--

ALTER TABLE ONLY public.employees ALTER COLUMN id SET DEFAULT nextval('public.employees_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: birthdayproject
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: companies companies_pk; Type: CONSTRAINT; Schema: public; Owner: birthdayproject
--

ALTER TABLE ONLY public.companies
    ADD CONSTRAINT companies_pk PRIMARY KEY (id);


--
-- Name: companies company_name_uniqueness; Type: CONSTRAINT; Schema: public; Owner: birthdayproject
--

ALTER TABLE ONLY public.companies
    ADD CONSTRAINT company_name_uniqueness UNIQUE (name);


--
-- Name: employees employees_email_uniqueness; Type: CONSTRAINT; Schema: public; Owner: birthdayproject
--

ALTER TABLE ONLY public.employees
    ADD CONSTRAINT employees_email_uniqueness UNIQUE (email);


--
-- Name: employees employees_pk; Type: CONSTRAINT; Schema: public; Owner: birthdayproject
--

ALTER TABLE ONLY public.employees
    ADD CONSTRAINT employees_pk PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: birthdayproject
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: users users_email_uniqueness; Type: CONSTRAINT; Schema: public; Owner: birthdayproject
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_uniqueness UNIQUE (email);


--
-- Name: users users_pk; Type: CONSTRAINT; Schema: public; Owner: birthdayproject
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pk PRIMARY KEY (id);


--
-- Name: employees employee_companies_fk; Type: FK CONSTRAINT; Schema: public; Owner: birthdayproject
--

ALTER TABLE ONLY public.employees
    ADD CONSTRAINT employee_companies_fk FOREIGN KEY (company_id) REFERENCES public.companies(id);


--
-- Name: users users_company_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: birthdayproject
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_company_id_fk FOREIGN KEY (company_id) REFERENCES public.companies(id);


--
-- PostgreSQL database dump complete
--

