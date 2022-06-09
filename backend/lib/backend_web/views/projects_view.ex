defmodule BackendWeb.Views.Projects do
  use BackendWeb, :view
  alias BackendWeb.Views

  def render("index.json", %{projects: projects}) do
    %{data: render_many(projects, Views.Projects, "project.json", as: :project)}
  end

  def render("view.json", %{project: project}) do
    %{data: render_one(project, Views.Projects, "project.json", as: :project)}
  end

  def render("project.json", %{project: project}) do
    %{
      id: project.id,
      title: project.title,
      slug: project.slug,
      tags: render_many(project.tags, Views.Tags, "tag.json", as: :tag)
    }
  end
end
