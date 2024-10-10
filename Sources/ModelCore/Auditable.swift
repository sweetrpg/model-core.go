import Foundation

public protocol Auditable {
    var createdAt : Date
    var createdBy : String
    var updatedAt : Date
    var updatedBy : String
    var deletedAt : Date?
    var deletedBy : String?
}
