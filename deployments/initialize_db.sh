DROP DATABASE practice;

CREATE DATABASE practice;
USE practice;

CREATE TABLE instances (
  uuid VARCHAR(255),
  name VARCHAR(255),
  owner VARCHAR(255),
PRIMARY KEY (uuid) );

CREATE TABLE images (
  uuid VARCHAR(255),
  name VARCHAR(255),
  owner VARCHAR(255),
PRIMARY KEY (uuid) );

INSERT INTO practice.instances (uuid, name, owner)
  VALUES ("cde11bfa-e973-11ea-9c6e-0a15bbaed69b", "web01-2020-08-29", "bd64e65e-e972-11ea-842a-0a15bbaed69b");
INSERT INTO practice.instances (uuid, name, owner)
  VALUES ("b38f005a-e973-11ea-b90d-0a15bbaed69b", "mysql01-2020-08-29", "fba89220-4eb3-4c6e-9990-f2dea16fb928");

INSERT INTO practice.images (uuid, name, owner)
  VALUES ("cde11bfa-e973-11ea-9c6e-0a15bbaed69b", "web01-2020-08-29", "bd64e65e-e972-11ea-842a-0a15bbaed69b");
INSERT INTO practice.images (uuid, name, owner)
  VALUES ("b38f005a-e973-11ea-b90d-0a15bbaed69b", "mysql01-2020-08-29", "bd64e65e-e972-11ea-842a-0a15bbaed69b");
