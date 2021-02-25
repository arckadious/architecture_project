FROM debian:9

ENV GOROOT=/usr/local/go
ENV PATH=$GOROOT/bin:$PATH 

EXPOSE 8000

RUN apt-get update -yq \
&& apt-get -y upgrade \
&& apt-get install curl gnupg -yq 

RUN curl -o go1.15.3.tar.gz https://dl.google.com/go/go1.15.3.linux-amd64.tar.gz \
&& tar -xf go1.15.3.tar.gz

RUN mv go /usr/local

RUN apt-get clean -y

WORKDIR /users-api-crossfitlov 

ENTRYPOINT ["go", "run", "main.go", "-debug=INFO"]
