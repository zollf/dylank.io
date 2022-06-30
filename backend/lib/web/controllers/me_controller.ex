defmodule Web.Controllers.Me do
  use Web, :controller

  alias Web.Controllers

  action_fallback(Web.Controllers.Fallback)

  def index(conn, _params) do
    Controllers.Users.get(conn, %{
      "username" => conn.assigns.current_user.username,
    })
  end

  def change_password(conn, %{"current_password" => current_password, "new_password" => password}) do
    Controllers.Users.reset_password(conn, %{
      "username" => conn.assigns.current_user.username,
      "current_password" => current_password,
      "new_password" => password
    })
  end

  def edit_details(conn, %{"user" => user_params}) do
    Controllers.Users.edit_details(conn, %{
      "username" => conn.assigns.current_user.username,
      "user" => user_params
    })
  end
end
