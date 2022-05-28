defmodule BackendWeb.ChangesetView do
  use BackendWeb, :view
  alias BackendWeb.Helpers.ErrorsHelpers

  def render("error.json", %{changeset: changeset}) do
    %{errors: ErrorsHelpers.traverse_errors(changeset)}
  end
end
