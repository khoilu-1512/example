CREATE TABLE IF NOT EXISTS students (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    age INT NOT NULL
);

INSERT INTO students (name, age) VALUES ('John Doe', 20);
INSERT INTO students (name, age) VALUES ('Jane Doe', 21);
INSERT INTO students (name, age) VALUES ('John Smith', 22);
INSERT INTO students (name, age) VALUES ('Jane Smith', 23);
INSERT INTO students (name, age) VALUES ('John Johnson', 24);
INSERT INTO students (name, age) VALUES ('Jane Johnson', 25);
INSERT INTO students (name, age) VALUES ('John Jones', 26);
