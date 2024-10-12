import Testing
import Foundation
@testable import ModelCore

@Suite("Test Auditable struct", .serialized)
struct AuditiableTests {
    struct TestAuditable : Auditable {
        var createdAt : Date
        var createdBy : URL
        var updatedAt : Date
        var updatedBy : URL
        var deletedAt : Date?
        var deletedBy : URL?
    }

    @Test("Test properties")
    func testProperties() async throws {
        let now = Date()
        let user = URL(string: "tester")!
        let auditable = TestAuditable(createdAt: now, createdBy: user,
                                      updatedAt: now, updatedBy: user)

        #expect(auditable != nil)
        #expect(auditable.createdAt == now)
        #expect(auditable.createdBy == user)
        #expect(auditable.updatedAt == now)
        #expect(auditable.updatedBy == user)
        #expect(auditable.deletedAt == nil)
        #expect(auditable.deletedBy == nil)
    }
}
