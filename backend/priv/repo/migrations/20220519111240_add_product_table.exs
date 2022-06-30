defmodule Backend.Repo.Migrations.AddProductTable do
  use Ecto.Migration

  def change do
    create table(:projects) do
      add :title, :string
      add :slug, :string
      add :short_description, :string
      add :page_content, :string
      add :preview_link, :string
      add :git_link, :string

      timestamps()
    end

    create(unique_index(:projects, [:slug]))

    create table(:project_tag) do
      add :project_id, references(:projects)
      add :tag_id, references(:tags)

      timestamps()
    end
  end
end
