defmodule Web.GraphQL.Resolvers.Project do
  alias Backend.Models
  alias Backend.Schema
  alias Backend.Repo
  import Ecto.Query

  def list_projects(_parent, args, _resolution) do
    projects = Schema.Project
      |> from()
      |> join(:left, [project], project_tag in "project_tag", on: project_tag.project_id == project.id)
      |> join(:left, [_, project_tag], tag in Schema.Tag, on: project_tag.tag_id == tag.id)
      |> maybe_filter_tags(Map.get(args, :tags, []))
      |> maybe_limit(Map.get(args, :limit))
      |> maybe_offset(Map.get(args, :offset))
      |> Repo.all()
      |> Repo.preload(:tags)

    tags = projects
      |> Enum.filter(fn project -> length(project.tags) > 0 end)
      |> Enum.map(fn project -> project.tags end)
      |> List.flatten()
      |> add_count()
      |> Enum.uniq()

    {:ok, %{
      items: projects,
      total: Models.Project.get_total_projects_count(),
      tags: tags,
      items_total: length(projects)
    }}
  end

  defp count_tags(tags, tag_to_count) do
    Enum.count(tags, fn tag -> tag.id === tag_to_count.id end)
  end

  defp add_count(tags) do
    Enum.map(tags, fn tag -> Map.merge(tag, %{count: count_tags(tags, tag)}) end)
  end

  defp maybe_filter_tags(queryable, []), do: queryable
  defp maybe_filter_tags(queryable, tags), do: queryable |> where([_, _, tag], tag.slug in ^tags)

  defp maybe_offset(queryable, nil), do: queryable
  defp maybe_offset(queryable, offset), do: queryable |> offset(^offset)

  defp maybe_limit(queryable, nil), do: queryable
  defp maybe_limit(queryable, limit), do: queryable |> limit(^limit)
end
