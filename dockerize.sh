docker build -t forum .
docker run -d -p 8000:8000 --rm --name c_forum  forum
