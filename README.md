# GoAF (Go App Framework)

A modern framework for building cross-platform applications using Go, WebAssembly, and Capacitor. GoAF allows you to write your core application logic in Go, compile it to WebAssembly, and deploy it as a web app, iOS app, or Android app.

## 🚀 Quick Start

1. **Use this template**
   ```bash
   git clone https://github.com/yourusername/GoAF.git my-app
   cd my-app
   ```

2. **Update app configuration**
   - Edit `capacitor/capacitor.config.json`:
     ```json
     {
       "appId": "io.mycompany.myapp",
       "appName": "My App Name"
     }
     ```
   - Update package.json name and description

3. **Install dependencies**
   ```bash
   # Install Go dependencies
   go mod tidy
   
   # Install Capacitor dependencies
   cd capacitor
   npm install
   ```

## 📁 Project Structure

```
.
├── .github/workflows/
│   ├── ios-match-init.yml  # Initialize iOS certificates using Fastlane Match
│   └── ios.yml            # Build and deploy to TestFlight
├── capacitor/             # Capacitor app wrapper
│   ├── ios/              # iOS specific code
│   ├── src/              # Web assets
│   └── scripts/          # Build scripts
├── web/                  # Web server
└── main.go               # Your Go application code
```

## 🔧 Development

1. **Write your Go code**
   - Modify `main.go` with your application logic
   - Use the `syscall/js` package to interact with JavaScript
   - Example:
     ```go
     package main

     import (
         "syscall/js"
     )

     func main() {
         c := make(chan struct{}, 0)
         js.Global().Set("myFunction", js.FuncOf(myFunction))
         <-c
     }
     ```

2. **Build for development**
   ```bash
   # Build WASM
   ./build.sh
   
   # Start development server
   cd web
   python -m http.server 8080
   ```

## 📱 Mobile Development

### iOS
Prerequisites:
- Xcode
- Ruby (for Fastlane)
- CocoaPods

```bash
cd capacitor
npm run ios:prepare  # Build and prepare iOS assets
npm run ios:build   # Open in Xcode
```

### Android
Prerequisites:
- Android Studio
- JDK 17+

```bash
cd capacitor
npm run android:prepare  # Build and prepare Android assets
npm run android:build   # Open in Android Studio
```

## 🚢 Deployment

### Web Deployment
1. Build the WASM binary:
   ```bash
   GOOS=js GOARCH=wasm go build -o web/main.wasm
   ```

2. Deploy using Docker:
   ```bash
   docker-compose up -d
   ```
   Or deploy to Fly.io:
   ```bash
   fly deploy
   ```

### iOS Deployment
1. Set up certificates using Match:
   - Create a private GitHub repo for certificates
   - Configure secrets in GitHub repository settings
   - Run the iOS Match Init workflow

2. Deploy to TestFlight:
   - Push to main branch or manually trigger the iOS deployment workflow
   - Configure the following secrets:
     ```
     MATCH_GIT_URL
     GIT_AUTHORIZATION
     APPLE_KEY_ID
     APPLE_ISSUER_ID
     APPLE_KEY_CONTENT
     APP_STORE_CONNECT_TEAM_ID
     DEVELOPER_APP_ID
     DEVELOPER_APP_IDENTIFIER
     DEVELOPER_PORTAL_TEAM_ID
     MATCH_PASSWORD
     TEMP_KEYCHAIN_PASSWORD
     ```

## 🔒 Security

- Never commit sensitive keys or certificates
- Use GitHub Secrets for all sensitive values
- Follow Apple's code signing best practices
- Keep dependencies updated

## 🛠 Build Options

### WASM Build Flags
- Debug build: `GOOS=js GOARCH=wasm go build`
- Release build: `GOOS=js GOARCH=wasm go build -ldflags="-w -s"`

### Environment Variables
- `BUILD_TYPE`: 'debug' or 'release'
- `VERSION_NUMBER`: Semantic version (e.g., '1.0.0')

## 📚 Resources

- [Go WebAssembly Guide](https://github.com/golang/go/wiki/WebAssembly)
- [Capacitor Documentation](https://capacitorjs.com/docs)
- [Fastlane Documentation](https://docs.fastlane.tools)

## 📄 License

MIT License - see [LICENSE](LICENSE) for details

### Attribution
When using GoAF in your project, please include appropriate attribution as specified in the [NOTICE](NOTICE) file. We kindly request that you credit Micah Longmire as the original creator of this framework in your project's documentation.
