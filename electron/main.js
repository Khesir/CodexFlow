const {app, BrowserWindow} = require("electron")
const path = require("path")

function createWindow() {
  const win = new BrowserWindow({
    width: 1200,
    height: 800,
    webPreferences: {
      preload: path.join(__dirname, "preload.js")
    }
  });
  win.loadURL("http://localhost:5173");
  // if (process.env.NODE_ENVE === "developmennt") {
    
  // } else {
  //   win.loadFile(path.join(__dirname, "../frontend/dist/index.html"))
  // }
}

function startBackend() {
  if (process.env.NODE_ENV === "development") {
    console.log("Starting backend manually (go run)...");
    backendProcess = spawn("go", ["run", "./cmd/main.go"], {
      cwd: path.join(__dirname, ".."),
      shell: true,
      stdio: "inherit",
    });
  } else {
    console.log("Starting backend executable...");
    const exePath = path.join(process.resourcesPath, "backend", "main.exe");
    backendProcess = spawn(exePath, [], {
      stdio: "inherit",
    });
  }
}

app.whenReady().then(() => {
  createWindow();
  startBackend();
  
  app.on("activate", () => {
    if (BrowserWindow.getAllWindows().length === 0) createWindow();
  })
});
app.on("window-all-closed", () => {
  if (process.platform !== "darwin") app.quit();
});
