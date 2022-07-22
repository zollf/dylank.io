defmodule Web.Pages.Tags do
  @moduledoc false
  use Web.PageBuilder

  alias Backend.Models

  @impl true
  def render_page(_assigns) do
    table(
      title: "Tags",
      columns: table_columns(),
      fetch_data: &fetch_data/0
    )
  end

  defp table_columns() do
    [
      %{
        header: "# ID",
        field: :id
      },
      %{
        header: "Title",
        field: :title
      },
      %{
        header: "Slug",
        field: :slug
      },
      %{
        header: "Inserted At",
        field: :inserted_at
      }
    ]
  end

  defp fetch_data() do
    data = Models.Tag.get_tags()
    |> Enum.map(&construct_row/1)
    IO.inspect(Models.Tag.get_tags())
    total = length(data)
    {data, total}
  end

  defp construct_row(tag) do
    %{
      id: tag.id,
      title: tag.title,
      slug: tag.slug,
      inserted_at: tag.inserted_at
    }
  end
end
