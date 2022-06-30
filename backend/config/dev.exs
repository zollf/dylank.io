import Config

config :backend, Backend.Repo,
  username: System.get_env("MYSQL_USER"),
  password: System.get_env("MYSQL_PASSWORD"),
  hostname: System.get_env("MYSQL_HOSTNAME"),
  database: System.get_env("MYSQL_DATABASE"),
  stacktrace: true,
  show_sensitive_data_on_connection_error: true,
  pool_size: 10

config :backend, Web.Endpoint,
  http: [ip: {0, 0, 0, 0}, port: 8080],
  check_origin: false,
  code_reloader: true,
  debug_errors: true,
  secret_key_base: "HW9WaQ9yb3r56tvS2vC6UJCveq+UZNHKm5efu3OLR2I4arnbibCvdRIMoT+/OCeF",
  watchers: []

config :backend, Web.Endpoint,
  live_reload: [
    patterns: [
      ~r"priv/static/.*(js|css|png|jpeg|jpg|gif|svg)$",
      ~r"priv/gettext/.*(po)$",
      ~r"lib/backend_web/(live|views)/.*(ex)$",
      ~r"lib/backend_web/templates/.*(eex)$"
    ]
  ]

config :logger, :console, format: "[$level] $message\n"
config :phoenix, :stacktrace_depth, 20
config :phoenix, :plug_init_mode, :runtime
