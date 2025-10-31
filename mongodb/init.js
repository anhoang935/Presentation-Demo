// MongoDB Initialization Script
// Run this script to set up the MongoDB database and collections

// Switch to or create the demo_db database
db = db.getSiblingDB('demo_db');

// Create the orders collection with validation schema
db.createCollection("orders", {
    validator: {
        $jsonSchema: {
            bsonType: "object",
            required: ["account_id", "food_id", "restaurant_id", "total_price", "created_at"],
            properties: {
                account_id: {
                    bsonType: "int",
                    description: "Account ID must be an integer and is required"
                },
                food_id: {
                    bsonType: "int",
                    description: "Food ID must be an integer and is required"
                },
                restaurant_id: {
                    bsonType: "int",
                    description: "Restaurant ID must be an integer and is required"
                },
                total_price: {
                    bsonType: "double",
                    minimum: 0,
                    description: "Total price must be a positive number and is required"
                },
                created_at: {
                    bsonType: "date",
                    description: "Created at must be a date and is required"
                }
            }
        }
    }
});

// Create indexes for better query performance
db.orders.createIndex({ "account_id": 1 });
db.orders.createIndex({ "restaurant_id": 1 });
db.orders.createIndex({ "created_at": -1 });
db.orders.createIndex({ "account_id": 1, "created_at": -1 });

// Insert sample orders for testing (optional)
// Uncomment the lines below to add test orders
/*
db.orders.insertMany([
    {
        account_id: 1,
        food_id: 1,
        restaurant_id: 1,
        total_price: 12.99,
        created_at: new Date()
    },
    {
        account_id: 1,
        food_id: 3,
        restaurant_id: 2,
        total_price: 8.99,
        created_at: new Date()
    }
]);
*/

// Display collection information
print("\n=== MongoDB Database Setup Complete ===");
print("Database: demo_db");
print("\nCollections:");
db.getCollectionNames().forEach(function(collection) {
    print("  - " + collection);
});

print("\nIndexes on 'orders' collection:");
db.orders.getIndexes().forEach(function(index) {
    print("  - " + index.name);
});

print("\nDocument count in 'orders':");
print("  " + db.orders.countDocuments({}));

print("\n=== Setup Complete ===\n");
