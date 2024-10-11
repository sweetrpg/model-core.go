import XCTest
import Foundation
@testable import ModelCore

final class AuditiableTests: XCTestCase {
    struct TestAuditable : Auditable {
        var createdAt : Date
        var createdBy : URL
        var updatedAt : Date
        var updatedBy : URL
        var deletedAt : Date?
        var deletedBy : URL?
    }

    func testAuditable() throws {
        // This is an example of a functional test case.
        // Use XCTAssert and related functions to verify your tests produce the correct
        // results.
        // XCTAssertEqual(ModelCore().text, "Hello, World!")
        let now = Date()
        let user = URL(string: "tester")!
        let auditable = TestAuditable(createdAt: now, createdBy: user,
                                      updatedAt: now, updatedBy: user)

        XCTAssertNotNil(auditable)
        XCTAssertEqual(auditable.createdAt, now)
        XCTAssertEqual(auditable.createdBy, user)
        XCTAssertEqual(auditable.updatedAt, now)
        XCTAssertEqual(auditable.updatedBy, user)
        XCTAssertNil(auditable.deletedAt)
        XCTAssertNil(auditable.deletedBy)
    }
}
