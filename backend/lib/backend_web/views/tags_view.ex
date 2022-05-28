defmodule BackendWeb.TagsView do
  use BackendWeb, :view
  alias BackendWeb.TagsView

  def render("index.json", %{tags: tags}) do
    %{data: render_many(tags, TagsView, "tag.json", as: :tag)}
  end

  def render("view.json", %{tag: tag}) do
    %{data: render_one(tag, TagsView, "tag.json", as: :tag)}
  end

  def render("tag.json", %{tag: tag}) do
    %{
      id: tag.id,
      title: tag.title,
      slug: tag.slug,
    }
  end
end
