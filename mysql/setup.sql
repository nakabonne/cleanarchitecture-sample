CREATE TABLE users (
    id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    name VARCHAR(20) NOT NULL,
    email VARCHAR(100) NOT NULL,
    age smallint,
    PRIMARY KEY (id)
);

insert into users(name, email) values("joe", "joe.com");
insert into users(name, email) values("john", "john.com");
