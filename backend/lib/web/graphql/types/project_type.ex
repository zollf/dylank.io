defmodule Web.GraphQL.Types.Project do
  use Absinthe.Schema.Notation

  object :project do
    field :id, non_null(:id)
    field :title, non_null(:string)
    field :slug, non_null(:string)
    field :short_description, :string
    field :page_content, :string
    field :preview_link, :string
    field :git_link, :string
    field :updated_at, non_null(:string)
    field :inserted_at, non_null(:string)

    field :tags, list_of(:tag)
  end

  object :project_interface do
    field :items, non_null(list_of(:project))
    field :tags, non_null(list_of(:tag_interface))
    field :total, non_null(:integer)
    field :items_total, non_null(:integer)
  end
end
