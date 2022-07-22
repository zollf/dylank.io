defmodule Web.Helpers.Component do
  def validate_required(params, list) do
    case Enum.find(list, fn param -> !Map.has_key?(params, param) end) do
      nil -> :ok
      key -> raise ArgumentError, "Param #{inspect(key)} is required"
    end
    params
  end
end
