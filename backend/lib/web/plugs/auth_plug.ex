defmodule Web.Plugs.Auth do
  use Web, :controller

  import Plug.Conn

  alias Backend.Models.UserToken

  @remember_me_cookie "user_remember_me"

  def fetch_current_user(conn, _opts) do
    {user_token_string, conn} = ensure_user_token(conn)
    user_token = user_token_string && UserToken.get_user_by_session_token(user_token_string)
    if user_token && user_token_string do
      assign(conn, :current_user, user_token.user)
    else
      assign(conn, :current_user, nil)
    end
  end

  def require_authenticated_user(conn, _opts) do
    if conn.assigns[:current_user] do
      conn
    else
      conn |> unauthorized_and_halt()
    end
  end

  def require_authenticated_admin(conn, _opts) do
    if conn.assigns[:current_user] && conn.assigns.current_user.admin do
      conn
    else
      conn |> unauthorized_and_halt()
    end
  end

  def require_authenticated_super_admin(conn, _opts) do
    if conn.assigns[:current_user] && conn.assigns.current_user.super_admin do
      conn
    else
      conn |> unauthorized_and_halt()
    end
  end

  defp ensure_user_token(conn) do
    if user_token_string = get_session(conn, :user_token_string) do
      {user_token_string, conn}
    else
      conn = fetch_cookies(conn, signed: [@remember_me_cookie])

      if user_token_string = conn.cookies[@remember_me_cookie] do
        {user_token_string, put_session(conn, :user_token_string, user_token_string)}
      else
        {nil, conn}
      end
    end
  end

  defp unauthorized_and_halt(conn) do
    conn
    |> put_status(:unauthorized)
    |> json(%{error: "unauthorized"})
    |> halt()
  end
end
