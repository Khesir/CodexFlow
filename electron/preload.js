const {contextBridge} = require("electron");

contextBridge.exposeInMainWorld("codexfloa", {
  ping: () => "pong  from preload",
})
