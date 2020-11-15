USE fullstackAsociacion;

CREATE TABLE assocPartner (
  id INT PRIMARY KEY,
  partnerName VARCHAR(30) NOT NULL
);

CREATE TABLE borrowedGame (
  idGame INT PRIMARY KEY,
  idBorrower INT NOT NULL,
  borrowDate DATE,
  FOREIGN KEY (idBorrower) REFERENCES assocPartner (id)
);

CREATE TABLE game (
  id INT PRIMARY KEY,
  idOwner INT,
  entryDate DATE,
  disponibility BIT(1) NOT NULL,
  comments VARCHAR(200),
  CONSTRAINT fk_idOwner FOREIGN KEY (idOwner) REFERENCES assocPartner (id),
  CONSTRAINT fk_idGame FOREIGN KEY (id) REFERENCES borrowedGame (idGame)
);

CREATE TABLE user (
  id INT PRIMARY KEY,
  userName VARCHAR(30) NOT NULL,
  userPassword VARCHAR(30) NOT NULL
);




