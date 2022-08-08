# Golang CRUD

Crud maked from scratch, using **PostgreSQL** as Data Base, without libraries.

## Table of Content

1. [Data Base Content](#data-base-content).

   1. [Table Clients](#table-clients).
   2. [Table Products](#table-products).
   3. [Table Orders](#table-orders).
2. [Database Package](#database-package).

   1. [Structs to use](#structs-to-use).
   2. [Connection to PostgreSQL](#connection-to-postgresql).
   3. [Functions for querying the Clients and Products Table](#functions-for-querying-clients-and-products-tables).
   4. [Functions for querying the Orders Table](#functions-for-querying-the-orders-table).
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

```
| Column  | Type                   | Null  | Key  |
|---------|------------------------|-------|------|
| dni     | integer                | NO    | PRI  |
| name    | character varying(120) | NO    |      |
| phone   | character varying(22)  | NO    |      |
```

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

```
| Column          | Type                   | Null | Key |
|-----------------|------------------------|------|-----|
| cod             | integer                | NO   | PRI |
| name            | character varying(170) | NO   |     |
| brand           | character varying(60)  | NO   |     |
| description     | text                   | NO   |     |
| price           | numeric                | NO   |     |
| inventory_count | integer                | NO   |     |
```

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

```
| Column  | Type    | Null | Key |
|---------|---------|------|-----|
| id      | integer | NO   | PRI |
| dni     | integer | NO   | FOR |
| cod     | integer | NO   | FOR |
```

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

## Database Package

The CRUD has a package that I named ```database````, in this are sotarged all the functions for querying data from the
data base. Here is the half of the magic of this CRUD.

### Structs to use

In the path ```internal/database/structs.go``` is located all the structs for receive the data of the tables. In the
file are *three* structs, for each table, obviously.

### Connection to PostgreSQL

For make the connection to the database, the function ```makeCN``` returns a pointer object to the database and an
error, this function is located on the ```internal/database/database.go``` file.

### Functions for querying Clients and Products Tables

The functions for each table are located in separated files, ```internal/database/cliens.go```
and ```internal/database/products.go``` respecctively. In these files, are located the below functions:

```
| Function          | Parameters                                                                                | Return             |
|-------------------|-------------------------------------------------------------------------------------------|--------------------|
| CreateClients     | dni int, name string, phone string                                                        | error              |
| GetClientByDNI    | dni int                                                                                   | error              |
| GetAllClients     | none                                                                                      | []*Clients, error  |
| UpdateClient      | columnEdit string, newValue string, dni int                                               | error              |
| DeleteClient      | dni int                                                                                   | error              |
| InsertProducts    | cod int, name string, brand string, description string, price float64, inventoryCount int | error              |
| GetProductsByCode | cod int                                                                                   | error              |
| GetAllProducts    | none                                                                                      | []*Products, error |
| UpdateProducts    | columnEdit string, newValue any, cod int                                                  | error              |
| DeleteProducts    | cod int                                                                                   | error              |
```

The functions ```CreateClients``` and ```InsertProducts``` do the same, but for its respective tables, only changes the
parameters count. Also, ```GetClientByDNI```, ```GetAllClients```, ```GetProductsByCode``` and ```GetAllProducts```
functions are similar, only are different on the data returned, this is the same case for the
functions ```DeleteClient``` and ```DeleteProducts```.

The functions what are more different (relatively) are ```UpdateClient``` and ```UpdateProducts```, only because in the
first I use if statement, the second uses switch statement.

### Functions for querying the Orders Table

The functions used here are similar to the mentioned previously.

```
| Function        | Parameters | Return           |
|-----------------|------------|------------------|
| GetOrdersByDNI  | dni int    | error            |
| GetOrdersByCode | cod int    | error            |
| GetAllOrders    | none       | []*Orders, error |
```

In ```GetOrdersByDNI``` we print all the orders by this client, the case is similar for ```GetOrdersByCode```, only
changes the client by the product. Finishing, in ```GetAllOrders``` we pass the data of this table.