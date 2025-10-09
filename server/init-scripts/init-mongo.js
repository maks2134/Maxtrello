db.createUser({
    user: "app_user",
    pwd: "app_password",
    roles: [
        {
            role: "readWrite",
            db: "boards"
        }
    ]
});

db.createCollection("boards");
db.createCollection("columns");

db.boards.createIndex({ "owner_id": 1 });
db.boards.createIndex({ "members.user_id": 1 });
db.boards.createIndex({ "created_at": -1 });

db.columns.createIndex({ "board_id": 1 });
db.columns.createIndex({ "position": 1 });

db.boards.insertMany([
    {
        _id: "111111111111111111111111",
        title: "Проект Разработка",
        description: "Доска для управления разработкой проекта",
        owner_id: "000000000000000000000001",
        members: [
            {
                user_id: "000000000000000000000001",
                role: "owner",
                joined_at: new Date()
            }
        ],
        settings: {
            color: "#3498db",
            is_public: false
        },
        created_at: new Date(),
        updated_at: new Date()
    },
    {
        _id: "222222222222222222222222",
        title: "Личные задачи",
        description: "Мои персональные задачи",
        owner_id: "000000000000000000000001",
        members: [
            {
                user_id: "000000000000000000000001",
                role: "owner",
                joined_at: new Date()
            }
        ],
        settings: {
            color: "#2ecc71",
            is_public: false
        },
        created_at: new Date(),
        updated_at: new Date()
    }
]);

db.columns.insertMany([
    {
        _id: "333333333333333333333333",
        title: "Бэклог",
        position: 0,
        board_id: "111111111111111111111111",
        created_at: new Date(),
        updated_at: new Date()
    },
    {
        _id: "444444444444444444444444",
        title: "В работе",
        position: 1,
        board_id: "111111111111111111111111",
        created_at: new Date(),
        updated_at: new Date()
    },
    {
        _id: "555555555555555555555555",
        title: "На проверке",
        position: 2,
        board_id: "111111111111111111111111",
        created_at: new Date(),
        updated_at: new Date()
    },
    {
        _id: "666666666666666666666666",
        title: "Готово",
        position: 3,
        board_id: "111111111111111111111111",
        created_at: new Date(),
        updated_at: new Date()
    }
]);