defmodule BackendWeb.ProjectsController do
  use BackendWeb, :controller
  alias Backend.Model.Project

  action_fallback(BackendWeb.FallbackController)

  def index(conn, _params) do
    conn
    |> put_view(BackendWeb.ProjectsView)
    |> render("index.json", projects: Project.get_projects())
  end
  def create(conn, %{"project" => project_params}) do
    with {:ok, project} <- Project.create_project(project_params) do
      conn
      |> put_view(BackendWeb.ProjectsView)
      |> render("view.json", project: project)
    end
  end

  def edit(conn, %{"project" => project_params, "slug" => slug}) do
    with {:ok, project} <- Project.edit_project(project_params, slug) do
      conn
      |> put_view(BackendWeb.ProjectsView)
      |> render("view.json", project: project)
    end
  end

  def delete(conn, %{"slug" => slug}) do
    with {:ok, _tag} <- Project.delete_project(slug) do
      json(conn, %{success: true})
    end
  end

  def get(conn, %{"slug" => slug}) do
    with {:ok, project} <- Project.get_project(slug) do
      conn
      |> put_view(BackendWeb.ProjectsView)
      |> render("view.json", project: project)
    end
  end
end
