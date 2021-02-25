USE CL;

CREATE TABLE IF NOT EXISTS dataCL(
    crossfitlovID INT(10) UNSIGNED NOT NULL,
    firstname varchar(255) NOT NULL,
    gender ENUM ('boy', 'girl') NOT NULL,
    age INT(7) UNSIGNED NOT NULL,
    boxCity varchar(50) DEFAULT NULL, /*ville où se situe la box de crossfit*/
    email varchar(50) DEFAULT NULL, 
    biography varchar(255) DEFAULT NULL,
    job varchar(70) DEFAULT NULL,
	created_at DATETIME NOT NULL, /*date dinscription*/
    PRIMARY KEY(crossfitlovID)
);

ALTER TABLE `CL`.`dataCL` ADD INDEX `created_at` (`created_at`);


/*INSERT INTO `dataCL` (`crossfitlovID`, `firstname`, `gender`, `age`, `boxCity`, `email`, `biography`, `job`, `created_at`) VALUES ('123', 'Pierre', 'boy', '22', 'Montévrain', 'beau, jeune et riche', 'étudiant', '2021-02-02 09:03:02');
*/
/*prénom age genre box de crossfit bio job date d'inscription*/
