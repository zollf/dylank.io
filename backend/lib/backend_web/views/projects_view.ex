defmodule BackendWeb.ProjectsView do
  use BackendWeb, :view
  alias BackendWeb.ProjectsView
  alias BackendWeb.TagsView

  def render("index.json", %{projects: projects}) do
    %{data: render_many(projects, ProjectsView, "project.json", as: :project)}
  end

  def render("view.json", %{project: project}) do
    %{data: render_one(project, ProjectsView, "project.json", as: :project)}
  end

  def render("project.json", %{project: project}) do
    %{
      id: project.id,
      title: project.title,
      slug: project.slug,
      tags: render_many(project.tags, TagsView, "tag.json", as: :tag)
    }
  end
end
