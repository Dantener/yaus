CREATE TABLE urls
(
    id        varchar(10)   not null unique,
    shortURL  varchar(2100) not null unique,
    originURL varchar(2100) not null unique

);

CREATE INDEX IF NOT EXISTS origin_url_idx ON urls (originURL);
