FROM elixir:latest

WORKDIR /usr/app

COPY . .

RUN apt-get update
RUN apt-get -y install inotify-tools
RUN apt-get install curl
RUN curl -sL https://deb.nodesource.com/setup_16.x | bash
RUN apt-get install nodejs

RUN mix local.hex --force
RUN mix local.rebar --force
