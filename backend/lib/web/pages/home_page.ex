defmodule Web.Pages.Home do
  @moduledoc false
  use Web.PageBuilder

  @impl true
  def render_page(assigns) do
    {Web.Components.Table, assigns}
  end
end
