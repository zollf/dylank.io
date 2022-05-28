import Config

config :backend, Backend.Repo,
  username: "mysql",
  password: "mysql",
  hostname: "mysql",
  database: "db",
  stacktrace: true,
  show_sensitive_data_on_connection_error: true,
  pool_size: 10

config :backend, BackendWeb.Endpoint,
  # Binding to loopback ipv4 address prevents access from other machines.
  # Change to `ip: {0, 0, 0, 0}` to allow access from other machines.
  http: [ip: {0, 0, 0, 0}, port: 8080],
  check_origin: false,
  code_reloader: true,
  debug_errors: true,
  secret_key_base: "HW9WaQ9yb3r56tvS2vC6UJCveq+UZNHKm5efu3OLR2I4arnbibCvdRIMoT+/OCeF",
  watchers: [
    # Start the esbuild watcher by calling Esbuild.install_and_run(:default, args)
    # esbuild: {Esbuild, :install_and_run, [:default, ~w(--sourcemap=inline --watch)]}
  ]

config :backend, BackendWeb.Endpoint,
  live_reload: [
    patterns: [
      ~r"priv/static/.*(js|css|png|jpeg|jpg|gif|svg)$",
      ~r"priv/gettext/.*(po)$",
      ~r"lib/backend_web/(live|views)/.*(ex)$",
      ~r"lib/backend_web/templates/.*(eex)$"
    ]
  ]

# Do not include metadata nor timestamps in development logs
config :logger, :console, format: "[$level] $message\n"

# Set a higher stacktrace during development. Avoid configuring such
# in production as building large stacktraces may be expensive.
config :phoenix, :stacktrace_depth, 20

# Initialize plugs at runtime for faster development compilation
config :phoenix, :plug_init_mode, :runtime
