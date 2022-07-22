defmodule Web.Pages.Home do
  @moduledoc false
  use Web.PageBuilder

  alias Web.Components

  @impl true
  def render_page(assigns) do
    {Components.Null, assigns}
  end
end
