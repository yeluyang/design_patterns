use std::sync::Arc;

// --- Element: 两个实现展示双重分发 ---

pub trait Element {
    fn accept(&self, visitor: &dyn Visitor);
}

pub struct Circle {
    pub radius: f64,
}
pub struct Rect {
    pub width: f64,
    pub height: f64,
}

impl Element for Circle {
    fn accept(&self, visitor: &dyn Visitor) {
        visitor.visit_circle(self);
    }
}

impl Element for Rect {
    fn accept(&self, visitor: &dyn Visitor) {
        visitor.visit_rect(self);
    }
}

// --- Visitor: 一个实现足够展示多态 ---

pub trait Visitor {
    fn visit_circle(&self, circle: &Circle);
    fn visit_rect(&self, rect: &Rect);
}

pub struct AreaPrinter;

impl Visitor for AreaPrinter {
    fn visit_circle(&self, c: &Circle) {
        println!(
            "⭕ Circle area: {:.2}",
            std::f64::consts::PI * c.radius * c.radius
        );
    }
    fn visit_rect(&self, r: &Rect) {
        println!("🟦 Rect area: {:.2}", r.width * r.height);
    }
}

// --- 消费端: 持有 visitor，由外部注入 ---

pub struct Scene {
    pub visitor: Arc<dyn Visitor>,
}

impl Scene {
    pub fn run(&self, elements: &[Box<dyn Element>]) {
        for e in elements {
            e.accept(self.visitor.as_ref());
        }
    }
}
