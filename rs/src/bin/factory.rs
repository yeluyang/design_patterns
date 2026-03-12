use std::sync::Arc;

use design_patterns::factory::{GameFactory, Factory, Player, Shop};

fn main() {
    // 构造
    let factory: Arc<dyn Factory> = Arc::new(GameFactory);

    // 注入
    let player = Player {
        factory: Arc::clone(&factory),
    };
    let shop = Shop {
        factory: Arc::clone(&factory),
    };

    // 使用
    player.equip(&["sword", "shield"]);
    shop.preview("sword");
}
