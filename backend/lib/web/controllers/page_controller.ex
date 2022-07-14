defmodule Web.Controllers.Page do
  use Web, :controller

  alias Web.Views

  def index(conn, _params) do
    if conn.assigns[:current_user] do
      conn
      |> put_view(Views.Page)
      |> render("index.html")
    else
      conn
      |> put_view(Views.Page)
      |> render("login.html")
    end
  end
end
