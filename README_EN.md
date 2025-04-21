# MultiTranslatorUnifier

A powerful unified translation service supporting multiple translation sources (Google, Bing, DeepLX) with intelligent caching and failover mechanisms.

## Features

- **Multiple Translation Sources**
  - Google Translate
  - Microsoft Bing Translator
  - DeepLX API

- **Intelligent Caching System**
  - Automatic translation result caching
  - Fast historical translation retrieval
  - Reduced duplicate requests

- **High Availability Design**
  - Parallel requests to multiple translation sources
  - Automatic failover
  - Timeout auto-retry mechanism

- **Cross-Platform Support**
  - Google and Bing translation support on Linux systems
  - DeepLX API support on all platforms

## Requirements

- Go 1.24
- MySQL database (for storing translation history)
- `translate-shell` installation required for Linux environments

## Quick Start

### Installation

```bash
# Clone the repository
git clone https://github.com/zhangyiming748/MultiTranslatorUnifier.git

# Navigate to project directory
cd MultiTranslatorUnifier

# Install dependencies
go mod download
```

### Configuration

1. **Environment Variables**

```bash
# Google Translate proxy settings (optional)
export PROXY="your-proxy-address"

# DeepLX API key
export LINUXDO="your-api-key"
```

2. **Database Configuration**

Ensure MySQL service is running and database connection information is properly configured.

### Usage

```go
import "github.com/zhangyiming748/MultiTranslatorUnifier/logic"

func main() {
    source := "Hello World"
    from, result := logic.Trans(source)
    fmt.Printf("Translation source: %s\nResult: %s\n", from, result)
}
```

## Technical Architecture

### Core Components

- **Trans Module**: Unified translation entry point, responsible for scheduling and managing translation processes
- **Storage Module**: Handles translation history storage and retrieval
- **Translate-Shell Module**: Encapsulates Google and Bing translation interfaces
- **LinuxDo Module**: Integrates DeepLX API service

### Workflow

1. Receive translation request
2. Check translation cache
3. Send parallel requests to available translation sources
4. Use the fastest valid response
5. Cache new translation results

## Performance Optimization

- Parallel translation requests using goroutines
- sync.Once ensures only the fastest result is used
- Intelligent caching mechanism to avoid duplicate translations
- 30-second timeout protection with automatic retry mechanism

## Error Handling

- Automatic failover between translation sources
- Timeout auto-retry
- Detailed error logging

## Docker Support

The project includes Docker support for quick deployment:

```bash
# Build image
docker-compose build

# Start service
docker-compose up -d
```

## License

This project is open-sourced under the MIT License.

## Contributing

Issues and Pull Requests are welcome to help improve the project. Before submitting code, please ensure:

1. Code follows Go coding standards
2. Necessary test cases are added
3. All test cases pass
4. Related documentation is updated