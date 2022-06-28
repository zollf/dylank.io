defmodule Web.GraphQL.Types.Tag do
  use Absinthe.Schema.Notation

  object :tag do
    field :id, :id
    field :title, :string
    field :slug, :string
    field :updated_at, non_null(:string)
    field :inserted_at, non_null(:string)
  end

  object :tag_interface do
    field :id, :id
    field :title, :string
    field :slug, :string
    field :updated_at, non_null(:string)
    field :inserted_at, non_null(:string)
    field :count, :integer
  end
end
