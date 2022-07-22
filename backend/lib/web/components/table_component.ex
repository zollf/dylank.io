defmodule Web.Components.Table do
  use Web, :live_component

  alias Web.Helpers

  @impl true
  def mount(socket) do
    {:ok, socket}
  end

  @impl true
  def render(assigns) do
    ~H"""
    <div class="table-component">
      <div class="top">
        <h2><%= @title %></h2>
      </div>
      <div class="actions">
        <button class="large">New Entry</button>
      </div>
      <div class="table">
        <table>
          <thead>
            <tr>
              <%= for column <- @columns do %>
                <th>
                  <%= column.header %>
                </th>
              <% end %>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            <%= for row <- @data do %>
              <tr>
                <%= for column <- @columns do %>
                  <td>
                    <%= row[column.field] %>
                  </td>
                <% end %>
                <td class="action-buttons">
                  <button>Edit</button>
                  <button>Delete</button>
                </td>
              </tr>
            <% end %>
          </tbody>
        </table>
      </div>
    </div>
    """
  end

  def normalize_params(params) do
    params
    |> Helpers.Component.validate_required([:title, :columns, :fetch_data])
  end

  @impl true
  def update(assigns, socket) do
    %{fetch_data: fetch_data} = assigns
    {data, total} = fetch_data.()
    assigns = Map.merge(assigns, %{data: data, total: total})
    {:ok, assign(socket, assigns)}
  end
end
