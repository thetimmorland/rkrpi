import sqlite3


SCHEMA = """
CREATE TABLE IF NOT EXISTS msgs (
  id INTEGER PRIMARY KEY NOT NULL,
  msg TEXT NOT NULL
);
"""

TRIGGER = """
CREATE TRIGGER IF NOT EXISTS deleteold
  AFTER INSERT ON msgs
  WHEN (SELECT count(*) FROM msgs) > 50000000
BEGIN
  DELETE FROM msgs
  WHERE id < (SELECT MIN(id) from msgs) + 10000;
END;
"""

CREATE = """
INSERT INTO msgs(msg) VALUES (?);
"""

READ = """
WITH _msgs AS (
  SELECT * from msgs
  WHERE id >= ?
  ORDER BY id
  LIMIT ?
) SELECT
  MIN(id) as start,
  MAX(id) as stop,
  GROUP_CONCAT(msg) as msgs
FROM _msgs;
"""


class Database:
    def __init__(self):
        self._conn = sqlite3.connect("rkrpi.db", check_same_thread=False)
        self._conn.row_factory = sqlite3.Row

    def init_tables(self):
        self._conn.execute(SCHEMA)
        self._conn.execute(TRIGGER)

    def create_msgs(self, msgs):
        with self._conn as cur:
            cur.executemany(CREATE, msgs)

    def read_msgs(self, offset, limit):
        return self._conn.execute(READ, (offset, limit)).fetchone()
