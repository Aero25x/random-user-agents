# Random User Agent Generator - Python, JavaScript, Rust, TypeScript, Go

[![Join our Telegram](https://img.shields.io/badge/Telegram-2CA5E0?style=for-the-badge&logo=telegram&logoColor=white)](https://t.me/hidden_coding)
[![GitHub](https://img.shields.io/badge/GitHub-181717?style=for-the-badge&logo=github&logoColor=white)](https://github.com/aero25x)
[![Twitter](https://img.shields.io/badge/Twitter-1DA1F2?style=for-the-badge&logo=x&logoColor=white)](https://x.com/aero25x)
[![YouTube](https://img.shields.io/badge/YouTube-FF0000?style=for-the-badge&logo=youtube&logoColor=white)](https://www.youtube.com/@flaming_chameleon)
[![Reddit](https://img.shields.io/badge/Reddit-FF3A00?style=for-the-badge&logo=reddit&logoColor=white)](https://www.reddit.com/r/HiddenCode/)
[![GitHub Stars](https://img.shields.io/github/stars/Aero25x/random-user-agents?style=for-the-badge)](https://github.com/Aero25x/random-user-agents/stargazers)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge)](https://opensource.org/licenses/MIT)

> **Generate realistic random user agents for web scraping, browser automation, and testing. Supports Chrome, Firefox, Safari on Android, iOS, Windows, macOS, Linux, and Ubuntu devices.**

🌐 **[Try Online Generator](https://multitools.ovh/random-user-agent-generator/)** | 📚 **[Documentation](#usage)** | 🤝 **[Contributing](#contributing)**

## 🚀 Why Random User Agent Generator?

**Random User Agent Generator** is a lightweight, cross-platform library for generating authentic browser user agent strings. Perfect for:

- 🕷️ **Web Scraping** - Rotate user agents to avoid detection and blocking
- 🤖 **Browser Automation** - Selenium, Puppeteer, Playwright testing with realistic agents
- 🧪 **API Testing** - Test endpoints with different client configurations
- 🔒 **Privacy Tools** - Mask browser fingerprints and enhance anonymity
- 📊 **Analytics Testing** - Simulate traffic from various devices and browsers

### ✨ Key Features

✅ **Multi-language support** - Python, JavaScript, TypeScript, Rust, Go  
✅ **Realistic user agents** - Based on real browser version distributions  
✅ **Device targeting** - Android, iOS, Windows, macOS, Linux, Ubuntu  
✅ **Browser selection** - Chrome, Firefox, Safari, Edge support  
✅ **Zero dependencies** - Lightweight and fast  
✅ **MIT License** - Free for commercial use  
✅ **Actively maintained** - Regular updates with latest browser versions  

## 📦 Installation

### Python (pip)

```bash
git clonse https://github.com/Aero25x/random-user-agents.git
```

```python
from random_user_agents import generate_random_user_agent

# Generate random user agent
user_agent = generate_random_user_agent()
print(user_agent)

# Specific device and browser
android_chrome = generate_random_user_agent(device_type='android', browser_type='chrome')
ios_safari = generate_random_user_agent(device_type='ios', browser_type='safari')
```

### JavaScript / Node.js (npm)

```bash
git clonse https://github.com/Aero25x/random-user-agents.git
```

```javascript
const { generateRandomUserAgent } = require('random-user-agents');

// Random user agent
console.log(generateRandomUserAgent());

// Chrome on Windows
console.log(generateRandomUserAgent('windows', 'chrome'));

// Firefox on Linux
console.log(generateRandomUserAgent('linux', 'firefox'));
```

### TypeScript

```bash
git clonse https://github.com/Aero25x/random-user-agents.git
```

```typescript
import { generateRandomUserAgent } from 'random-user-agents';

const userAgent: string = generateRandomUserAgent('android', 'chrome');
console.log(userAgent);
```

### Rust (Cargo)

```bash
git clonse https://github.com/Aero25x/random-user-agents.git
```

```rust
use random_user_agents::generate_random_user_agent;

fn main() {
    let user_agent = generate_random_user_agent(None, None);
    println!("{}", user_agent);
    
    // Android Chrome
    let android_ua = generate_random_user_agent(Some("android"), Some("chrome"));
    println!("{}", android_ua);
}
```

### Go (Golang)

```bash
go get github.com/Aero25x/random-user-agents
```

```go
package main

import (
    "fmt"
    rua "github.com/Aero25x/random-user-agents"
)

func main() {
    // Random user agent
    userAgent := rua.GenerateRandomUserAgent(nil, nil, nil, nil)
    fmt.Println(userAgent)
    
    // Android Chrome with version range
    deviceType := "android"
    browserType := "chrome"
    chromeVersions := []int{120, 130}
    
    ua := rua.GenerateRandomUserAgent(&deviceType, &browserType, chromeVersions, nil)
    fmt.Println(ua)
}
```

### Manual Installation (Git Clone)

```bash
git clone https://github.com/Aero25x/random-user-agents.git
cd random-user-agents
```

## 💡 Usage Examples

### Web Scraping with Python (Requests)

```python
import requests
from random_user_agents import generate_random_user_agent

headers = {
    'User-Agent': generate_random_user_agent()
}

response = requests.get('https://example.com', headers=headers)
print(response.status_code)
```

### Selenium Browser Automation

```python
from selenium import webdriver
from random_user_agents import generate_random_user_agent

options = webdriver.ChromeOptions()
options.add_argument(f'user-agent={generate_random_user_agent("windows", "chrome")}')

driver = webdriver.Chrome(options=options)
driver.get('https://example.com')
```

### Puppeteer with JavaScript

```javascript
const puppeteer = require('puppeteer');
const { generateRandomUserAgent } = require('random-user-agents');

(async () => {
    const browser = await puppeteer.launch();
    const page = await browser.newPage();
    
    await page.setUserAgent(generateRandomUserAgent('windows', 'chrome'));
    await page.goto('https://example.com');
    
    await browser.close();
})();
```

### Axios HTTP Client

```javascript
const axios = require('axios');
const { generateRandomUserAgent } = require('random-user-agents');

axios.get('https://api.example.com', {
    headers: {
        'User-Agent': generateRandomUserAgent()
    }
}).then(response => {
    console.log(response.data);
});
```

### Fetch API (Browser/Node.js)

```javascript
fetch('https://api.example.com', {
    headers: {
        'User-Agent': generateRandomUserAgent('windows', 'firefox')
    }
})
.then(response => response.json())
.then(data => console.log(data));
```

## 🎯 Supported Configurations

### Device Types
- `android` - Android smartphones and tablets
- `ios` - iPhone and iPad devices
- `windows` - Windows desktop (7, 10, 11)
- `macos` - macOS desktop
- `linux` - Linux distributions
- `ubuntu` - Ubuntu Linux

### Browser Types
- `chrome` - Google Chrome (80-130)
- `firefox` - Mozilla Firefox (90-125)
- `safari` - Apple Safari (iOS and macOS)
- `edge` - Microsoft Edge

### Example Output

```
Mozilla/5.0 (Linux; Android 13; SM-S908B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Mobile Safari/537.36

Mozilla/5.0 (iPhone; CPU iPhone OS 17_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2 Mobile/15E148 Safari/604.1

Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36

Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:121.0) Gecko/20100101 Firefox/121.0
```

## 🔧 Advanced Usage

### Custom Version Ranges (Go)

```go
chromeVersions := []int{115, 120}  // Chrome 115 to 120
firefoxVersions := []int{100, 110} // Firefox 100 to 110

userAgent := GenerateRandomUserAgent(nil, nil, chromeVersions, firefoxVersions)
```

### Batch Generation (Python)

```python
from random_user_agents import generate_random_user_agent

# Generate 100 unique user agents
user_agents = [generate_random_user_agent() for _ in range(100)]

# Generate mobile-only user agents
mobile_agents = [
    generate_random_user_agent(device_type=device) 
    for device in ['android', 'ios'] * 50
]
```

### Rate Limiting Protection

```python
import time
import random
from random_user_agents import generate_random_user_agent

for url in urls:
    headers = {'User-Agent': generate_random_user_agent()}
    response = requests.get(url, headers=headers)
    
    # Random delay between requests
    time.sleep(random.uniform(1, 3))
```

## 🛠️ Useful Tools & Resources

| Tool | Description | Link |
|------|-------------|------|
| 🌐 **User Agent Generator** | Online generator with copy function | [Open Tool](https://multitools.ovh/random-user-agent-generator/) |
| 🧬 **Base64 Converter** | Encode/decode Base64 strings | [Open Tool](https://multitools.ovh/base64-converter/) |
| 🔍 **RegEx Validator** | Test and validate Regular Expressions | [Open Tool](https://multitools.ovh/regex-validator/) |
| 🔐 **JWT Decoder** | Decode and verify JSON Web Tokens | [Open Tool](https://multitools.ovh/jwt-converter/) |
| ⏱️ **Timestamp Converter** | Convert Unix timestamps | [Open Tool](https://multitools.ovh/timestamp/) |
| 🌎 **Time Zone Converter** | World time zone conversion | [Open Tool](https://multitools.ovh/world-time-zone/) |

## 🤝 Contributing

Contributions are welcome! Here's how you can help:

1. 🍴 Fork the repository
2. 🌿 Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. ✅ Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. 📤 Push to the branch (`git push origin feature/AmazingFeature`)
5. 🎉 Open a Pull Request

### Development Setup

```bash
# Clone the repository
git clone https://github.com/Aero25x/random-user-agents.git
cd random-user-agents

# Install dependencies
pip install -r requirements.txt  # Python
npm install                       # JavaScript/TypeScript
cargo build                       # Rust
go mod download                   # Go
```

## 📝 License

This project is licensed under the **MIT License** - see the [LICENSE](LICENSE) file for details.

Free for personal and commercial use, modification, and distribution.

## 🌟 Star History

If this project helped you, please consider giving it a ⭐️ on GitHub!

[![Star History Chart](https://api.star-history.com/svg?repos=Aero25x/random-user-agents&type=Date)](https://star-history.com/#Aero25x/random-user-agents&Date)

## 📞 Support & Contact

- 💬 **Telegram Community**: [@hidden_coding](https://t.me/hidden_coding)
- 🐛 **Bug Reports**: [GitHub Issues](https://github.com/Aero25x/random-user-agents/issues)
- 💡 **Feature Requests**: [GitHub Discussions](https://github.com/Aero25x/random-user-agents/discussions)
- 📧 **Email**: Contact via GitHub profile

## 🔗 Related Projects

- [fake-useragent](https://github.com/fake-useragent/fake-useragent) - Python user agent faker
- [user-agents](https://github.com/intoli/user-agents) - JavaScript user agent parser
- [random-useragent](https://github.com/skratchdot/random-useragent) - Node.js random user agents

## 📊 Statistics

![GitHub repo size](https://img.shields.io/github/repo-size/Aero25x/random-user-agents)
![GitHub language count](https://img.shields.io/github/languages/count/Aero25x/random-user-agents)
![GitHub top language](https://img.shields.io/github/languages/top/Aero25x/random-user-agents)
![GitHub last commit](https://img.shields.io/github/last-commit/Aero25x/random-user-agents)
![GitHub issues](https://img.shields.io/github/issues/Aero25x/random-user-agents)
![GitHub pull requests](https://img.shields.io/github/issues-pr/Aero25x/random-user-agents)

---

**Made with ❤️ by [Aero25x](https://github.com/aero25x)**

Keywords: user agent generator, random user agent, fake user agent, web scraping, browser automation, selenium user agent, puppeteer user agent, playwright, http headers, user agent string, mobile user agent, desktop user agent, chrome user agent, firefox user agent, safari user agent, android user agent, ios user agent, windows user agent, linux user agent, python user agent, javascript user agent, rust user agent, golang user agent, typescript user agent, web crawler, bot detection bypass, anti bot, scraping tool, automation testing
