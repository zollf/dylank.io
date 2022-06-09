defmodule BackendWeb.Views.UserToken do
  use BackendWeb, :view
  alias BackendWeb.Views

  def render("view.json", %{user_token: user_token}) do
    %{data: render_one(user_token, Views.UserToken, "user_token.json", as: :user_token)}
  end

  def render("user_token.json", %{user_token: user_token}) do
    %{
      token: user_token.token,
      last_used: user_token.last_used,
      inserted_at: user_token.inserted_at,
      user: render_one(user_token.user, Views.Users, "user.json", as: :user)
    }
  end
end
