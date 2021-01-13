--
-- Table structure for table `rooms`
--

DROP TABLE IF EXISTS room;

CREATE TABLE room (
    id SERIAL PRIMARY KEY,
    price NUMERIC NOT NULL,
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
    start_date VARCHAR(13) NOT NULL,
    expiration_date VARCHAR(13) NOT NULL
);

CREATE INDEX booking_room_id_fkey ON booking (room_id);
-- TODO delete cascade room and booking