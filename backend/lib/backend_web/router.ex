defmodule BackendWeb.Router do
  use BackendWeb, :router

  import BackendWeb.UserAuth

  pipeline :browser do
    plug :accepts, ["html"]
    plug :fetch_session
    plug :fetch_live_flash
    plug :put_root_layout, {BackendWeb.LayoutView, :root}
    plug :protect_from_forgery
    plug :put_secure_browser_headers
  end

  pipeline :api do
    plug :accepts, ["json"]
  end

  scope "/", BackendWeb do
    pipe_through :browser

    get "/", PageController, :index
  end

  scope "/api", BackendWeb do
    pipe_through :api

    get "/ping", PingController, :index

    get "/tags", TagsController, :index
    get "/tags/:slug", TagsController, :get
    post "/tags/create", TagsController, :create
    post "/tags/delete", TagsController, :delete
    post "/tags/edit", TagsController, :edit

    get "/projects", ProjectsController, :index
    get "/projects/:slug", ProjectsController, :get
    post "/projects/create", ProjectsController, :create
    post "/projects/delete", ProjectsController, :delete
    post "/projects/edit", ProjectsController, :edit

    get "/users", UsersController, :index
    get "/users/:username", UsersController, :get
    post "/users/create", UsersController, :create
    post "/users/delete", UsersController, :delete
    post "/users/edit", UsersController, :edit_details
    post "/users/reset_password", UsersController, :reset_password
  end

  if Mix.env() in [:dev, :test] do
    import Phoenix.LiveDashboard.Router

    scope "/" do
      pipe_through :browser

      live_dashboard "/dashboard", metrics: BackendWeb.Telemetry
    end
  end

  # Enables the Swoosh mailbox preview in development.
  #
  # Note that preview only shows emails that were sent by the same
  # node running the Phoenix server.
  if Mix.env() == :dev do
    scope "/dev" do
      pipe_through :browser

      forward "/mailbox", Plug.Swoosh.MailboxPreview
    end
  end
end
