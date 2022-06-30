defmodule Backend.Models.Project do
  alias Backend.Schema.Project
  alias Backend.Repo

  def create_project(attrs \\ %{}) do
    %Project{}
    |> Project.changeset(attrs)
    |> Repo.insert()
  end

  def delete_project(slug) do
    case Repo.get_by(Project, slug: slug) do
      nil -> {:not_found, "Cannot find project to delete"}
      project -> Repo.delete project
    end
  end

  def edit_project(attrs \\ %{}, slug) do
    case Repo.get_by(Project, slug: slug) do
      nil -> {:not_found, "Cannot find project to edit"}
      project -> project
      |> Project.changeset(attrs)
      |> Repo.update()
    end
  end

  def get_project(slug) do
    case Repo.get_by(Project, slug: slug) do
      nil -> {:not_found, "Cannot find project"}
      tag -> {:ok, tag |> Repo.preload(:tags)}
    end
  end

  def get_projects(), do: Repo.all(Project) |> Repo.preload(:tags)

  def get_total_projects_count, do: Repo.aggregate(Project, :count, :id)
end
