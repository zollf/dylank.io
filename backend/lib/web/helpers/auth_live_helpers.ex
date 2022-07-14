defmodule Web.Helpers.AuthLive do
  alias Backend.Models.UserToken
  import Phoenix.LiveView

  def on_mount(:default, _params, session, socket) do
    user_token_string = session["user_token_string"]
    user_token = user_token_string && UserToken.get_user_by_session_token(user_token_string)

    if user_token && user_token_string do
      {:cont, assign(socket, :current_user, user_token.user)}
    else
      {:halt, redirect(socket, to: "/admin/login")}
    end
  end
end
