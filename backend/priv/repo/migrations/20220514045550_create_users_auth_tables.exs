defmodule Backend.Repo.Migrations.CreateUsersAuthTables do
  use Ecto.Migration

  def change do
    create table(:users) do
      add :username, :string
      add :email, :string, null: false, size: 160
      add :hashed_password, :string, null: false
      add :last_login, :naive_datetime
      add :failed_attempts, :integer
      add :admin, :boolean, null: false, default: false
      add :super_admin, :boolean, null: false, default: false

      timestamps()
    end

    create unique_index(:users, [:email])
    create unique_index(:users, [:username])
  end
end
