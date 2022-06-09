defmodule Backend.Repo.Migrations.AddUserSessionTokens do
  use Ecto.Migration

  def change do
    create table(:users_tokens) do
      add :user_id, references(:users, on_delete: :delete_all), null: false
      add :token, :string, null: false
      add :last_used, :naive_datetime

      timestamps(updated_at: false)
    end

    create index(:users_tokens, [:user_id])
    create unique_index(:users_tokens, [:token])
  end
end
