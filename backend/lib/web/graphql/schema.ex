defmodule Web.GraphQL.Schema do
  use Absinthe.Schema

  import_types Web.GraphQL.Queries.Tag
  import_types Web.GraphQL.Queries.Project

  query do
    import_fields :tag_queries
    import_fields :project_queries
  end
end
