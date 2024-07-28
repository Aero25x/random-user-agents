function getRandomElement<T>(arr: T[]): T {
    return arr[Math.floor(Math.random() * arr.length)];
}

function generateRandomUserAgent(deviceType?: string, browserType?: string): string {
    const devices = ['android', 'ios', 'windows', 'ubuntu'];
    const browsers = ['chrome', 'firefox'];

    if (!deviceType) {
        deviceType = getRandomElement(devices);
    }

    if (!browserType) {
        browserType = getRandomElement(browsers);
    }

    let browserVersion: string;
    if (browserType === 'chrome') {
        const majorVersion = Math.floor(Math.random() * (127 - 110) + 110);
        const minorVersion = Math.floor(Math.random() * 10);
        const buildVersion = Math.floor(Math.random() * (10000 - 1000) + 1000);
        const patchVersion = Math.floor(Math.random() * 100);
        browserVersion = `${majorVersion}.${minorVersion}.${buildVersion}.${patchVersion}`;
    } else {
        const firefoxVersions = Array.from({ length: 10 }, (_, i) => 90 + i);
        browserVersion = getRandomElement(firefoxVersions).toString();
    }

    if (deviceType === 'android') {
        const androidVersions = ['10.0', '11.0', '12.0', '13.0'];
        const androidDevices = [
            'SM-G960F', 'Pixel 5', 'SM-A505F', 'Pixel 4a', 'Pixel 6 Pro', 'SM-N975F',
            'SM-G973F', 'Pixel 3', 'SM-G980F', 'Pixel 5a', 'SM-G998B', 'Pixel 4',
            'SM-G991B', 'SM-G996B', 'SM-F711B', 'SM-F916B', 'SM-G781B', 'SM-N986B',
            'SM-N981B', 'Pixel 2', 'Pixel 2 XL', 'Pixel 3 XL', 'Pixel 4 XL',
            'Pixel 5 XL', 'Pixel 6', 'Pixel 6 XL', 'Pixel 6a', 'Pixel 7', 'Pixel 7 Pro',
            'OnePlus 8', 'OnePlus 8 Pro', 'OnePlus 9', 'OnePlus 9 Pro', 'OnePlus Nord', 'OnePlus Nord 2', 'OnePlus Nord CE', 'OnePlus 10', 'OnePlus 10 Pro', 'OnePlus 10T', 'OnePlus 10T Pro',
            'Xiaomi Mi 9', 'Xiaomi Mi 10', 'Xiaomi Mi 11', 'Xiaomi Redmi Note 8', 'Xiaomi Redmi Note 9',
            'Huawei P30', 'Huawei P40', 'Huawei Mate 30', 'Huawei Mate 40', 'Sony Xperia 1',
            'Sony Xperia 5', 'LG G8', 'LG V50', 'LG V60', 'Nokia 8.3', 'Nokia 9 PureView'
        ];
        const androidVersion = getRandomElement(androidVersions);
        const androidDevice = getRandomElement(androidDevices);
        if (browserType === 'chrome') {
            return `Mozilla/5.0 (Linux; Android ${androidVersion}; ${androidDevice}) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/${browserVersion} Mobile Safari/537.36`;
        } else {
            return `Mozilla/5.0 (Android ${androidVersion}; Mobile; rv:${browserVersion}.0) Gecko/${browserVersion}.0 Firefox/${browserVersion}.0`;
        }
    } else if (deviceType === 'ios') {
        const iosVersions = ['13.0', '14.0', '15.0', '16.0'];
        const iosDevices = [
            'iPhone X', 'iPhone 11', 'iPhone 12', 'iPhone 13', 'iPad Pro', 'iPad Mini'
        ];
        const iosVersion = getRandomElement(iosVersions);
        const iosDevice = getRandomElement(iosDevices);
        if (browserType === 'chrome') {
            return `Mozilla/5.0 (iPhone; CPU iPhone OS ${iosVersion.replace('.', '_')} like Mac OS X) AppleWebKit/537.36 (KHTML, like Gecko) CriOS/${browserVersion} Mobile/15E148 Safari/604.1`;
        } else {
            return `Mozilla/5.0 (iPhone; CPU iPhone OS ${iosVersion.replace('.', '_')} like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/${browserVersion}.0 Mobile/15E148 Safari/605.1.15`;
        }
    } else if (deviceType === 'windows') {
        const windowsVersions = ['10.0', '11.0'];
        const windowsVersion = getRandomElement(windowsVersions);
        if (browserType === 'chrome') {
            return `Mozilla/5.0 (Windows NT ${windowsVersion}; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/${browserVersion} Safari/537.36`;
        } else {
            return `Mozilla/5.0 (Windows NT ${windowsVersion}; Win64; x64; rv:${browserVersion}.0) Gecko/${browserVersion}.0 Firefox/${browserVersion}.0`;
        }
    } else if (deviceType === 'ubuntu') {
        const ubuntuVersions = ['20.04', '22.04'];
        const ubuntuVersion = getRandomElement(ubuntuVersions);
        if (browserType === 'chrome') {
            return `Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:94.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/${browserVersion} Safari/537.36`;
        } else {
            return `Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:${browserVersion}.0) Gecko/${browserVersion}.0 Firefox/${browserVersion}.0`;
        }
    }

    return '';
}

// Example usage
for (let i = 0; i < 5; i++) {
    console.log(generateRandomUserAgent());
}
