defmodule BackendWeb.PingController do
  use BackendWeb, :controller
  def index(conn, _params), do: json(conn, %{ping: "pong"})
end
