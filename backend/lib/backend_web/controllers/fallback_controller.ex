defmodule BackendWeb.FallbackController do
  use BackendWeb, :controller

  @doc """
  Error call
  """
  def call(conn, {:error, %Ecto.Changeset{} = changeset}) do
    conn
    |> put_status(:unprocessable_entity)
    |> put_view(BackendWeb.ChangesetView)
    |> render("error.json", changeset: changeset)
  end

  def call(conn, {:not_found, msg}) do
    conn
    |> put_status(:unprocessable_entity)
    |> json(%{error: msg})
  end
end
