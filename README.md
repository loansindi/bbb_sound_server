# PLAYS SOUNDS ON A DOOR

## Description

A Golang-based REST server. Starts on port 3030. Serves up a little form where you can upload an audio file on `/`.

`/play/` accepts `POST` requests with content-type `multipart/form-data`. Must have a form field `soundFile`, which is the audio file to be played. Audio file must be <10 MB. 

Multiple audio files stack up one on top of the other.


## Requirements

Mac:

`brew install mplayer`

Beaglebone Black:

(probably shouldn't have to do anything special)





## Usage

	./sounds
	curl --form "soundFile=@SOMEFILE.mp3" localhost:3030/play/

Alternatively, start the server and browse to localhost:3030

## Probably don't

- Put this on the open internet.
- Whatever, I'm not the boss of you
- Go for it
- Shine on you crazy diamond