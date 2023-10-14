use std::time::Instant;

#[derive(Debug, PartialEq, Eq, Default)]
/// Base abstract model object.
pub struct BaseModel {}

#[derive(Debug, PartialEq, Eq, Default)]
pub struct SimpleModel {
    id: String,
}

#[derive(Debug, PartialEq, Eq)]
/// Model object.
/// Contains ID and audit fields.
pub struct Model {
    id: String,
    created_at: Instant,
    created_by: String,
    updated_at: Instant,
    updated_by: String,
    deleted_at: Option<Instant>,
    deleted_by: Option<String>,
}

#[derive(Debug, PartialEq, Eq, Default)]
pub struct EmbeddedModel {}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::constants::*;

    #[test]
    fn test_base_model() {
        let _ = BaseModel {};
        // assert_eq!(bm, 1);
    }

    #[test]
    fn test_simple_model() {
        let sm = SimpleModel {
            id: String::from("1"),
        };
        assert_eq!(sm.id, "1");
    }

    #[test]
    fn test_model() {
        let m = Model {
            id: String::from("1"),
            created_at: Instant::now(),
            created_by: String::from(SYSTEM_USER_ID),
            updated_at: Instant::now(),
            updated_by: String::from(SYSTEM_USER_ID),
            deleted_at: None,
            deleted_by: None,
        };

        assert_eq!(m.id, "1");
    }
}
