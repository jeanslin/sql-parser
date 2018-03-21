CREATE TYPE session_status AS ENUM ('new', 'finished', 'active', 'declined');


CREATE TABLE sessions
(
  id        VARCHAR(255)                                     NOT NULL
    CONSTRAINT sessions_pkey
    PRIMARY KEY,
  creatorid INTEGER                                          NOT NULL,
  abonentid INTEGER                                          NOT NULL,
  status    SESSION_STATUS DEFAULT 'new' :: SESSION_STATUS   NOT NULL,
  createdat TIMESTAMP DEFAULT timezone('utc' :: TEXT, now()) NOT NULL,
  updatedat TIMESTAMP DEFAULT timezone('UTC' :: TEXT, now()) NOT NULL
);

CREATE UNIQUE INDEX sessions_id_uindex
  ON sessions (id);


CREATE OR REPLACE FUNCTION trigger_upd_time()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updatedat = (NOW() AT TIME ZONE 'UTC');
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER set_upd_time
BEFORE UPDATE ON sessions
FOR EACH ROW
EXECUTE PROCEDURE trigger_upd_time();