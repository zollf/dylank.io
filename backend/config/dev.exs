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
  secret_key_base: System.get_env("SECRET_KEY_BASE"),
  watchers: [
    node: [
      "node_modules/webpack/bin/webpack.js",
      "watch",
      "--mode",
      System.get_env("NODE_ENV") || "production",
      "--watch-options-stdin",
      cd: "assets"
    ]
  ]

config :backend, Web.Endpoint,
  live_reload: [
    patterns: [
      ~r"priv/static/.*(js|css|png|jpeg|jpg|gif|svg)$",
      ~r"lib/web/(live|views)/.*(ex)$",
      ~r"lib/web/templates/.*(eex)$"
    ]
  ]

config :logger, :console, format: "[$level] $message\n"
config :phoenix, :stacktrace_depth, 20
config :phoenix, :plug_init_mode, :runtime
