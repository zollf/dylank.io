defmodule Web.Pages.Dashboard do
  use Web, :live_view

  def mount(socket) do
    {:ok, socket}
  end

  def render(assigns) do
    ~H"""
    <div>
    Hello User
    </div>
    """
  end
end
