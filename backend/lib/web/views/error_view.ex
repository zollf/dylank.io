defmodule Web.Views.Error do
  use Web, :view

  def render("500.html", assigns) do
    IO.inspect(assigns)
    "Internal Server Error"
  end
end
