PLAYS SOUNDS ON A DOOR


Mac:
`brew install portaudio`
`brew install ffmpeg`

Nerdmachine:
`apt-get install portaudio19-dev`


Then:
`go get code.google.com/p/portaudio-go/portaudio`

Usage:

	./sounds
	curl --form "soundFile=@SOMEFILE.mp3" localhost:3030/play/