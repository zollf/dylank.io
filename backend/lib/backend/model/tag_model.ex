defmodule Backend.Models.Tag do
  alias Backend.Schema.Tag
  alias Backend.Repo

  def create_tag(attrs \\ %{}) do
    %Tag{}
    |> Tag.create_changeset(attrs)
    |> Repo.insert()
  end

  def delete_tag(slug) do
    case Repo.get_by(Tag, slug: slug) do
      nil ->  {:not_found, "Cannot find tag to delete"}
      tag -> Repo.delete tag
    end
  end

  def edit_tag(attrs \\ %{}, slug) do
    case Repo.get_by(Tag, slug: slug) do
      nil -> {:not_found, "Cannot find tag to edit"}
      tag -> tag
      |> Tag.create_changeset(attrs)
      |> Repo.update()
    end
  end

  def get_tag(slug) do
    case Repo.get_by(Tag, slug: slug) do
      nil -> {:not_found, "Cannot find tag"}
      tag -> {:ok, tag}
    end
  end

  def get_tags(), do: Repo.all(Tag)
end
