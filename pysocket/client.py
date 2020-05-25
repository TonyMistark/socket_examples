import socket

HOST = "127.0.0.1"
PORT = 9091

with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
    s.connect((HOST, PORT))
    s.sendall(bytes("Hello, World", encoding="utf-8"))
    data = s.recv(1024)
    print("Received", repr(data))
    s.close()
