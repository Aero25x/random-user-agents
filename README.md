[![Join our Telegram](https://img.shields.io/badge/Telegram-2CA5E0?style=for-the-badge&logo=telegram&logoColor=white)](https://t.me/hidden_coding)
[![GitHub](https://img.shields.io/badge/GitHub-181717?style=for-the-badge&logo=github&logoColor=white)](https://github.com/aero25x)
[![Twitter](https://img.shields.io/badge/Twitter-1DA1F2?style=for-the-badge&logo=x&logoColor=white)](https://x.com/aero25x)
[![YouTube](https://img.shields.io/badge/YouTube-FF0000?style=for-the-badge&logo=youtube&logoColor=white)](https://www.youtube.com/@flaming_chameleon)
[![Reddit](https://img.shields.io/badge/Reddit-FF3A00?style=for-the-badge&logo=reddit&logoColor=white)](https://www.reddit.com/r/HiddenCode/)


### README.md (English)

# Random User Agents

This project provides functions in Python, Rust, JavaScript, and TypeScript to generate random user-agent strings for different devices and browsers. These functions can be used for web scraping, automated testing, or any other scenario where you need to simulate different user agents.


[Generate user agents you could here](https://multitools.ovh/random-user-agent-generator/)


## Features

- Generate random user-agent strings for Android, iOS, Windows, and Ubuntu devices.
- Supports Chrome and Firefox browsers.
- Randomly selects versions and devices to provide realistic user-agent strings.

## Installation

To use this project, simply clone the repository and import the function into your script.

```bash
git clone https://github.com/Aero25x/random-user-agents.git
```

## Usage

### Python

```python
from random_user_agents import generate_random_user_agent

# Generate a random user-agent for an Android device using Chrome
print(generate_random_user_agent(device_type='android', browser_type='chrome'))

# Generate a random user-agent for an iOS device using Firefox
print(generate_random_user_agent(device_type='ios', browser_type='firefox'))
```

### Rust

```rust
use rand::Rng;
use rand::seq::SliceRandom;

fn generate_random_user_agent(device_type: Option<&str>, browser_type: Option<&str>) -> String {
    // Function implementation here
}

fn main() {
    for _ in 0..5 {
        println!("{}", generate_random_user_agent(None, None));
    }
}
```

### JavaScript

[Similar JS code is used here](https://multitools.ovh/random-user-agent-generator/)

```javascript
function getRandomElement(arr) {
    return arr[Math.floor(Math.random() * arr.length)];
}

function generateRandomUserAgent(deviceType, browserType) {
    // Function implementation here
}

// Example usage
for (let i = 0; i < 5; i++) {
    console.log(generateRandomUserAgent());
}
```

### TypeScript

```typescript
function getRandomElement<T>(arr: T[]): T {
    return arr[Math.floor(Math.random() * arr.length)];
}

function generateRandomUserAgent(deviceType?: string, browserType?: string): string {
    // Function implementation here
}

// Example usage
for (let i = 0; i < 5; i++) {
    console.log(generateRandomUserAgent());
}
```

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue.

## License

This project is licensed under the MIT License.

### README.md (Russian)

# Случайные Пользовательские Агенты

Этот проект предоставляет функции на Python, Rust, JavaScript и TypeScript для генерации случайных строк user-agent для различных устройств и браузеров. Эти функции могут использоваться для веб-скрапинга, автоматического тестирования или в любых других сценариях, где необходимо имитировать различные пользовательские агенты.

## Возможности

- Генерация случайных строк user-agent для устройств Android, iOS, Windows и Ubuntu.
- Поддержка браузеров Chrome и Firefox.
- Случайный выбор версий и устройств для создания реалистичных строк user-agent.

## Установка

Чтобы использовать этот проект, просто клонируйте репозиторий и импортируйте функцию в свой скрипт.

```bash
git clone https://github.com/Aero25x/random-user-agents.git
```

## Использование

### Python

```python
from random_user_agents import generate_random_user_agent

# Генерация случайного user-agent для устройства Android с использованием Chrome
print(generate_random_user_agent(device_type='android', browser_type='chrome'))

# Генерация случайного user-agent для устройства iOS с использованием Firefox
print(generate_random_user_agent(device_type='ios', browser_type='firefox'))
```

### Rust

```rust
use rand::Rng;
use rand::seq::SliceRandom;

fn generate_random_user_agent(device_type: Option<&str>, browser_type: Option<&str>) -> String {
    // Function implementation here
}

fn main() {
    for _ in 0..5 {
        println!("{}", generate_random_user_agent(None, None));
    }
}
```

### JavaScript

```javascript
function getRandomElement(arr) {
    return arr[Math.floor(Math.random() * arr.length)];
}

function generateRandomUserAgent(deviceType, browserType) {
    // Function implementation here
}

// Пример использования
for (let i = 0; i < 5; i++) {
    console.log(generateRandomUserAgent());
}
```

### TypeScript

```typescript
function getRandomElement<T>(arr: T[]): T {
    return arr[Math.floor(Math.random() * arr.length)];
}

function generateRandomUserAgent(deviceType?: string, browserType?: string): string {
    // Function implementation here
}

// Пример использования
for (let i = 0; i < 5; i++) {
    console.log(generateRandomUserAgent());
}
```

## Вклад

Приветствуются любые вклады! Пожалуйста, не стесняйтесь отправлять запросы на добавление функций или открывать вопросы.

### Контакты

[![Join our Telegram](https://img.shields.io/badge/Telegram-2CA5E0?style=for-the-badge&logo=telegram&logoColor=white)](https://t.me/hidden_coding)
[![GitHub](https://img.shields.io/badge/GitHub-181717?style=for-the-badge&logo=github&logoColor=white)](https://github.com/aero25x)
[![Twitter](https://img.shields.io/badge/Twitter-1DA1F2?style=for-the-badge&logo=x&logoColor=white)](https://x.com/aero25x)
[![YouTube](https://img.shields.io/badge/YouTube-FF0000?style=for-the-badge&logo=youtube&logoColor=white)](https://www.youtube.com/@flaming_chameleon)

Для поддержки или вопросов, свяжитесь со мной в Telegram: [@hidden_coding](https://t.me/hidden_coding)


## Лицензия

Этот проект лицензирован по лицензии MIT.
