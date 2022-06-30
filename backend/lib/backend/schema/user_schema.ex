defmodule Backend.Schema.User do
  use Ecto.Schema
  import Ecto.Changeset

  schema "users" do
    field :username, :string
    field :email, :string

    field :current_password, :string, virtual: true # guess at the current_password that needs to be verified
    field :password, :string, virtual: true # non_hashed password that needs to be hashed
    field :hashed_password, :string

    field :last_login, :naive_datetime
    field :failed_attempts, :integer
    field :admin, :boolean
    field :super_admin, :boolean

    timestamps()
  end

  def create_changeset(user, attrs) do
    user
    |> cast(attrs, [:email, :password, :username, :admin, :super_admin])
    |> validate_username()
    |> validate_email()
    |> validate_password()
  end

  def update_changeset(user, attrs) do
    user
    |> cast(attrs, [:email])
    |> validate_username()
    |> validate_email()
  end

  def update_password_changeset(user, attrs) do
    user
    |> cast(attrs, [:current_password, :password, :hashed_password])
    |> verify_password()
    |> validate_password()
  end

  defp validate_username(changeset) do
    changeset
    |> validate_required([:username])
    |> unique_constraint(:username)
  end

  defp validate_email(changeset) do
    changeset
    |> validate_required([:email])
    |> validate_format(:email, ~r/^[^\s]+@[^\s]+$/, message: "must have the @ sign and no spaces")
    |> validate_length(:email, max: 160)
    |> unique_constraint(:email)
  end

  defp validate_password(changeset) do
    changeset
    |> validate_required([:password])
    |> validate_length(:password, min: 1, max: 72)
    # |> validate_format(:password, ~r/[a-z]/, message: "at least one lower case character")
    # |> validate_format(:password, ~r/[A-Z]/, message: "at least one upper case character")
    # |> validate_format(:password, ~r/[!?@#$%^&*_0-9]/, message: "at least one digit or punctuation character")
    |> hash_password()
  end

  defp hash_password(changeset) do
    password = get_change(changeset, :password)

    if password && changeset.valid? do
      changeset
      |> validate_length(:password, max: 72, count: :bytes)
      |> put_change(:hashed_password, Bcrypt.hash_pwd_salt(password))
      |> delete_change(:password)
    else
      changeset
    end
  end

  defp verify_password(changeset) do
    current_password = get_change(changeset, :current_password)
    hashed_password = changeset.data.hashed_password
    if is_binary(hashed_password) && Bcrypt.verify_pass(current_password, hashed_password) do
      changeset
    else
      add_error(changeset, :current_password, "is not valid")
    end
  end
end
