defmodule BackendWeb.Views.Tags do
  use BackendWeb, :view
  alias BackendWeb.Views

  def render("index.json", %{tags: tags}) do
    %{data: render_many(tags, Views.Tags, "tag.json", as: :tag)}
  end

  def render("view.json", %{tag: tag}) do
    %{data: render_one(tag, Views.Tags, "tag.json", as: :tag)}
  end

  def render("tag.json", %{tag: tag}) do
    %{
      id: tag.id,
      title: tag.title,
      slug: tag.slug,
    }
  end
end
