defmodule Backend.Models.UserToken do
  alias Backend.Schema.UserToken
  alias Backend.Repo

  def create_user_token(user) do
    user
    |> UserToken.create_changeset()
    |> Repo.insert()
  end

  def get_user_by_session_token(user_token_string) do
    case Repo.get_by(UserToken, token: user_token_string) do
      nil -> nil
      user_token -> user_token
      |> update_last_used()
      |> Repo.preload(:user)
    end
  end

  def delete_user_token(user_token_string) do
    case Repo.get_by(UserToken, token: user_token_string) do
      nil -> {:bad_request, "Bad Request"}
      user_token -> Repo.delete user_token
    end
  end

  defp update_last_used(user_token) do
    user_token
    |> UserToken.update_last_used()
    |> Repo.update()
    user_token
  end
end
