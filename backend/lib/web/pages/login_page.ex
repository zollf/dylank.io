defmodule Web.Pages.Login do
  use Web, :live_view
  alias Backend.Models.UserToken
  import Phoenix.LiveView

  def mount(_params, session, socket) do
    user_token_string = session["user_token_string"]
    user_token = user_token_string && UserToken.get_user_by_session_token(user_token_string)
    if user_token && user_token_string do
      {:ok, redirect(socket, to: "/admin")}
    else
      {:ok, socket}
    end
  end

  def render(assigns) do
    ~H"""
    <div class="login">
      <div class="login-block">
        <div class="login-block-left">
        </div>
        <div class="login-block-right">
          <h1>Login</h1>
          <.live_component module={Web.Component.Login} id="login" />
        </div>
      </div>
    </div>
    """
  end
end
