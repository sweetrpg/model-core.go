import Foundation

public protocol Auditable {
    var createdAt : Date { get set }
    var createdBy : URL { get set }
    var updatedAt : Date { get set }
    var updatedBy : URL { get set }
    var deletedAt : Date? { get set }
    var deletedBy : URL? { get set }
}
