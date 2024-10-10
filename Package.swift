// swift-tools-version: 5.7
// The swift-tools-version declares the minimum version of Swift required to build this package.

import PackageDescription

let package = Package(
    name: "ModelCore",
    products: [
        .library(
            name: "ModelCore",
            targets: ["ModelCore"]),
    ],
    targets: [
        .target(
            name: "ModelCore",
            dependencies: []),
        .testTarget(
            name: "ModelCoreTests",
            dependencies: ["ModelCore"]),
    ],
    dependencies: [
        .package(url: "https://github.com/apple/swift-docc-plugin", from: "1.0.0"),
    ]
)
