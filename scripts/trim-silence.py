import IPython
import numpy as np
from scipy.io import wavfile
import plotly.graph_objs as go
import plotly

file = "09-16-2022_13-17-44.wav"
newfile = "processed-" + file
file = "data/barkley/" + file

# Fix-sized segmentation (breaks a signal into non-overlapping segments)
fs, signal = wavfile.read(file)
signal = signal / (2**15)
signal_len = len(signal)
segment_size_t = .1 # segment size in seconds
segment_size = int(segment_size_t * fs)  # segment size in samples
# Break signal into list of segments in a single-line Python code
segments = np.array([signal[x:x + segment_size] for x in
                     np.arange(0, signal_len, segment_size)])


# Remove pauses using an energy threshold
energies = [(s**2).sum() / len(s) for s in segments]
# (attention: integer overflow would occure without normalization here!)
thres = np.percentile(energies, 95)
index_of_segments_to_keep = (np.where(energies > thres)[0])
# get segments that have energies higher than a the threshold:
segments2 = segments[index_of_segments_to_keep]

print(energies)

# concatenate segments to signal:
new_signal = np.concatenate(segments2)
# and write to file:
wavfile.write(newfile, fs, new_signal)
plotly.offline.iplot({ "data": [go.Scatter(y=energies, name="energy"),
                                go.Scatter(y=np.ones(len(energies)) * thres,
                                           name="thres")]})
# play the initial and the generated files in notebook:
IPython.display.display(IPython.display.Audio(file))
IPython.display.display(IPython.display.Audio(newfile))