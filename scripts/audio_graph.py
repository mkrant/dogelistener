from pydub import AudioSegment
import numpy as np
from scipy.io import wavfile
from plotly.offline import init_notebook_mode
import plotly.graph_objs as go
import plotly

fs_wav, data_wav = wavfile.read("../data/barkley/09-16-2022_13-17-44.wav")
data_wav_norm = data_wav / (2**15)
time_wav = np.arange(0, len(data_wav)) / fs_wav
plotly.offline.iplot({ "data": [go.Scatter(x=time_wav,
                                           y=data_wav_norm,
                                           name='normalized audio signal')]})

