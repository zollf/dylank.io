defmodule Web.Views.Changeset do
  use Web, :view
  alias Web.Helpers.Errors

  def render("error.json", %{changeset: changeset}) do
    %{errors: Errors.traverse_errors(changeset)}
  end
end
