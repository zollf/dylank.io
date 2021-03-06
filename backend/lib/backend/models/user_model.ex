defmodule Backend.Models.User do
  alias Backend.Repo
  alias Backend.Schema.User

  def create_user(attrs \\ %{}) do
    %User{}
    |> User.create_changeset(attrs)
    |> Repo.insert()
  end

  def get_users(), do: Repo.all(User)

  def get_user(username) do
    case Repo.get_by(User, username: username) do
      nil -> {:not_found, "Cannot find user"}
      user -> {:ok, user}
    end
  end

  def delete_user(username) do
    case Repo.get_by(User, username: username) do
      nil -> {:not_found, "Cannot find user to delete"}
      user -> Repo.delete user
    end
  end

  def edit_user_details(attrs \\ %{}, username) do
    case Repo.get_by(User, username: username) do
      nil -> {:not_found, "Cannot get user to edit details for"}
      user -> user
      |> User.update_changeset(attrs)
      |> Repo.update()
    end
  end

  def edit_user_password(username, current_password, new_password) do
    case Repo.get_by(User, username: username) do
      nil -> {:not_found, "Cannot get user to change password for"}
      user -> user
      |> User.update_password_changeset(%{current_password: current_password, password: new_password})
      |> Repo.update()
    end
  end

  def get_user_with_password(username, password) do
    case Repo.get_by(User, username: username) do
      nil -> {:unauthorized, "Username or password is incorrect"}
      user -> case verify_user_password?(user, password) do
        true -> {:ok, user}
        false -> {:unauthorized, "Username or password is incorrect"}
      end
    end
  end

  defp verify_user_password?(user, password) do
    hashed_password = user.hashed_password
    if Bcrypt.verify_pass(password, hashed_password), do: true, else: false
  end
end
