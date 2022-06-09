defmodule BackendWeb.Controllers.Tags do
  use BackendWeb, :controller

  alias Backend.Models.Tag
  alias BackendWeb.Views

  action_fallback(BackendWeb.Controllers.Fallback)

  def index(conn, _params) do
    tags = Tag.get_tags()
    conn
    |> put_view(Views.Tags)
    |> render("index.json", tags: tags)
  end

  def create(conn, %{"tag" => tag_params}) do
    with {:ok, tag} <- Tag.create_tag(tag_params) do
      conn
      |> put_view(Views.Tags)
      |> render("view.json", tag: tag)
    end
  end

  def edit(conn, %{"tag" => tag_params, "slug" => slug}) do
    with {:ok, tag} <- Tag.edit_tag(tag_params, slug) do
      conn
      |> put_view(Views.Tags)
      |> render("view.json", tag: tag)
    end
  end

  def delete(conn, %{"slug" => slug}) do
    with {:ok, _tag} <- Tag.delete_tag(slug) do
      json(conn, %{success: true})
    end
  end

  def get(conn, %{"slug" => slug}) do
    with {:ok, tag} <- Tag.get_tag(slug) do
      conn
      |> put_view(Views.Tags)
      |> render("view.json", tag: tag)
    end
  end
end
