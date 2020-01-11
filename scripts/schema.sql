CREATE TABLE hospital
(
    id                 SERIAL PRIMARY KEY,
    name               VARCHAR(100),
    max_patient_amount INT
);

CREATE TABLE location
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(100),
    hospital_id INT,
    FOREIGN KEY (hospital_id) REFERENCES hospital (id)
);

CREATE TABLE patient
(
    id             SERIAL PRIMARY KEY,
    name           VARCHAR(100),
    illness        VARCHAR(200),
    birth_date     DATE,
    location_id    INT NOT NULL,
    FOREIGN KEY (location_id) REFERENCES location (id)
);