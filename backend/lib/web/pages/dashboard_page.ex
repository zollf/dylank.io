defmodule Web.Pages.Dashboard do
  use Web, :live_view
  alias Web.Pages

  alias Web.PageBuilder

  @pages %{
    :home => {Pages.Home, %{}},
    :tags => {Pages.Tags, %{}}
  }

  def mount(%{"page" => page} = _params, _session, socket) do
    with {_id, { module, _page_session }} <- (Enum.find(@pages, :error, fn {key, _} -> Atom.to_string(key) == page end))  do
      page = %PageBuilder{module: module}
      socket = assign(socket, page: page)
      {:ok, socket}
    end
  end

  def mount(%{}, _session, socket) do
    page = %PageBuilder{module: Pages.Home}
    socket = assign(socket, page: page)
    {:ok, socket}
  end

  def render(assigns) do
    ~H"""
    <div class="dashboard">
      <div class="sidebar">
        <div class="top">
          <div class="header">
            <div class="logo"></div>
            <h1>dylank.io</h1>
          </div>
          <nav class="nav">
            <a class="nav-item home" href="/admin">Home</a>
            <a class="nav-item entries">Entries</a>
            <div class="sub-nav">
              <a class="sub-nav-item projects">- Projects</a>
              <a class="sub-nav-item tags" href="/admin/tags">- Tags</a>
            </div>
            <a class="nav-item assets">Assets</a>
            <a class="nav-item users">Users</a>
          </nav>
        </div>
        <div class="bottom">
          <nav class="nav">
            <a class="nav-item profile">Profile</a>
            <a class="nav-item logout" href="/admin/logout">Logout</a>
          </nav>
        </div>
      </div>
      <div class="page">
        <%= render_page(@page.module, assigns) %>
      </div>
    </div>
    """
  end

  defp render_page(module, assigns) do
    {component, component_assigns} = module.render_page(assigns)
    component_assigns = Map.put(component_assigns, :page, assigns.page)
    live_component(component, component_assigns)
  end
end
