db.createUser(
    {
        user: 'user',
        pwd: 'password',
        roles: [
            {role: 'readWrite', db: 'your_app'},
            {role: 'readWrite', db: 'test_your_app'}
        ],
    },
);