defmodule BackendWeb.UsersView do
  use BackendWeb, :view
  alias BackendWeb.UsersView

  def render("index.json", %{users: users}) do
    %{data: render_many(users, UsersView, "user.json", as: :user)}
  end

  def render("view.json", %{user: user}) do
    %{data: render_one(user, UsersView, "user.json", as: :user)}
  end

  def render("user.json", %{user: user}) do
    %{
      id: user.id,
      username: user.username,
      hashed_password: user.hashed_password,
      last_login: user.last_login,
      failed_attempts: user.failed_attempts,
      email: user.email,
      admin: user.admin,
      super_admin: user.super_admin
    }
  end
end
