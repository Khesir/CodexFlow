const {app, BrowserWindow, ipcMain} = require("electron")
const path = require("path")
const dotenv = require("dotenv")

dotenv.config({ path: path.join(__dirname, "../env/.env.electron")})

let backendProcess
let loginWindow;
let mainWindow;

function startBackend() {
  if (process.env.NODE_ENV === "production") { 
    console.log("Starting backend executable...");
    const exePath = path.join(process.resourcesPath, "backend", "main.exe");
    backendProcess = spawn(exePath, [], {
      stdio: "inherit",
    });
    backendProcess.on("error", (err)=> {
      console.error("Faild to start backend: ", err)
    })
    backendProcess.on("exit", (code) => {
      console.log(`Backend exited with code ${code}`)
    })
  }
}

// Create the login window
function createLoginWindow() {
  loginWindow = new BrowserWindow({
    width: 500,
    height: 600,
    webPreferences: {
      preload: path.join(__dirname, "preload.js"),
    }
  });

  if (process.env.NODE_ENV === "development"){
    loginWindow.loadURL(process.env.FE_URL).catch(console.error);
  } else {
    loginWindow.loadFile(path.join(__dirname, "../frontend/dist/index.html"))
  }

  loginWindow.webContents.once("did-finish-load", () => {
    loginWindow.webContents.send("navigate","/login")
  })

  loginWindow.on("closed" , () => {
    loginWindow = null;
  })
}

// Create the main window
function createMainWindow() {
   mainWindow = new BrowserWindow({
    width: 1200,
    height: 800,
    webPreferences: {
      preload: path.join(__dirname, "preload.js")
    }
  });
  
  if (process.env.NODE_ENV === "development") {
    mainWindow.loadURL("http://localhost:5173").catch(console.error);
  } else {
    mainWindow.loadFile(path.join(__dirname, "../frontend/dist/index.html"))
  }

  mainWindow.on("closed", () => {
    mainWindow = null;
  })
}

ipcMain.on("login-success", () => {
  if (loginWindow) loginWindow.close();
  createMainWindow();
})

app.whenReady().then(() => {
  startBackend();
  createLoginWindow();
  
  app.on("activate", () => {
    if (BrowserWindow.getAllWindows().length === 0) createWindow();
  })
});

app.on("window-all-closed", () => {
  if (process.platform !== "darwin") app.quit();
});
