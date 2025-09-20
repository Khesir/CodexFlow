import { defineConfig, loadEnv } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
// Vite supports loading custom .env files with the --mode flag.
export default defineConfig(({mode}) => {
  const env = loadEnv(mode, process.cwd() + "/env", "")
  return {
  plugins: [react()],
  define: {
    "process.env": env,
  },
  server: {
    proxy: {
      "/api": {
        target: "http://localhost:8080",
        changeOrigin: true,
        secure: true,
      }
    }
  }
}
})
