use rand::Rng;
use rand::seq::SliceRandom;

fn generate_random_user_agent(device_type: Option<&str>, browser_type: Option<&str>) -> String {
    let device_type = device_type.unwrap_or_else(|| {
        let choices = vec!["android", "ios", "windows", "ubuntu"];
        choices.choose(&mut rand::thread_rng()).unwrap()
    });

    let browser_type = browser_type.unwrap_or_else(|| {
        let choices = vec!["chrome", "firefox"];
        choices.choose(&mut rand::thread_rng()).unwrap()
    });

    let browser_version = if browser_type == "chrome" {
        let major_version = rand::thread_rng().gen_range(110..127);
        let minor_version = rand::thread_rng().gen_range(0..10);
        let build_version = rand::thread_rng().gen_range(1000..10000);
        let patch_version = rand::thread_rng().gen_range(0..100);
        format!("{}.{}.{}.{}", major_version, minor_version, build_version, patch_version)
    } else {
        let firefox_versions = (90..100).collect::<Vec<i32>>();
        firefox_versions.choose(&mut rand::thread_rng()).unwrap().to_string()
    };

    match device_type {
        "android" => {
            let android_versions = vec!["10.0", "11.0", "12.0", "13.0"];
            let android_devices = vec![
                "SM-G960F", "Pixel 5", "SM-A505F", "Pixel 4a", "Pixel 6 Pro", "SM-N975F",
                "SM-G973F", "Pixel 3", "SM-G980F", "Pixel 5a", "SM-G998B", "Pixel 4",
                "SM-G991B", "SM-G996B", "SM-F711B", "SM-F916B", "SM-G781B", "SM-N986B",
                "SM-N981B", "Pixel 2", "Pixel 2 XL", "Pixel 3 XL", "Pixel 4 XL",
                "Pixel 5 XL", "Pixel 6", "Pixel 6 XL", "Pixel 6a", "Pixel 7", "Pixel 7 Pro",
                "OnePlus 8", "OnePlus 8 Pro", "OnePlus 9", "OnePlus 9 Pro", "OnePlus Nord", "OnePlus Nord 2", "OnePlus Nord CE", "OnePlus 10", "OnePlus 10 Pro", "OnePlus 10T", "OnePlus 10T Pro",
                "Xiaomi Mi 9", "Xiaomi Mi 10", "Xiaomi Mi 11", "Xiaomi Redmi Note 8", "Xiaomi Redmi Note 9",
                "Huawei P30", "Huawei P40", "Huawei Mate 30", "Huawei Mate 40", "Sony Xperia 1",
                "Sony Xperia 5", "LG G8", "LG V50", "LG V60", "Nokia 8.3", "Nokia 9 PureView"
            ];
            let android_version = android_versions.choose(&mut rand::thread_rng()).unwrap();
            let android_device = android_devices.choose(&mut rand::thread_rng()).unwrap();
            if browser_type == "chrome" {
                format!(
                    "Mozilla/5.0 (Linux; Android {}; {}) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/{} Mobile Safari/537.36",
                    android_version, android_device, browser_version
                )
            } else {
                format!(
                    "Mozilla/5.0 (Android {}; Mobile; rv:{}.) Gecko/{} Firefox/{}.0",
                    android_version, browser_version, browser_version, browser_version
                )
            }
        }
        "ios" => {
            let ios_versions = vec!["13.0", "14.0", "15.0", "16.0"];
            let ios_devices = vec![
                "iPhone X", "iPhone 11", "iPhone 12", "iPhone 13", "iPad Pro", "iPad Mini"
            ];
            let ios_version = ios_versions.choose(&mut rand::thread_rng()).unwrap();
            let ios_device = ios_devices.choose(&mut rand::thread_rng()).unwrap();
            if browser_type == "chrome" {
                format!(
                    "Mozilla/5.0 (iPhone; CPU iPhone OS {} like Mac OS X) AppleWebKit/537.36 (KHTML, like Gecko) CriOS/{} Mobile/15E148 Safari/604.1",
                    ios_version.replace(".", "_"), browser_version
                )
            } else {
                format!(
                    "Mozilla/5.0 (iPhone; CPU iPhone OS {} like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/{}.0 Mobile/15E148 Safari/605.1.15",
                    ios_version.replace(".", "_"), browser_version
                )
            }
        }
        "windows" => {
            let windows_versions = vec!["10.0", "11.0"];
            let windows_version = windows_versions.choose(&mut rand::thread_rng()).unwrap();
            if browser_type == "chrome" {
                format!(
                    "Mozilla/5.0 (Windows NT {}; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/{} Safari/537.36",
                    windows_version, browser_version
                )
            } else {
                format!(
                    "Mozilla/5.0 (Windows NT {}; Win64; x64; rv:{}.) Gecko/{} Firefox/{}.0",
                    windows_version, browser_version, browser_version, browser_version
                )
            }
        }
        "ubuntu" => {
            let ubuntu_versions = vec!["20.04", "22.04"];
            let ubuntu_version = ubuntu_versions.choose(&mut rand::thread_rng()).unwrap();
            if browser_type == "chrome" {
                format!(
                    "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:94.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/{} Safari/537.36",
                    browser_version
                )
            } else {
                format!(
                    "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:{}.) Gecko/{} Firefox/{}.0",
                    browser_version, browser_version, browser_version
                )
            }
        }
        _ => String::new()
    }
}

fn main() {
    for _ in 0..5 {
        println!("{}", generate_random_user_agent(None, None));
    }
}
