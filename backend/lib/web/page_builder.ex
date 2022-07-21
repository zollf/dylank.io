defmodule Web.PageBuilder do
  defstruct module: nil

  @opaque component :: {module, map}

  @callback render_page(assigns :: Socket.assigns()) :: component()

  def render_page() do

  end

  defmacro __using__(opts) do
    quote bind_quoted: [opts: opts] do
      import Phoenix.LiveView
      import Phoenix.LiveView.Helpers
      import Web.PageBuilder

      @behaviour Web.PageBuilder
      refresher? = Keyword.get(opts, :refresher?, true)

      def __page_live__(:refresher?) do
        unquote(refresher?)
      end

      def init(opts), do: {:ok, opts}
      defoverridable init: 1
    end
  end
end
