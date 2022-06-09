defmodule BackendWeb.Views.Changeset do
  use BackendWeb, :view
  alias BackendWeb.Helpers.Errors

  def render("error.json", %{changeset: changeset}) do
    %{errors: Errors.traverse_errors(changeset)}
  end
end
