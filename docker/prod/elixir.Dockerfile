FROM elixir:latest

ENV MIX_ENV prod
ENV ENVRIONMENT prod

WORKDIR /usr/app

COPY /backend .

RUN apt-get update

RUN mix local.hex --force
RUN mix local.rebar --force
RUN mix deps.get --only prod
RUN mix compile
CMD mix ecto.migrate && mix phx.server