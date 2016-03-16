SET @@global.show_compatibility_56=ON;

DROP DATABASE IF EXISTS test_db;
CREATE DATABASE test_db;

USE test_db;

CREATE TABLE creditcards (
  id INT AUTO_INCREMENT PRIMARY KEY,
  ccnumber VARCHAR(20),
  cctype TEXT
);

INSERT INTO creditcards (ccnumber, cctype) VALUES ("372079560813168", "amex");
INSERT INTO creditcards (ccnumber, cctype) VALUES ("3760-3161-5608-974", "amex");
INSERT INTO creditcards (ccnumber, cctype) VALUES ("36157344819566", "diners");
INSERT INTO creditcards (ccnumber, cctype) VALUES ("3007-094-9008-897", "diners");
INSERT INTO creditcards (ccnumber, cctype) VALUES ("6011947984608576", "discover");
INSERT INTO creditcards (ccnumber, cctype) VALUES ("6011-3430-0923-2891", "discover");
INSERT INTO creditcards (ccnumber, cctype) VALUES ("214918445574804", "enroute");
INSERT INTO creditcards (ccnumber, cctype) VALUES ("2014-0582-1051-586", "enroute");
INSERT INTO creditcards (ccnumber, cctype) VALUES ("3088412336401191", "jcb");
INSERT INTO creditcards (ccnumber, cctype) VALUES ("3158-0990-2497-7528", "jcb");
INSERT INTO creditcards (ccnumber, cctype) VALUES ("180026367124713", "jcb");
INSERT INTO creditcards (ccnumber, cctype) VALUES ("2100-1773-6605-916", "jcb");
INSERT INTO creditcards (ccnumber, cctype) VALUES ("5481223188541697", "mastercard");
INSERT INTO creditcards (ccnumber, cctype) VALUES ("5128-7128-4843-4755", "mastercard");
INSERT INTO creditcards (ccnumber, cctype) VALUES ("4539524757964623", "visa");
INSERT INTO creditcards (ccnumber, cctype) VALUES ("4916-7273-3707-8631", "visa");
INSERT INTO creditcards (ccnumber, cctype) VALUES ("4539772651264", "visa");
INSERT INTO creditcards (ccnumber, cctype) VALUES ("4024-0071-2290-2", "visa");
INSERT INTO creditcards (ccnumber, cctype) VALUES ("869934845122647", "voyager");
INSERT INTO creditcards (ccnumber, cctype) VALUES ("8699-3169-9537-961", "voyager");
INSERT INTO creditcards (ccnumber, cctype) VALUES ("1234123412341234", "invalid");
INSERT INTO creditcards (ccnumber, cctype) VALUES ("1234-1234-1234-1234", "invalid");
INSERT INTO creditcards (ccnumber, cctype) VALUES ("000000000000", "invalid");
INSERT INTO creditcards (ccnumber, cctype) VALUES ("0000000000000000", "invalid");
INSERT INTO creditcards (ccnumber, cctype) VALUES ("0000-0000-0000-0000", "invalid");
INSERT INTO creditcards (ccnumber, cctype) VALUES ("111111111111", "invalid");
INSERT INTO creditcards (ccnumber, cctype) VALUES ("1111111111111111", "invalid");
INSERT INTO creditcards (ccnumber, cctype) VALUES ("1111-1111-1111-1111", "invalid");
INSERT INTO creditcards (ccnumber, cctype) VALUES ("888888888888", "invalid");
INSERT INTO creditcards (ccnumber, cctype) VALUES ("8888888888888888", "invalid");
INSERT INTO creditcards (ccnumber, cctype) VALUES ("8888-8888-8888-8888", "invalid");

CREATE TABLE posts (
  id INT AUTO_INCREMENT PRIMARY KEY,
  published_date DATE,
  published_time TIME,
  title CHAR(100),
  body TEXT
);

INSERT INTO posts (published_date, published_time, title, body) VALUES (NOW(), NOW(), "Stuff", "It's amazing what stuff can do!");
INSERT INTO posts (published_date, published_time, title, body) VALUES (NOW(), NOW(), "I want stuff!", "I want the ultrastuff 3000. My credit card number is 5520557624359492");
INSERT INTO posts (published_date, published_time, title, body) VALUES (NOW(), NOW(), "Hey wiz1234567890123", "Want to increase the size of your cardboard?!? Ping me!");
