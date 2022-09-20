# Example 1: short-term feature extraction
from pyAudioAnalysis import ShortTermFeatures as aF
from pyAudioAnalysis import audioBasicIO as aIO
import numpy as np
import plotly.graph_objs as go
import plotly
import IPython

file = "../data/barkley/09-15-2022_12-00-00.wav"

# read audio data from file
# (returns sampling freq and signal as a numpy array)
fs, s = aIO.read_audio_file(file)

# play the initial and the generated files in notebook:
IPython.display.display(IPython.display.Audio(file))

# print duration in seconds:
duration = len(s) / float(fs)
print(f'duration = {duration} seconds')

win, step = 1, 1
[f, fn] = aF.feature_extraction(s, fs, int(fs * win),
                                int(fs * step))
print(f'{f.shape[1]} frames, {f.shape[0]} short-term features')
print('Feature names:')
for i, nam in enumerate(fn):
    print(f'{i}:{nam}')
time = np.arange(0, duration - step, win)
energy = f[fn.index('energy'), :]

for idx in range(0, len(time)):
    time[idx] /= 60

mylayout = go.Layout(yaxis=dict(title="frame energy value"),
                     xaxis=dict(title="time (minutes)"))
plotly.offline.iplot(go.Figure(data=[go.Scatter(x=time,
                                                y=energy)],
                               layout=mylayout))
