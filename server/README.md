## To launch server
```
make build-docker-image run-docker-image
```
Then: [http://localhost:9080](http://localhost:9080)

To encode : ffmpeg -i  in.mp4 -y -vf scale=640:380  -vcodec h264 -b:v 100000 -acodec aac -b:a 640000 -ss 00:10:00.0 -t 00:00:30.0 out.mp4