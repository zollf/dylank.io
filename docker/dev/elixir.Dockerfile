FROM elixir:latest

WORKDIR /usr/app

COPY . .

RUN apt-get update
RUN apt-get -y install inotify-tools

RUN mix local.hex --force
RUN mix local.rebar --force
