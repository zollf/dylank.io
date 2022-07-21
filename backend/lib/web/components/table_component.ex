defmodule Web.Components.Table do
  use Web, :live_component

  @impl true
  def mount(socket) do
    {:ok, socket}
  end

  @impl true
  def render(assigns) do
    ~H"""
    <div>
    Table Component
    </div>
    """
  end
end
