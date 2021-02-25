FROM debian:10 AS dev-image

RUN apt-get update -yq \
&& apt-get install curl gnupg -yq \
&& curl -sL https://deb.nodesource.com/setup_10.x | bash \
&& apt-get install nodejs -yq \
&& apt-get clean -y


# set working directory
WORKDIR /app

# install and cache app dependencies
COPY package.json /app/package.json
RUN npm install
RUN npm install -g @angular/cli@11.0.5

EXPOSE 4200

#ENTRYPOINT ["/bin/sh", "-c" , "npm install && ng serve --host 0.0.0.0"]

CMD ["ng","serve","--host", "0.0.0.0"]
