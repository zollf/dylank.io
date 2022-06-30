defmodule Web.GraphQL.Queries.Project do
  use Absinthe.Schema.Notation

  import_types Web.GraphQL.Types.Project

  alias Web.GraphQL.Resolvers

  object :project_queries do
    @desc "Get all projects"
    field :projects, :project_interface do
      arg :tags, list_of(:string)
      arg :offset, :integer
      arg :limit, :integer
      resolve &Resolvers.Project.list_projects/3
    end
  end
end
