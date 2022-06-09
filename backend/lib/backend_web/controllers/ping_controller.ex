defmodule BackendWeb.Controllers.Ping do
  use BackendWeb, :controller
  def index(conn, _params), do: json(conn, %{ping: "pong"})
  def user_ping(conn, _params), do: json(conn, %{user_ping: "user_pong"})

  def admin_ping(conn, _params), do: json(conn, %{admin_ping: "admin_pong"})
  def super_admin_ping(conn, _params), do: json(conn, %{super_admin_ping: "super_admin_pong"})
end
