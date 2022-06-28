defmodule Web.Controllers.Page do
  use Web, :controller

  alias Web.Views

  def index(conn, _params) do
    conn
    |> put_view(Views.Page)
    |> render("index.html")
  end
end
