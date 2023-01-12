CREATE TABLE IF NOT EXISTS customers(
    id varchar(50) PRIMARY KEY,
    surname varchar(90) NOT NULL,
    firstname varchar(90) NOT NULL,
    patronym varchar(90) NOT NULL,
    age varchar(4),
    date_created timestamp NOT NULL
    );
CREATE TABLE IF NOT EXISTS shops(
    id varchar(50) PRIMARY KEY,
    shopname varchar(255) NOT NULL,
    address text NOT NULL,
    work_status boolean,
    owner text
    );