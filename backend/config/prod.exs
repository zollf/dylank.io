import Config

config :backend, Backend.Repo,
  stacktrace: true,
  show_sensitive_data_on_connection_error: false,
  pool_size: 10

config :backend, Web.Endpoint,
  http: [ip: {0, 0, 0, 0}, port: 8080],
  server: true,
  load_from_system_env: true,
  url: [host: "dylank.io", path: "/admin"]

# config :backend, Web.Endpoint, cache_static_manifest: "priv/static/cache_manifest.json"
config :logger, level: :info
