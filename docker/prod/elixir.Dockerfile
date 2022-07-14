ARG DEBIAN_VERSION=bullseye-20210902-slim

ARG BUILDER_IMAGE="elixir:latest"
ARG RUNNER_IMAGE="debian:${DEBIAN_VERSION}"

FROM ${BUILDER_IMAGE} as builder

RUN apt-get update -y && apt-get install -y build-essential git curl \
    && apt-get clean && rm -f /var/lib/apt/lists/*_*

RUN curl -sL https://deb.nodesource.com/setup_16.x | bash
RUN apt-get install nodejs
RUN npm install -g npm

WORKDIR /app

RUN mix local.hex --force && \
    mix local.rebar --force

ENV MIX_ENV="prod"

# COMPILE DEPS
COPY mix.exs mix.lock ./
RUN mix deps.get --only $MIX_ENV
RUN mkdir config

COPY config/config.exs config/${MIX_ENV}.exs config/
RUN mix deps.compile

COPY priv priv

# COMPILE ASSETS
COPY assets assets
RUN mix assets.install
RUN mix assets.deploy

# COMPILE RELEASE
COPY lib lib
RUN mix compile

COPY config/runtime.exs config/
RUN mix release

# start a new build stage so that the final image will only contain
# the compiled release and other runtime necessities
FROM ${RUNNER_IMAGE}

RUN apt-get update -y && apt-get install -y libstdc++6 openssl libncurses5 locales \
  && apt-get clean && rm -f /var/lib/apt/lists/*_*

# Set the locale
RUN sed -i '/en_US.UTF-8/s/^# //g' /etc/locale.gen && locale-gen

ENV LANG en_US.UTF-8
ENV LANGUAGE en_US:en
ENV LC_ALL en_US.UTF-8

WORKDIR "/app"
RUN chown nobody /app

ENV MIX_ENV="prod"
ENV PHX_SERVER="true"

COPY --from=builder --chown=nobody:root /app/_build/${MIX_ENV}/rel/backend ./

USER nobody

CMD /app/bin/backend eval "Backend.Release.migrate" && /app/bin/server