USE CL;

CREATE TABLE IF NOT EXISTS swipes(
    swipes_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    usr_id INT,
    swipe_with INT
);

CREATE TABLE IF NOT EXISTS matches(
    match_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    usr_id_1 INT, 
    usr_id_2 INT
);