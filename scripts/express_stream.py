import sys
from pyAudioAnalysis import ShortTermFeatures as aF
from pyAudioAnalysis import audioBasicIO as aIO
import numpy as np
import json
import os
from typing import List
from datetime import datetime
import socket

def crunch_file(file: str):
    wav_time, wave_energy = read_wav_file(file)
    # send_socket_msg(file, wav_time, wave_energy)
    write_json_file('out.json', wav_time, wave_energy)


def read_wav_file(file: str) -> (np.ndarray, np.ndarray):
    # read audio data from file
    # (returns sampling freq and signal as a numpy array)
    fs, s = aIO.read_audio_file(file)

    # print duration in seconds:
    duration = len(s) / float(fs)
    print(f'duration = {duration} seconds')

    win, step = 1, 1
    [f, fn] = aF.feature_extraction(s, fs, int(fs * win),
                                    int(fs * step))
    time = np.arange(0, duration, win)
    energy = f[fn.index('energy'), :]
    for i in range(len(time)):
        time[i] = time[i] / 60

    return time, energy


def write_json_file(file: str, time: np.ndarray, energy: np.ndarray):
    dictionary = {
        "time": time.tolist(),
        "energy": energy.tolist(),
    }
    json_object = json.dumps(dictionary)

    with open(file, "w") as outfile:
        outfile.write(json_object)


def send_socket_msg(file: str, time: np.ndarray, energy: np.ndarray):
    sock = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
    # Connect the socket to the port where the server is listening
    server_address = '/tmp/echo.sock'
    sock.connect(server_address)

    data = json.dumps({
        'type': 'upload_data',
        'payload': {
            'time': time.tolist(),
            'energy': energy.tolist(),
            'index': file,
        }
    }) + "\n"

    sock.sendall(bytes(data, 'utf-8'))
    sock.close()


args = sys.argv[1:]
if len(args) != 1:
    print("File name as cli arg required")
    exit(1)

crunch_file(args[0])
