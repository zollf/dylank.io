defmodule Web.Controllers.Fallback do
  use Web, :controller
  alias Web.Views

  @doc """
  Error call
  """
  def call(conn, {:error, %Ecto.Changeset{} = changeset}) do
    conn
    |> put_status(:unprocessable_entity)
    |> put_view(Views.Changeset)
    |> render("error.json", changeset: changeset)
  end

  def call(conn, {:bad_request, msg}) do
    conn
    |> put_status(:bad_request)
    |> json(%{error: msg})
  end

  def call(conn, {:not_found, msg}) do
    conn
    |> put_status(:not_found)
    |> json(%{error: msg})
  end


  def call(conn, {:unauthorized, msg}) do
    conn
    |> put_status(:unauthorized)
    |> json(%{error: msg})
  end
end
