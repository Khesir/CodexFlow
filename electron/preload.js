const {contextBridge, ipcRenderer} = require("electron");

contextBridge.exposeInMainWorld("codexflow", {
  ping: () => "pong  from preload",
})


contextBridge.exposeInMainWorld("electronAPI", {
  loginSuccess: () => ipcRenderer.send("login-success")
})