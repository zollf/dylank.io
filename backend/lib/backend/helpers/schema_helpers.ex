defmodule Backend.Helpers.SchemaHelpers do
  alias Backend.Helpers.StringHelpers
  import Ecto.Changeset

  def slugify_field(changeset, key_to_use, key_to_change) do
    if value = get_change(changeset, key_to_use) do
      put_change(changeset, key_to_change, StringHelpers.slugify(value))
    else
      changeset
    end
  end
end
