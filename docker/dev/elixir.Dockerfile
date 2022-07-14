FROM elixir:latest

WORKDIR /usr/app

COPY . .

RUN apt-get update -y && apt-get install -y build-essential inotify-tools curl
RUN curl -sL https://deb.nodesource.com/setup_16.x | bash
RUN apt-get install nodejs

RUN mix local.hex --force
RUN mix local.rebar --force
