--
-- Table structure for table `rooms`
--

DROP TABLE IF EXISTS room;

CREATE TABLE room (
    id SERIAL PRIMARY KEY,
    price INTEGER NOT NULL,
    description TEXT NOT NULL,
    date_added TIMESTAMP NOT NULL
);


--
-- Table structure for table `booking`
--

DROP TABLE IF EXISTS booking;

CREATE TABLE booking (
    id SERIAL PRIMARY KEY,
    room_id INTEGER NOT NULL REFERENCES room,
    start_date TIMESTAMP NOT NULL,
    expiration_date TIMESTAMP NOT NULL
);

-- TODO Create index on room_id in booking table