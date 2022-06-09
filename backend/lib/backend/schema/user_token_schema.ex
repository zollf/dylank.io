defmodule Backend.Schema.UserToken do
  use Ecto.Schema
  import Ecto.Changeset

  alias Backend.Schema.UserToken
  alias Backend.Schema.User
  alias Backend.Repo
  alias Backend.Helpers.StringHelpers

  @rand_size 32

  schema "users_tokens" do
    field :token, :string
    field :last_used, :naive_datetime
    belongs_to :user, User

    timestamps(updated_at: false)
  end

  def create_changeset(user) do
    %UserToken{}
    |> Repo.preload(:user)
    |> change()
    |> put_assoc(:user, user)
    |> put_token()
    |> validate_required([:user, :token])
    |> unique_constraint(:token)
  end

  def update_last_used(user_token) do
    user_token
    |> Repo.preload(:user)
    |> change()
    |> put_new_last_used()
    |> validate_required([:user, :token])
  end

  defp put_new_last_used(changeset) do
    changeset
    |> put_change(:last_used, NaiveDateTime.utc_now() |> NaiveDateTime.truncate(:second))
  end

  defp put_token(changeset) do
    changeset
    |> put_change(:token, StringHelpers.random_string(@rand_size))
  end
end
