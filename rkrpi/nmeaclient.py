import socket
from .database import Database


if __name__ == "__main__":
    db = Database()
    db.init_tables()

    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        s.connect(("192.168.76.1", 10110))
        s_file = s.makefile(newline="\r\n")
        while True:
            msgs = [(s_file.readline(),) for _ in range(100)]
            db.create_msgs(msgs)
