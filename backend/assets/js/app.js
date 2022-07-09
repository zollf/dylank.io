import "phoenix_html"
import "../scss/app.scss"

import { LiveSocket } from "phoenix_live_view"
import { Socket } from "phoenix"

let socketPath = document.querySelector("html").getAttribute("phx-socket") || "/live";
let csrfToken = document.querySelector("meta[name='csrf-token']").getAttribute("content");
let liveSocket = new LiveSocket(socketPath, Socket, {
  params: (liveViewName) => {
    return {
      _csrf_token: csrfToken,
    }
  }
});

liveSocket.connect()

window.liveSocket = liveSocket