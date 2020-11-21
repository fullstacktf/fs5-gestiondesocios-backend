USE fullstackAsociacion;

CREATE TABLE assoc_partners (
  id INT PRIMARY KEY,
  partnerName VARCHAR(30) NOT NULL
);

CREATE TABLE games (
  id INT PRIMARY KEY,
  id_owner INT,
  entry_date VARCHAR(200),
  disponibility BOOL NOT NULL,
  comments VARCHAR(200),
  CONSTRAINT fk_idOwner FOREIGN KEY (id_owner) REFERENCES assoc_partners (id)
);

CREATE TABLE borrowedGames (
  idGame INT PRIMARY KEY,
  idBorrower INT NOT NULL,
  borrowDate DATE,
  FOREIGN KEY (idBorrower) REFERENCES assoc_partners (id),
  CONSTRAINT fk_idGame FOREIGN KEY (idGame) REFERENCES games (id)

);

CREATE TABLE assoc_users (
  id INT PRIMARY KEY,
  username VARCHAR(30) NOT NULL,
  user_password VARCHAR(30) NOT NULL
);

INSERT INTO assoc_users VALUES (1, "David", "password");
INSERT INTO assoc_partners VALUES (1, "Pepe");


