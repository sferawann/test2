CREATE TABLE requirements
(
    id bigserial NOT NULL,
	nik character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
	alamat character varying(255) NOT NULL,
	phone_number character varying(20) NOT NULL,
	created_at timestamptz NOT NULL DEFAULT 'NOW()',
    CONSTRAINT requirements_pkey PRIMARY KEY (id)
);

CREATE TABLE lenders
(
    id bigserial NOT NULL,
    name character varying(255) NOT NULL,
	limits decimal(10, 2) NOT NULL,
	bunga decimal(10, 2) NOT NULL,
	created_at timestamptz NOT NULL,
    CONSTRAINT lenders_pkey PRIMARY KEY (id)
);

CREATE TABLE users
(
    id bigserial NOT NULL,
    username character varying(255) NOT NULL,
	password character varying(255) NOT NULL,
	id_role int REFERENCES roles(id) NOT NULL,
	created_at timestamptz NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id)
);

CREATE TABLE payment_methods
(
    id bigserial NOT NULL,
	name character varying(255) NOT NULL, 
	created_at timestamptz NOT NULL,
    CONSTRAINT payment_methods_pkey PRIMARY KEY (id)
);

CREATE TABLE accept_statuss
(
    id bigserial NOT NULL,
	id_transaction int REFERENCES transactions(id) NOT NULL,
	status bool NOT NULL, 
	created_at timestamptz NOT NULL,
    CONSTRAINT accept_statuss_pkey PRIMARY KEY (id)
);

CREATE TABLE transactions 
(
  id bigserial NOT NULL,
  id_requirement int REFERENCES requirements(id) NOT NULL,
  id_lender int REFERENCES lenders(id) NOT NULL,
	id_user int REFERENCES users(id) NOT NULL,
  amount decimal(10, 2) NOT NULL,
  transaction_date timestamptz NOT NULL,
  due_date timestamptz NOT NULL,
	CONSTRAINT transactions_pkey PRIMARY KEY (id)
);

CREATE TABLE payments 
(
  id bigserial NOT NULL,
  id_transaction int REFERENCES transactions(id) NOT NULL,
  payment_amount decimal(10, 2) NOT NULL,
  payment_date timestamptz NOT NULL,
  id_payment_method int REFERENCES payment_methods(id) NOT NULL,
	CONSTRAINT payments_pkey PRIMARY KEY (id)
);

CREATE TABLE status_loans 
(
  id bigserial NOT NULL,
  id_transaction int REFERENCES transactions(id) NOT NULL,
  id_accept_status int REFERENCES accept_statuss(id) NOT NULL,
	created_at timestamptz NOT NULL,
	CONSTRAINT status_loans_pkey PRIMARY KEY (id)
);

CREATE TABLE roles
(
    id bigserial NOT NULL,
    name character varying(255) NOT NULL,
	created_at timestamptz NOT NULL,
    CONSTRAINT roles_pkey PRIMARY KEY (id)
);