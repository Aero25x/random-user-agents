package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Created by Aero25x
// https://github.com/Aero25x/random-user-agents.git

// GenerateRandomUserAgent generates a random user agent string
func GenerateRandomUserAgent(deviceType, browserType *string, chromeVersions, firefoxVersions []int) string {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())
	
	// Default device types and browser types
	deviceTypes := []string{"android", "ios", "windows", "ubuntu"}
	browserTypes := []string{"chrome", "firefox"}
	
	// Set default device type if not provided
	var selectedDeviceType string
	if deviceType == nil {
		selectedDeviceType = deviceTypes[rand.Intn(len(deviceTypes))]
	} else {
		selectedDeviceType = *deviceType
	}
	
	// Set default browser type if not provided
	var selectedBrowserType string
	if browserType == nil {
		selectedBrowserType = browserTypes[rand.Intn(len(browserTypes))]
	} else {
		selectedBrowserType = *browserType
	}
	
	// Set default version ranges if not provided
	if chromeVersions == nil {
		chromeVersions = []int{125, 138}
	}
	if firefoxVersions == nil {
		firefoxVersions = []int{120, 135}
	}
	
	var browserVersion string
	
	// Generate browser version
	if selectedBrowserType == "chrome" {
		majorVersion := rand.Intn(chromeVersions[1]-chromeVersions[0]) + chromeVersions[0]
		minorVersion := rand.Intn(10)
		buildVersion := rand.Intn(8999) + 1000
		patchVersion := rand.Intn(100)
		browserVersion = fmt.Sprintf("%d.%d.%d.%d", majorVersion, minorVersion, buildVersion, patchVersion)
	} else if selectedBrowserType == "firefox" {
		version := rand.Intn(firefoxVersions[1]-firefoxVersions[0]) + firefoxVersions[0]
		browserVersion = strconv.Itoa(version)
	}
	
	// Generate user agent based on device type
	switch selectedDeviceType {
	case "android":
		return generateAndroidUserAgent(selectedBrowserType, browserVersion)
	case "ios":
		return generateIOSUserAgent(selectedBrowserType, browserVersion)
	case "windows":
		return generateWindowsUserAgent(selectedBrowserType, browserVersion)
	case "ubuntu":
		return generateUbuntuUserAgent(selectedBrowserType, browserVersion)
	default:
		return ""
	}
}

func generateAndroidUserAgent(browserType, browserVersion string) string {
	androidVersions := []string{"10.0", "11.0", "12.0", "13.0", "14.0", "15.0", "16.0"}
	androidDevices := []string{
		"SM-G960F", "Pixel 5", "SM-A505F", "Pixel 4a", "Pixel 6 Pro", "SM-N975F",
		"SM-G973F", "Pixel 3", "SM-G980F", "Pixel 5a", "SM-G998B", "Pixel 4",
		"SM-G991B", "SM-G996B", "SM-F711B", "SM-F916B", "SM-G781B", "SM-N986B",
		"SM-N981B", "Pixel 2", "Pixel 2 XL", "Pixel 3 XL", "Pixel 4 XL",
		"Pixel 5 XL", "Pixel 6", "Pixel 6 XL", "Pixel 6a", "Pixel 7", "Pixel 7 Pro",
		"OnePlus 8", "OnePlus 8 Pro", "OnePlus 9", "OnePlus 9 Pro", "OnePlus Nord",
		"OnePlus Nord 2", "OnePlus Nord CE", "OnePlus 10", "OnePlus 10 Pro", "OnePlus 10T", "OnePlus 10T Pro",
		"Xiaomi Mi 9", "Xiaomi Mi 10", "Xiaomi Mi 11", "Xiaomi Redmi Note 8", "Xiaomi Redmi Note 9",
		"Huawei P30", "Huawei P40", "Huawei Mate 30", "Huawei Mate 40", "Sony Xperia 1",
		"Sony Xperia 5", "LG G8", "LG V50", "LG V60", "Nokia 8.3", "Nokia 9 PureView",
	}
	
	androidVersion := androidVersions[rand.Intn(len(androidVersions))]
	androidDevice := androidDevices[rand.Intn(len(androidDevices))]
	
	if browserType == "chrome" {
		return fmt.Sprintf("Mozilla/5.0 (Linux; Android %s; %s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Mobile Safari/537.36",
			androidVersion, androidDevice, browserVersion)
	} else if browserType == "firefox" {
		return fmt.Sprintf("Mozilla/5.0 (Android %s; Mobile; rv:%s.0) Gecko/%s.0 Firefox/%s.0",
			androidVersion, browserVersion, browserVersion, browserVersion)
	}
	return ""
}

func generateIOSUserAgent(browserType, browserVersion string) string {
	iosVersions := []string{"13.0", "14.0", "15.0", "16.0"}
	iosDevices := []string{"iPhone X", "iPhone 11", "iPhone 12", "iPhone 13", "iPad Pro", "iPad Mini"}
	
	iosVersion := iosVersions[rand.Intn(len(iosVersions))]
	iosDevice := iosDevices[rand.Intn(len(iosDevices))]
	iosVersionFormatted := strings.ReplaceAll(iosVersion, ".", "_")
	
	if browserType == "chrome" {
		return fmt.Sprintf("Mozilla/5.0 (%s; CPU iPhone OS %s like Mac OS X) AppleWebKit/537.36 (KHTML, like Gecko) CriOS/%s Mobile/15E148 Safari/604.1",
			iosDevice, iosVersionFormatted, browserVersion)
	} else if browserType == "firefox" {
		return fmt.Sprintf("Mozilla/5.0 (%s; CPU iPhone OS %s like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/%s.0 Mobile/15E148 Safari/605.1.15",
			iosDevice, iosVersionFormatted, browserVersion)
	}
	return ""
}

func generateWindowsUserAgent(browserType, browserVersion string) string {
	windowsVersions := []string{"10.0", "11.0"}
	windowsVersion := windowsVersions[rand.Intn(len(windowsVersions))]
	
	if browserType == "chrome" {
		return fmt.Sprintf("Mozilla/5.0 (Windows NT %s; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36",
			windowsVersion, browserVersion)
	} else if browserType == "firefox" {
		return fmt.Sprintf("Mozilla/5.0 (Windows NT %s; Win64; x64; rv:%s.0) Gecko/%s.0 Firefox/%s.0",
			windowsVersion, browserVersion, browserVersion, browserVersion)
	}
	return ""
}

func generateUbuntuUserAgent(browserType, browserVersion string) string {
	if browserType == "chrome" {
		return fmt.Sprintf("Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:94.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36",
			browserVersion)
	} else if browserType == "firefox" {
		return fmt.Sprintf("Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:%s.0) Gecko/%s.0 Firefox/%s.0",
			browserVersion, browserVersion)
	}
	return ""
}

// Example usage
func main() {
	for i := 0; i < 5; i++ {
		userAgent := GenerateRandomUserAgent(nil, nil, nil, nil)
		fmt.Println(userAgent)
	}
}
