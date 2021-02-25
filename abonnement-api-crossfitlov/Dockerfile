FROM debian:9

RUN apt-get update -yq \
&& apt-get -y upgrade 

RUN apt-get install -y python3-pip python3-dev

COPY ./requirements.txt /app/requirements.txt

WORKDIR /app

RUN pip3 install -r requirements.txt

RUN apt-get clean -y

COPY . /app

EXPOSE 8080

ENTRYPOINT [ "python3" ]

CMD [ "abonnement_api.py" ]