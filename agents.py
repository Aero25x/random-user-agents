import random

# Created by Aero25x
# https://github.com/Aero25x/random-user-agents.git

def generate_random_user_agent(device_type=None, browser_type=None, chrome_versions=[125, 138], firefox_versions=[120, 135]):
    if not device_type:
        device_type = random.choice(['android', 'ios', 'windows', 'ubuntu'])

    if not browser_type:
        browser_type = random.choice(['chrome', 'firefox'])

    if browser_type == 'chrome':
        chrome_versions = list(range(chrome_versions[0], chrome_versions[1]))
        major_version = random.choice(chrome_versions)
        minor_version = random.randint(0, 9)
        build_version = random.randint(1000, 9999)
        patch_version = random.randint(0, 99)
        browser_version = f"{major_version}.{minor_version}.{build_version}.{patch_version}"
    elif browser_type == 'firefox':
        firefox_versions = list(range(firefox_versions[0], firefox_versions[1]))  # Last 10 versions of Firefox
        browser_version = random.choice(firefox_versions)

    if device_type == 'android':
        android_versions = ['10.0', '11.0', '12.0', '13.0', '14.0', '15.0', '16.0']
        android_device = random.choice([
            'SM-G960F', 'Pixel 5', 'SM-A505F', 'Pixel 4a', 'Pixel 6 Pro', 'SM-N975F',
            'SM-G973F', 'Pixel 3', 'SM-G980F', 'Pixel 5a', 'SM-G998B', 'Pixel 4',
            'SM-G991B', 'SM-G996B', 'SM-F711B', 'SM-F916B', 'SM-G781B', 'SM-N986B',
            'SM-N981B', 'Pixel 2', 'Pixel 2 XL', 'Pixel 3 XL', 'Pixel 4 XL',
            'Pixel 5 XL', 'Pixel 6', 'Pixel 6 XL', 'Pixel 6a', 'Pixel 7', 'Pixel 7 Pro',
            'OnePlus 8', 'OnePlus 8 Pro', 'OnePlus 9', 'OnePlus 9 Pro', 'OnePlus Nord', 'OnePlus Nord 2', 'OnePlus Nord CE', 'OnePlus 10', 'OnePlus 10 Pro', 'OnePlus 10T', 'OnePlus 10T Pro',
            'Xiaomi Mi 9', 'Xiaomi Mi 10', 'Xiaomi Mi 11', 'Xiaomi Redmi Note 8', 'Xiaomi Redmi Note 9',
            'Huawei P30', 'Huawei P40', 'Huawei Mate 30', 'Huawei Mate 40', 'Sony Xperia 1',
            'Sony Xperia 5', 'LG G8', 'LG V50', 'LG V60', 'Nokia 8.3', 'Nokia 9 PureView'
        ])
        android_version = random.choice(android_versions)
        if browser_type == 'chrome':
            return f"Mozilla/5.0 (Linux; Android {android_version}; {android_device}) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/{browser_version} Mobile Safari/537.36"
        elif browser_type == 'firefox':
            return f"Mozilla/5.0 (Android {android_version}; Mobile; rv:{browser_version}.0) Gecko/{browser_version}.0 Firefox/{browser_version}.0"

    elif device_type == 'ios':
        ios_versions = ['13.0', '14.0', '15.0', '16.0']
        ios_device = random.choice([
            'iPhone X', 'iPhone 11', 'iPhone 12', 'iPhone 13', 'iPad Pro', 'iPad Mini'
        ])
        ios_version = random.choice(ios_versions)
        if browser_type == 'chrome':
            return f"Mozilla/5.0 (iPhone; CPU iPhone OS {ios_version.replace('.', '_')} like Mac OS X) AppleWebKit/537.36 (KHTML, like Gecko) CriOS/{browser_version} Mobile/15E148 Safari/604.1"
        elif browser_type == 'firefox':
            return f"Mozilla/5.0 (iPhone; CPU iPhone OS {ios_version.replace('.', '_')} like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/{browser_version}.0 Mobile/15E148 Safari/605.1.15"

    elif device_type == 'windows':
        windows_versions = ['10.0', '11.0']
        windows_version = random.choice(windows_versions)
        if browser_type == 'chrome':
            return f"Mozilla/5.0 (Windows NT {windows_version}; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/{browser_version} Safari/537.36"
        elif browser_type == 'firefox':
            return f"Mozilla/5.0 (Windows NT {windows_version}; Win64; x64; rv:{browser_version}.0) Gecko/{browser_version}.0 Firefox/{browser_version}.0"

    elif device_type == 'ubuntu':
        ubuntu_versions = ['20.04', '22.04']
        ubuntu_version = random.choice(ubuntu_versions)
        if browser_type == 'chrome':
            return f"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:94.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/{browser_version} Safari/537.36"
        elif browser_type == 'firefox':
            return f"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:{browser_version}.0) Gecko/{browser_version}.0 Firefox/{browser_version}.0"

    return None


# Example usage
if __name__ == "__main__":
    for _ in range(5):
        print(generate_random_user_agent())

