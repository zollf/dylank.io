defmodule Backend.Repo do
  use Ecto.Repo,
    otp_app: :backend,
    adapter: Ecto.Adapters.MyXQL

  def init(_, opts) do
    case Application.get_env(:backend, :env) do
      :test ->
        {:ok, opts}

      _ ->
        opts =
          opts
          |> Keyword.put(:username, System.get_env("MYSQL_USER"))
          |> Keyword.put(:password, System.get_env("MYSQL_PASSWORD"))
          |> Keyword.put(:database, System.get_env("MYSQL_DATABASE"))
          |> Keyword.put(:hostname, System.get_env("MYSQL_HOSTNAME"))

        {:ok, opts}
    end
  end
end
