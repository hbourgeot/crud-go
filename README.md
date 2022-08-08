# Golang CRUD

Crud maked from scratch, using **PostgreSQL** as Data Base, without libraries.

## Table of Content

1. [Data Base Content](#data-base-content).

    1. [Table Clients](#table-clients).
    2. [Table Products](#table-products).
    3. Table Orders.
2. Database Package.

    1. Structs to use.
    2. Connection to PostgreSQL
    3. Functions for querying the Clients Table.
    4. Functions for querying the Products Table.
    5. Functions for querying the Orders Table.
3. Utilities Package.

    1. Functions for Print Titles and Separators.
    2. Reader and Read Entries.
4. Main Package.

    1. Principal menus.
    2. Show menus.
    3. Options menu.

## Data Base Content

### Table Clients

This table has *three* columns, two **VARCHAR** type and one **INTEGER** type. Here you can see a preview of this table:

| Column | Type                   | Null | Key |
| ------ | ---------------------- | ---- | --- |
| dni    | integer                | NO   | PRI |
| name   | character varying(120) | NO   |     |
| phone  | character varying(22)  | NO   |     |

If you want create this table on your PostgreSQL installation, use this code:

```
CREATE TABLE IF NOT EXISTS public.clients
(
   dni integer NOT NULL,
   name character varying(120) COLLATE pg_catalog."default" NOT NULL,
   phone character varying(22) COLLATE pg_catalog."default" NOT NULL,
   CONSTRAINT clients_pkey PRIMARY KEY (dni)
);
```

### Table Products

This table has *six* columns, two **VARCHAR** type, two **INTEGER** type, one **TEXT** type and one **NUMERIC** type.
Here you can see a preview of this table:

| Column          | Type                   | Null | Key |
|-----------------|------------------------|------|-----|
| cod             | integer                | NO   | PRI |
| name            | character varying(170) | NO   |     |
| brand           | character varying(60)  | NO   |     |
| description     | text                   | NO   |     |
| price           | numeric                | NO   |     |
| inventory_count | integer                | NO   |     |

If you want create this table on your PostgreSQL installation, use this code:

```
CREATE TABLE IF NOT EXISTS public.products
(
    cod integer NOT NULL,
    name character varying(170) COLLATE pg_catalog."default" NOT NULL,
    brand character varying(60) COLLATE pg_catalog."default" NOT NULL,
    description text COLLATE pg_catalog."default" NOT NULL,
    price numeric(7,0) NOT NULL,
    inventory_count integer NOT NULL,
    CONSTRAINT products_pkey PRIMARY KEY (cod)
)
```

### Table Orders

This table has *three* columns, all are **INTEGER** type. Here you can see a preview of this table:

| Column | Type    | Null | Key |
|-------|---------| ---- |-----|
| id    | integer | NO   | PRI |
| dni   | integer | NO   | FOR |
| cod   | integer | NO   | FOR |

This table is for create the **Data Base Relation**, that is N-N, because one client can buy N products, and one product
can be purchased by N clients. Also, if the DNI or Code is updated in the Clients or Product Table, respectively, here
will updated too.

If you want create this table on your PostgreSQL installation, use this code:

```
CREATE TABLE IF NOT EXISTS public.orders
(
    id integer NOT NULL DEFAULT nextval('orders_id_seq'::regclass),
    dni integer NOT NULL,
    cod integer NOT NULL,
    CONSTRAINT "PK_Orders" PRIMARY KEY (id),
    CONSTRAINT "FK_Clients" FOREIGN KEY (dni)
        REFERENCES public.clients (dni) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT "FK_Products" FOREIGN KEY (cod)
        REFERENCES public.products (cod) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
)
```