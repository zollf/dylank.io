defmodule Web.Components.Null do
  use Web, :live_component

  @impl true
  def mount(socket) do
    {:ok, socket}
  end

  @impl true
  def render(assigns) do
    ~H"""
    <div>Hello World</div>
    """
  end
end
