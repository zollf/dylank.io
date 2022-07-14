defmodule Web.Views.Error do
  use Web, :view

  def render("500.html", _assigns) do
    "Internal Server Error"
  end

  def render("404.html", _assigns) do
    "Page Not Found"
  end
end
