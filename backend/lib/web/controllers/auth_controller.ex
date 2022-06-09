defmodule Web.Controllers.Auth do
  use Web, :controller

  import Plug.Conn

  alias Backend.Models.User
  alias Backend.Models.UserToken
  alias Web.Views

  @max_age 60 * 60 * 24 * 60
  @remember_me_cookie "user_remember_me"
  @remember_me_options [sign: true, max_age: @max_age]

  action_fallback(Web.Controllers.Fallback)

  def login(conn, %{"username" => username, "password" => password, "remember_me" => remember_me}) do
    with {:ok, user} <- User.get_user_with_password(username, password) do
      with {:ok, user_token} <- UserToken.create_user_token(user) do
        conn
        |> renew_session()
        |> put_session(:user_token, user_token.token)
        |> maybe_write_remember_me_cookie(user_token.token, remember_me)
        |> json(%{"success" => true})
      end
    end
  end

  def session(conn, _params) do
    user_token_string = get_session(conn, :user_token_string)
    user_token = user_token_string && UserToken.get_user_by_session_token(user_token_string)
    conn
    |> put_view(Views.UserToken)
    |> render("view.json", user_token: user_token)
  end

  def logout(conn, _params) do
    user_token_string = get_session(conn, :user_token_string)
    user_token_string && UserToken.delete_user_token(user_token_string)
    conn
    |> renew_session()
    |> delete_resp_cookie(@remember_me_cookie)
    |> json(%{"success" => true})
  end

  defp renew_session(conn) do
    conn
    |> configure_session(renew: true)
    |> clear_session()
  end

  defp maybe_write_remember_me_cookie(conn, token, true) do
    put_resp_cookie(conn, @remember_me_cookie, token, @remember_me_options)
  end

  defp maybe_write_remember_me_cookie(conn, _token, false), do: conn
end
