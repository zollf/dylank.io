defmodule Backend.Schema.Project do
  use Ecto.Schema
  import Ecto.Changeset
  import Ecto.Query

  alias Backend.Schema.Project
  alias Backend.Helpers.SchemaHelpers
  alias Backend.Schema.Tag
  alias Backend.Repo

  schema "projects" do
    field :title, :string
    field :slug, :string
    field :short_description, :string
    field :page_content, :string
    field :preview_link, :string
    field :git_link, :string
    many_to_many :tags, Tag,
      join_through: "project_tag",
      on_replace: :delete

    timestamps()
  end

  def changeset(%Project{} = project, attrs) do
    project
    |> Repo.preload(:tags)
    |> cast(attrs, [:title, :slug, :short_description, :page_content, :preview_link, :git_link])
    |> put_assoc(:tags, get_assoc(attrs["tags"] || []))
    |> validate_required([:title, :short_description])
    |> unique_constraint(:slug)
    |> SchemaHelpers.slugify_field(:title, :slug)
  end


  defp get_assoc([]), do: []
  defp get_assoc(slugs) do
    Repo.all(from t in Tag, where: t.slug in ^slugs)
  end
end
