import json
import socket
import sys

# Create a UDS socket
sock = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)

# Connect the socket to the port where the server is listening
server_address = '/tmp/echo.sock'
print(sys.stderr, f'connecting to {server_address}')
try:
    sock.connect(server_address)
    payload = {
        'type': 'process_file',
        'payload': {
            'filename': 'yoo'
        }
    }
    sock.sendall(bytes(json.dumps(payload), 'utf-8'))

except socket.error as msg:
    print(sys.stderr, msg)
    sys.exit(1)