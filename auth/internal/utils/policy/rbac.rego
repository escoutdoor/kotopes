package auth

default allow = false

permissions := {
    "admin": [
        {
            "path": "/*",
            "methods": ["GET", "POST", "PUT", "DELETE", "PATCH"]
        }
    ],
    "user": [
        {
            "path": "/pets/v1/*",
            "methods": ["PATCH", "DELETE"]
        },
        {
            "path": "/pets/v1/",
            "methods": ["POST"]
        },
        {
            "path": "/users/v1/",
            "methods": ["PATCH", "DELETE"]
        },
        {
            "path": "/favorites/v1/*",
            "methods": ["POST", "DELETE"]
        },
        {
            "path": "/favorites/v1/",
            "methods": ["GET"]
        }
    ]
}

allow {
    role_permissions := permissions[input.role]
    
    some_permission_matches(role_permissions)
}

some_permission_matches(permissions) {
    permission := permissions[_]

    glob.match(permission.path, ["/"], input.endpoint)

    input.method == permission.methods[_]
}
