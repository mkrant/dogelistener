# dogelistener

Listen and track your doge.

## Dependencies

* Go
* Python 3.8.16

## Run

```bash
pip install -r scripts/requirements.txt

source scripts/venv/bin/activate

# Capture audio until max time reached or cancelled (Ctrl + C)
python scripts/listen.py

# Process audio, create graphs and html files
python scripts/express.py
```

Resolve missing pip packages manually, requirements.txt is not up to date.

Some manual things I had to do:

* Clone PortAudio and build from source since it was failing to install with brew
* Clone pyAudioAnalysis locally and install those pip requirements

## Go server

```bash
go build -o server backend/main.go
```

## AWS Instance

[AWS instance file server](http://ec2-35-91-124-34.us-west-2.compute.amazonaws.com)

```bash
# scp the whole backend folder originally
scp -i ~/.ssh/aws/key -r backend ec2-user@35.91.124.34:/home/ec2-user

# after each run, scp the static data again
scp -i ~/.ssh/aws/key -r backend/static ec2-user@35.91.124.34:/home/ec2-user/backend
```

http://localhost:3000
