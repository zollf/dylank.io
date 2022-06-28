defmodule Web.GraphQL.Resolvers.Tag do
  alias Backend.Models.Tag

  def list_tags(_parent, _args, _resolution) do
    {:ok, Tag.get_tags()}
  end
end
