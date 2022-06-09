defmodule Backend.Schema.Tag do
  use Ecto.Schema
  import Ecto.Changeset

  alias Backend.Schema.Tag
  alias Backend.Helpers.SchemaHelpers

  schema "tags" do
    field :slug, :string
    field :title, :string

    timestamps()
  end

  @doc """
  Returns a Tag changest for new tags
  """
  @spec create_changeset(%Tag{}, map) :: Ecto.Changeset.t()
  def create_changeset(%Tag{} = tag, attrs) do
    tag
    |> cast(attrs, [:title, :slug])
    |> validate_required([:title])
    |> unique_constraint(:slug)
    |> SchemaHelpers.slugify_field(:title, :slug)
  end
end
