git pull && docker build -f ./build/Dockerfile -t my-blog-docker:latest . && docker stop blog-server && docker rm blog-server && docker run -d -p 9090:9090 --name blog-server my-blog-docker