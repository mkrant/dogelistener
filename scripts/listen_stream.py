import json

import pyaudio
import signal
import wave
import time
import socket
from datetime import datetime

chunk = 1024  # Record in chunks of 1024 samples
sample_format = pyaudio.paInt16  # 16 bits per sample
channels = 1
fs = 8000  # Record at 8000 samples per second *8 KHz
maxDurationSeconds = 15800 # 1 hour

now = datetime.now()
dt_string = now.strftime("%m-%d-%Y_%H-%M-%S")

p = pyaudio.PyAudio()  # Create an interface to PortAudio

print(f'Recording for maximum {maxDurationSeconds / 60} minutes ({maxDurationSeconds} seconds). Ctrl + C to stop')

stream = p.open(format=sample_format,
                channels=channels,
                rate=fs,
                frames_per_buffer=chunk,
                input=True)

signal.signal(signal.SIGINT, signal.default_int_handler)
start = time.time()

num = 1

sock = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
# Connect the socket to the port where the server is listening
server_address = '/tmp/echo.sock'
sock.connect(server_address)

try:
    while 1:
        data = [stream.read(chunk)]

        filename = "{:08d}.wav".format(num)
        wf = wave.open(f'scripts/tmp/{filename}', 'wb')
        wf.setnchannels(channels)
        wf.setsampwidth(p.get_sample_size(sample_format))
        wf.setframerate(fs)
        wf.writeframes(b''.join(data))
        wf.close()

        payload = {
            'type': 'process_file',
            'payload': {
                'filename': filename
            }
        }
        data = json.dumps(payload)
        sock.sendall(bytes(data, 'utf-8'))

        num += 1

        exit(0)

        if time.time() - start > maxDurationSeconds:
            print("Max time hit, all done")
            break
except KeyboardInterrupt:
    print("Cancelled, all done")


# Stop and close the stream
stream.stop_stream()
stream.close()
# Terminate the PortAudio interface
p.terminate()

print('Finished recording')
