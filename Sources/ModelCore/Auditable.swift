import Foundation

public protocol Auditable {
    var createdAt : Date { get set }
    var createdBy : String { get set }
    var updatedAt : Date { get set }
    var updatedBy : String { get set }
    var deletedAt : Date? { get set }
    var deletedBy : String? { get set }
}
