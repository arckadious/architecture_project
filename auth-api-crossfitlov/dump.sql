USE CL;

CREATE TABLE IF NOT EXISTS dataCL
(
    crossfitlovID int(10) unsigned NOT NULL AUTO_INCREMENT,
    user_login varchar(255) DEFAULT NULL,
    encrypted_passwd varchar(255) DEFAULT NULL,
    primary key(crossfitlovID),
    UNIQUE KEY `user_login` (`user_login`)
    
);