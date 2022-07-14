import Config

config :backend,
  ecto_repos: [Backend.Repo]

config :backend, Web.Endpoint,
  url: [host: "localhost", path: "/admin"],
  render_errors: [view: Web.Views.Error, accepts: ~w(html json), layout: false],
  pubsub_server: Backend.PubSub,
  live_view: [signing_salt: "ROhEx8NX"]

config :backend, Backend.Mailer, adapter: Swoosh.Adapters.Local

# Swoosh API client is needed for adapters other than SMTP.
config :swoosh, :api_client, false

config :logger, :console,
  format: "$time $metadata[$level] $message\n",
  metadata: [:request_id]

config :phoenix, :json_library, Jason

import_config "#{config_env()}.exs"
