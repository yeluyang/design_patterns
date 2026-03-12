use std::sync::Arc;

pub trait Product {
    fn describe(&self);
}

struct Sword;
struct Shield;

impl Product for Sword {
    fn describe(&self) {
        println!("🗡️ Sword");
    }
}
impl Product for Shield {
    fn describe(&self) {
        println!("🛡️ Shield");
    }
}

pub trait Factory {
    fn create(&self, name: &str) -> Option<Box<dyn Product>>;
}

pub struct GameFactory;

impl Factory for GameFactory {
    fn create(&self, name: &str) -> Option<Box<dyn Product>> {
        match name {
            "sword" => Some(Box::new(Sword)),
            "shield" => Some(Box::new(Shield)),
            _ => None,
        }
    }
}

pub struct Player {
    pub factory: Arc<dyn Factory>,
}

impl Player {
    pub fn equip(&self, items: &[&str]) {
        for name in items {
            match self.factory.create(name) {
                Some(p) => p.describe(),
                None => println!("❓ Unknown: {name}"),
            }
        }
    }
}

pub struct Shop {
    pub factory: Arc<dyn Factory>,
}

impl Shop {
    pub fn preview(&self, name: &str) {
        match self.factory.create(name) {
            Some(p) => {
                print!("Preview: ");
                p.describe();
            }
            None => println!("Not in stock: {name}"),
        }
    }
}
