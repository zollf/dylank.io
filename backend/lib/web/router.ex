defmodule Web.Router do
  use Web, :router

  import Web.Plugs.Auth

  pipeline :browser do
    plug :accepts, ["html"]
    plug :fetch_session
    plug :fetch_live_flash
    plug :put_root_layout, {Web.Views.Layout, :root}
    plug :protect_from_forgery
    plug :put_secure_browser_headers
    plug :fetch_current_user
  end

  pipeline :dashboard do
    plug :put_layout, {Web.Views.Layout, :dashboard}
  end

  pipeline :api do
    plug :accepts, ["json"]
    plug :fetch_session
    plug :fetch_current_user
  end

  scope "/", Web do
    pipe_through :browser

    live "/login", Pages.Login
    get "/logout", Controllers.Auth, :admin_logout
  end

  live_session :admin, on_mount: Web.Helpers.AuthLive do
    scope "/", Web do
      pipe_through [:browser, :dashboard]

      live "/", Pages.Dashboard
      live "/:page", Pages.Dashboard
    end
  end

  scope "/api" do
    pipe_through :api

    forward "/graphiql", Absinthe.Plug.GraphiQL,
      schema: Web.GraphQL.Schema

    forward "/graphql", Absinthe.Plug,
      schema: Web.GraphQL.Schema
  end

  scope "/api", Web do
    pipe_through :api

    get "/ping", Controllers.Ping, :index

    post "/auth/login", Controllers.Auth, :login

    get "/tags", Controllers.Tags, :index
    get "/tags/:slug", Controllers.Tags, :get

    get "/projects", Controllers.Projects, :index
    get "/projects/:slug", Controllers.Projects, :get
  end

  # Authenticated User Scope
  scope "/api", Web do
    pipe_through [:api, :require_authenticated_user]

    get "/user_ping", Controllers.Ping, :user_ping

    get "/auth/session", Controllers.Auth, :session
    get "/auth/logout", Controllers.Auth, :logout

    get "/me", Controllers.Me, :index
    post "/me/change_password", Controllers.Me, :change_password
    post "/me/edit_details", Controllers.Me, :edit_details
  end

  # Admin scope
  scope "/api", Web do
    pipe_through [:api, :require_authenticated_admin]

    get "/admin_ping", Controllers.Ping, :admin_ping

    post "/tags/create", Controllers.Tags, :create
    post "/tags/delete", Controllers.Tags, :delete
    post "/tags/edit", Controllers.Tags, :edit

    post "/projects/create", Controllers.Projects, :create
    post "/projects/delete", Controllers.Projects, :delete
    post "/projects/edit", Controllers.Projects, :edit

    get "/users", Controllers.Users, :index
    get "/users/:username", Controllers.Users, :get
  end

  # Super Admin Scope
  scope "/api", Web do
    pipe_through [:api, :require_authenticated_super_admin]

    get "/super_admin_ping", Controllers.Ping, :super_admin_ping

    post "/users/create", Controllers.Users, :create
    post "/users/delete", Controllers.Users, :delete
    post "/users/edit", Controllers.Users, :edit_details
    post "/users/reset_password", Controllers.Users, :reset_password
  end

  if Mix.env() in [:dev, :test] do
    import Phoenix.LiveDashboard.Router

    scope "/" do
      pipe_through :browser

      live_dashboard "/live_dashboard", metrics: Web.Telemetry
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
