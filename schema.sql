# CREATE TABLE AND FIELD
CREATE TABLE `students` (
    `id` Int( 8 ) AUTO_INCREMENT NOT NULL,
    `identifier` BigInt( 14 ) NOT NULL,
    `name` VarChar( 255 ) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
    `email` VarChar( 255 ) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL UNIQUE,
    `semester` VarChar( 255 ) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
    `created_at` DateTime NOT NULL,
    `updated_at` DateTime NOT NULL,
    CONSTRAINT `unique_id` UNIQUE( `id` )
)
    CHARACTER SET = utf8
    COLLATE = utf8_general_ci
    ENGINE = InnoDB
    AUTO_INCREMENT = 3;

# DESCRIBE TABLE
desc students

# DATA DUMMY
INSERT INTO students(identifier, name, email, semester, created_at, updated_at) VALUES
('2003113950', 'sammi1', 'sammi1@gmail.com', '1', '2020-01-06 07:01:05', '2020-01-06 07:01:05'),
('2003113951', 'sammi2', 'sammi2@gmail.com', '2', '2020-02-06 07:02:05', '2020-02-06 07:02:05'),
('2003113952', 'sammi3', 'sammi3@gmail.com', '3', '2020-03-06 07:03:05', '2020-03-06 07:03:05'),
('2003113953', 'sammi4', 'sammi4@gmail.com', '4', '2020-04-06 07:04:05', '2020-04-06 07:04:05');