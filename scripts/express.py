import pandas as pd
from plotly.subplots import make_subplots
import plotly.io as pio
import plotly.express as px
from pyAudioAnalysis import ShortTermFeatures as aF
from pyAudioAnalysis import audioBasicIO as aIO
import plotly.graph_objects as go
import numpy as np
import json
import os
from typing import List
from datetime import datetime
import math
from bs4 import BeautifulSoup


def create_graph(path: str):
    raw_runs = parse_data(path)
    subplots_all = create_subplots(raw_runs)

    max_samples_per_page = 5
    subplot_2d = chunks(subplots_all, max_samples_per_page)
    page = 1

    for subplots in subplot_2d:
        titles = []
        for subplot in subplots:
            time_nums = subplot.x.tolist()
            titles.append(f'{math.ceil(time_nums[len(time_nums) - 1])} minute sample')

        fig = make_subplots(rows=len(subplots), cols=1, x_title='Time (minutes)', y_title='Noise Level',
                            subplot_titles=titles, shared_xaxes=True, shared_yaxes='all')
        row = 1
        for subplot in subplots:
            fig.add_trace(subplot, row=row, col=1)
            row += 1

        fig.update_layout(title_text=f'Barkley Barking {subplots[0].name} to {subplots[len(subplots) - 1].name}')

        if page == 1:
            fig.show()

        pio.write_html(fig, f'cmd/staticserver/static/data/{page}/index.html')
        modify_html(page)
        page += 1


def create_subplots(raw_runs: []) -> List[go.Scatter]:
    raw_runs.sort(key=lambda x: x['date'], reverse=True)
    subplots = []

    for run in raw_runs:
        subplots.append(go.Scatter(
            name=run['date'].strftime("%b %d %Y"),
            x=run["Time (minutes)"],
            y=run["Energy"],
        ))

    return subplots


def parse_data(path: str) -> List[object]:
    dir_list = os.listdir(path)

    wav_files = []
    json_files = []
    data_runs = []

    date_fmt = "%m-%d-%Y_%H-%M-%S"

    for file in dir_list:
        if file.endswith('.wav'):
            wav_files.append(file)
        if file.endswith('.json'):
            name = os.path.splitext(file)[0]
            json_files.append(file)
            time, energy = read_json_file(os.path.join(path, file))

            obj = {
                "Time (minutes)": time,
                "Energy": energy,
                "date": datetime.strptime(name, date_fmt),
            }
            data_runs.append(obj)

    for wav_file in wav_files:
        name = os.path.splitext(wav_file)[0]
        if name + '.json' not in json_files:
            # we need to load audio file and crunch data
            wav_time, wave_energy = read_wav_file(os.path.join(path, wav_file))
            write_json_file(os.path.join(path, name + '.json'), wav_time, wave_energy)

            obj = {
                "Time (minutes)": wav_time,
                "Energy": wave_energy,
                "date": datetime.strptime(name, date_fmt),
            }

            data_runs.append(obj)

    return data_runs


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


def read_json_file(file: str) -> (np.ndarray, np.ndarray):
    with open(file, 'r') as openfile:
        dictionary = json.load(openfile)

    return np.asarray(dictionary["time"]), np.asarray(dictionary["energy"])


def write_json_file(file: str, time: np.ndarray, energy: np.ndarray):
    dictionary = {
        "time": time.tolist(),
        "energy": energy.tolist(),
    }
    json_object = json.dumps(dictionary)

    with open(file, "w") as outfile:
        outfile.write(json_object)


def chunks(lst, n) -> []:
    """Yield successive n-sized chunks from lst."""
    for i in range(0, len(lst), n):
        yield lst[i:i + n]


def modify_html(page: int):
    with open(f'cmd/staticserver/static/data/{page}/index.html', 'r') as file:
        html_doc = file.read()

        soup = BeautifulSoup(html_doc, 'html.parser')
        bodyTag = soup.find('body')
        divTag = soup.new_tag('div')

        prevTag = soup.new_tag('a')
        prevTag['style'] = 'font-size: 30px'
        prevTag['id'] = 'previous'
        prevTag['href'] = 'changeme'
        prevTag.insert(0, 'Previous')

        nextTag = soup.new_tag('a')
        nextTag['style'] = 'font-size: 30px'
        nextTag['id'] = 'next'
        nextTag['href'] = 'changeme'
        nextTag.insert(0, 'Next')

        scriptTag = soup.new_tag('script')
        scriptTag['type'] = "text/javascript"

        js = """
    let test = window.location.pathname.match("\\\\d+") ?? 1;
    console.log(test);
    let num = parseInt(test);
    console.log(num)
    
    if (num-1 > 0) {
        document.getElementById("previous").href = window.location.origin + "/data/" + (num-1);
    } else {
        console.log("lower");
        document.getElementById("previous").hidden = true;
    }
    
    document.getElementById("next").href = window.location.origin + "/data/" + (num+1);
    """
        scriptTag.insert(0, js)

        divTag.append(prevTag)
        divTag.append(nextTag)
        divTag.append(scriptTag)

        bodyTag.insert(0, divTag)

    with open(f'cmd/staticserver/static/data/{page}/index.html', "w") as wfile:
        wfile.write(str(soup))


create_graph('data/barkley')
