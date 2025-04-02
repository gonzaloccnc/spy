FROM ubuntu:latest

RUN apt-get update
RUN apt-get install -y curl
RUN curl -L https://yt-dl.org/downloads/latest/youtube-dl -o /usr/local/bin/youtube-dl
RUN chmod a+rx /usr/local/bin/youtube-dl
RUN curl -LO https://go.dev/dl/go1.24.1.linux-amd64.tar.gz
RUN rm -rf /usr/local/go
RUN tar -C /usr/local -xzf go1.24.1.linux-amd64.tar.gz
RUN rm go1.24.1.linux-amd64.tar.gz

ENV PATH="/usr/local/go/bin:${PATH}"

WORKDIR /app
COPY . .
RUN go build -o server ./main.go
RUN chmod +x ./server
CMD ["./server"]
