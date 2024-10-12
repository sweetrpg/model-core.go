// swift-tools-version: 6.0
// The swift-tools-version declares the minimum version of Swift required to build this package.

import PackageDescription

let package = Package(
    name: "ModelCore",
    platforms: [
       .macOS(.v13),
    ],
    products: [
        .library(
            name: "ModelCore",
            targets: ["ModelCore"]),
    ],
    dependencies: [
        .package(url: "https://github.com/apple/swift-docc-plugin", from: "1.0.0"),
    ],
    targets: [
        .target(
            name: "ModelCore",
            dependencies: [],
            swiftSettings: swiftSettings),
        .testTarget(
            name: "ModelCoreTests",
            dependencies: ["ModelCore"],
            swiftSettings: swiftSettings),
    ],
    swiftLanguageModes: [.v5]
)

var swiftSettings: [SwiftSetting] { [
    .enableUpcomingFeature("DisableOutwardActorInference"),
    .enableExperimentalFeature("StrictConcurrency"),
] }
