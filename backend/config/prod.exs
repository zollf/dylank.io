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
  secret_key_base: System.get_env("SECRET_KEY_BASE")

# config :backend, Web.Endpoint, cache_static_manifest: "priv/static/cache_manifest.json"
config :logger, level: :info
