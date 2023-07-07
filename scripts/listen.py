import pyaudio
import signal
import wave
import time
from datetime import datetime

chunk = 8000  # Record in chunks of 1024 samples
sample_format = pyaudio.paInt16  # 16 bits per sample
channels = 1

fs = 8000  # Record at 8000 samples per second *8 KHz

now = datetime.now()
dt_string = now.strftime("%m-%d-%Y_%H-%M-%S")

person = "barkley"
filename = f'data/{person}/{dt_string}.wav'

p = pyaudio.PyAudio()  # Create an interface to PortAudio

maxDurationSeconds = 7200
print(f'Recording for maximum {maxDurationSeconds / 60} minutes ({maxDurationSeconds} seconds). Ctrl + C to stop')

stream = p.open(format=sample_format,
                channels=channels,
                rate=fs,
                frames_per_buffer=chunk,
                input=True)

frames = []  # Initialize array to store frames

signal.signal(signal.SIGINT, signal.default_int_handler)
start = time.time()

try:
    while 1:
        data = stream.read(chunk)
        frames.append(data)
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

# Save the recorded data as a WAV file
wf = wave.open(filename, 'wb')
wf.setnchannels(channels)
wf.setsampwidth(p.get_sample_size(sample_format))
wf.setframerate(fs)
wf.writeframes(b''.join(frames))
wf.close()