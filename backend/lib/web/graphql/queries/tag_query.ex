defmodule Web.GraphQL.Queries.Tag do
  use Absinthe.Schema.Notation

  import_types Web.GraphQL.Types.Tag

  alias Web.GraphQL.Resolvers

  object :tag_queries do
    @desc "Get all tags"
    field :tags, list_of(:tag), resolve: &Resolvers.Tag.list_tags/3
  end
end
