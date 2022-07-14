import Config

if System.get_env("PHX_SERVER") do
  config :backend, Web.Endpoint, server: true
end

if config_env() == :prod do
  config :backend, Backend.Repo,
    username: System.get_env("MYSQL_USER"),
    password: System.get_env("MYSQL_PASSWORD"),
    hostname: System.get_env("MYSQL_HOSTNAME"),
    database: System.get_env("MYSQL_DATABASE")

  config :backend, Web.Endpoint,
    secret_key_base: System.get_env("SECRET_KEY_BASE")
end
