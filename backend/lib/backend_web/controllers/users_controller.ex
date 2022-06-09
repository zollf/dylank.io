defmodule BackendWeb.Controllers.Users do
  use BackendWeb, :controller

  alias Backend.Models.User
  alias BackendWeb.Views

  action_fallback(BackendWeb.Controllers.Fallback)

  def index(conn, _params) do
    users = User.get_users()
    conn
    |> put_view(Views.Users)
    |> render("index.json", users: users)
  end

  def create(conn, %{"user" => user_params}) do
    with {:ok, user} <- User.create_user(user_params) do
      conn
      |> put_view(Views.Users)
      |> render("view.json", user: user)
    end
  end

  def edit_details(conn, %{"username" => username, "user" => user_params}) do
    with {:ok, user} <- User.edit_user_details(user_params, username) do
      conn
      |> put_view(Views.Users)
      |> render("view.json", user: user)
    end
  end

  def delete(conn, %{"username" => username}) do
    with {:ok, _user} <- User.delete_user(username) do
      json(conn, %{success: true})
    end
  end

  def get(conn, %{"username" => username}) do
    with {:ok, user} <- User.get_user(username) do
      conn
      |> put_view(Views.Users)
      |> render("view.json", user: user)
    end
  end

  def reset_password(conn, %{"username" => username, "current_password" => current_password, "new_password" => password}) do
    with {:ok, user} <- User.edit_user_password(username, current_password, password) do
      conn
      |> put_view(Views.Users)
      |> render("view.json", user: user)
    end
  end
end
