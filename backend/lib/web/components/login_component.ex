defmodule Web.Component.Login do
  use Web, :live_component

  alias Backend.Models

  def mount(socket) do
    {:ok, assign(socket, %{
      incorrect_login: false,
      trigger_submit: false,
      form_params: %{}
    })}
  end

  def render(assigns) do
    ~H"""
    <div>
      <.form
        let={f}
        for={:user}
        action={Routes.auth_path(@socket, :login)}
        phx-submit="save"
        phx-target={@myself}
        phx-trigger-action={@trigger_submit}
        phx-change="validate"
      >
        <%= text_input f, :username, placeholder: "Username", class: "default", value: @form_params["username"], required: true %>
        <%= password_input f, :password, placeholder: "Password", class: "default", value: @form_params["password"], required: true %>

        <div class="checkbox-wrapper">
          <label class="checkbox">
            <%= checkbox f, :remember_me, class: "default", value: @form_params["remember_me"] %>
            <span class="checkmark"></span>
          </label>
          <div>Remember Me</div>
        </div>

        <%= if @incorrect_login do %>
          <div class="error">Username or password is incorrect</div>
        <% end %>
        <%= submit "Login", class: "large" %>
      </.form>
    </div>
    """
  end

  def handle_event("save", %{"user" => %{"username" => username, "password" => password, "remember_me" => _remember_me}}, socket) do
    case Models.User.get_user_with_password(username, password) do
      {:ok, _user} ->{:noreply, assign(socket, trigger_submit: true)}
      {:unauthorized, _} -> {:noreply, assign(socket, incorrect_login: true)}
    end
  end

  def handle_event("validate", %{"user" => params}, socket) do
    {:noreply, assign(socket, :form_params, params)}
  end
end
