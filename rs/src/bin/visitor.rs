use std::sync::Arc;

use design_patterns::visitor::{AreaPrinter, Circle, Element, Rect, Scene, Visitor};

fn main() {
    // 构造
    let visitor: Arc<dyn Visitor> = Arc::new(AreaPrinter);

    // 注入
    let scene = Scene { visitor };

    // 使用
    let elements: Vec<Box<dyn Element>> = vec![
        Box::new(Circle { radius: 5.0 }),
        Box::new(Rect {
            width: 3.0,
            height: 4.0,
        }),
    ];

    scene.run(&elements);
}
