USE CL;

CREATE TABLE IF NOT EXISTS users(
    id_usr INT PRIMARY KEY 
);

CREATE TABLE IF NOT EXISTS rooms(
    room_id INT PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS messages(
    msg_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    msg_text VARCHAR(500),
    date_time DATETIME, 
    room_id INT,
    FOREIGN KEY(room_id) REFERENCES rooms(room_id) ON DELETE SET NULL,
    id_usr INT,
    FOREIGN KEY(id_usr) REFERENCES users(id_usr) ON DELETE SET NULL
);