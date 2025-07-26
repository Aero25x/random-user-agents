# Created by Aero25x
# https://github.com/Aero25x/random-user-agents.git

def generate_random_user_agent(device_type: nil, browser_type: nil, chrome_versions: [125, 138], firefox_versions: [120, 135])
  device_type ||= ['android', 'ios', 'windows', 'ubuntu'].sample
  browser_type ||= ['chrome', 'firefox'].sample
  
  if browser_type == 'chrome'
    chrome_version_range = (chrome_versions[0]...chrome_versions[1]).to_a
    major_version = chrome_version_range.sample
    minor_version = rand(0..9)
    build_version = rand(1000..9999)
    patch_version = rand(0..99)
    browser_version = "#{major_version}.#{minor_version}.#{build_version}.#{patch_version}"
  elsif browser_type == 'firefox'
    firefox_version_range = (firefox_versions[0]...firefox_versions[1]).to_a
    browser_version = firefox_version_range.sample
  end
  
  case device_type
  when 'android'
    android_versions = ['10.0', '11.0', '12.0', '13.0', '14.0', '15.0', '16.0']
    android_devices = [
      'SM-G960F', 'Pixel 5', 'SM-A505F', 'Pixel 4a', 'Pixel 6 Pro', 'SM-N975F',
      'SM-G973F', 'Pixel 3', 'SM-G980F', 'Pixel 5a', 'SM-G998B', 'Pixel 4',
      'SM-G991B', 'SM-G996B', 'SM-F711B', 'SM-F916B', 'SM-G781B', 'SM-N986B',
      'SM-N981B', 'Pixel 2', 'Pixel 2 XL', 'Pixel 3 XL', 'Pixel 4 XL',
      'Pixel 5 XL', 'Pixel 6', 'Pixel 6 XL', 'Pixel 6a', 'Pixel 7', 'Pixel 7 Pro',
      'OnePlus 8', 'OnePlus 8 Pro', 'OnePlus 9', 'OnePlus 9 Pro', 'OnePlus Nord', 
      'OnePlus Nord 2', 'OnePlus Nord CE', 'OnePlus 10', 'OnePlus 10 Pro', 
      'OnePlus 10T', 'OnePlus 10T Pro', 'Xiaomi Mi 9', 'Xiaomi Mi 10', 
      'Xiaomi Mi 11', 'Xiaomi Redmi Note 8', 'Xiaomi Redmi Note 9',
      'Huawei P30', 'Huawei P40', 'Huawei Mate 30', 'Huawei Mate 40', 
      'Sony Xperia 1', 'Sony Xperia 5', 'LG G8', 'LG V50', 'LG V60', 
      'Nokia 8.3', 'Nokia 9 PureView'
    ]
    
    android_device = android_devices.sample
    android_version = android_versions.sample
    
    if browser_type == 'chrome'
      "Mozilla/5.0 (Linux; Android #{android_version}; #{android_device}) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/#{browser_version} Mobile Safari/537.36"
    elsif browser_type == 'firefox'
      "Mozilla/5.0 (Android #{android_version}; Mobile; rv:#{browser_version}.0) Gecko/#{browser_version}.0 Firefox/#{browser_version}.0"
    end
    
  when 'ios'
    ios_versions = ['13.0', '14.0', '15.0', '16.0']
    ios_devices = ['iPhone X', 'iPhone 11', 'iPhone 12', 'iPhone 13','iPhone 15', 'iPhone 16', 'iPad Pro', 'iPad Mini']
    
    ios_device = ios_devices.sample
    ios_version = ios_versions.sample
    
    if browser_type == 'chrome'
      "Mozilla/5.0 (iPhone; CPU iPhone OS #{ios_version.gsub('.', '_')} like Mac OS X) AppleWebKit/537.36 (KHTML, like Gecko) CriOS/#{browser_version} Mobile/15E148 Safari/604.1"
    elsif browser_type == 'firefox'
      "Mozilla/5.0 (iPhone; CPU iPhone OS #{ios_version.gsub('.', '_')} like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/#{browser_version}.0 Mobile/15E148 Safari/605.1.15"
    end
    
  when 'windows'
    windows_versions = ['10.0', '11.0']
    windows_version = windows_versions.sample
    
    if browser_type == 'chrome'
      "Mozilla/5.0 (Windows NT #{windows_version}; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/#{browser_version} Safari/537.36"
    elsif browser_type == 'firefox'
      "Mozilla/5.0 (Windows NT #{windows_version}; Win64; x64; rv:#{browser_version}.0) Gecko/#{browser_version}.0 Firefox/#{browser_version}.0"
    end
    
  when 'ubuntu'
    ubuntu_versions = ['20.04', '22.04', '24.04']
    ubuntu_version = ubuntu_versions.sample
    
    if browser_type == 'chrome'
      "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:94.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/#{browser_version} Safari/537.36"
    elsif browser_type == 'firefox'
      "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:#{browser_version}.0) Gecko/#{browser_version}.0 Firefox/#{browser_version}.0"
    end
  end
end

# Example usage
if __FILE__ == $0
  5.times do
    puts generate_random_user_agent
  end
end
