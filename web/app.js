// API Configuration
const API_BASE_URL = 'http://localhost:8080/api';

// State Management
let currentUser = {
    account: null,
    profile: null
};

// DOM Elements
const authSection = document.getElementById('authSection');
const appSection = document.getElementById('appSection');
const userInfo = document.getElementById('userInfo');
const welcomeMessage = document.getElementById('welcomeMessage');

// Initialize App
document.addEventListener('DOMContentLoaded', () => {
    initializeEventListeners();
    checkSession();
});

// Event Listeners
function initializeEventListeners() {
    // Auth tabs
    document.querySelectorAll('.tab-btn').forEach(btn => {
        btn.addEventListener('click', (e) => {
            switchAuthTab(e.target.dataset.tab);
        });
    });

    // Forms
    document.getElementById('loginForm').addEventListener('submit', handleLogin);
    document.getElementById('registerForm').addEventListener('submit', handleRegister);
    document.getElementById('profileForm').addEventListener('submit', handleProfileUpdate);
    document.getElementById('logoutBtn').addEventListener('click', handleLogout);

    // App tabs
    document.querySelectorAll('.app-tab-btn').forEach(btn => {
        btn.addEventListener('click', (e) => {
            switchAppSection(e.target.dataset.section);
        });
    });

    // Modal close
    document.querySelector('.close').addEventListener('click', () => {
        document.getElementById('menuModal').style.display = 'none';
    });

    window.addEventListener('click', (e) => {
        const modal = document.getElementById('menuModal');
        if (e.target === modal) {
            modal.style.display = 'none';
        }
    });
}

// Auth Tab Switching
function switchAuthTab(tab) {
    document.querySelectorAll('.tab-btn').forEach(btn => {
        btn.classList.remove('active');
    });
    document.querySelectorAll('.tab-content').forEach(content => {
        content.classList.remove('active');
    });

    event.target.classList.add('active');
    document.getElementById(`${tab}Tab`).classList.add('active');
}

// App Section Switching
function switchAppSection(section) {
    document.querySelectorAll('.app-tab-btn').forEach(btn => {
        btn.classList.remove('active');
    });
    document.querySelectorAll('.app-section').forEach(sec => {
        sec.classList.remove('active');
    });

    event.target.classList.add('active');
    document.getElementById(`${section}Section`).classList.add('active');

    // Load data based on section
    if (section === 'restaurants') {
        loadRestaurants();
    } else if (section === 'orders') {
        loadOrders();
    } else if (section === 'profile') {
        loadProfile();
    }
}

// Session Management
function checkSession() {
    const savedAccount = localStorage.getItem('account');
    const savedProfile = localStorage.getItem('profile');

    if (savedAccount && savedProfile) {
        currentUser.account = JSON.parse(savedAccount);
        currentUser.profile = JSON.parse(savedProfile);
        showApp();
    }
}

function saveSession() {
    localStorage.setItem('account', JSON.stringify(currentUser.account));
    localStorage.setItem('profile', JSON.stringify(currentUser.profile));
}

function clearSession() {
    localStorage.removeItem('account');
    localStorage.removeItem('profile');
    currentUser = { account: null, profile: null };
}

// Auth Handlers
async function handleLogin(e) {
    e.preventDefault();
    const email = document.getElementById('loginEmail').value;
    const password = document.getElementById('loginPassword').value;

    try {
        const response = await fetch(`${API_BASE_URL}/accounts/login`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ email, password })
        });

        if (!response.ok) {
            throw new Error('Invalid credentials');
        }

        const data = await response.json();
        currentUser.account = data.account;

        // Fetch user profile
        const profileResponse = await fetch(`${API_BASE_URL}/users/account/${data.account.id}`);
        if (profileResponse.ok) {
            currentUser.profile = await profileResponse.json();
            saveSession();
            showToast('Login successful!', 'success');
            showApp();
        } else {
            throw new Error('Profile not found');
        }
    } catch (error) {
        showToast(error.message, 'error');
    }
}

async function handleRegister(e) {
    e.preventDefault();
    const email = document.getElementById('registerEmail').value;
    const password = document.getElementById('registerPassword').value;
    const name = document.getElementById('registerName').value;
    const address = document.getElementById('registerAddress').value;

    try {
        // Create account
        const accountResponse = await fetch(`${API_BASE_URL}/accounts`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ email, password })
        });

        if (!accountResponse.ok) {
            throw new Error('Failed to create account');
        }

        const accountData = await accountResponse.json();

        // Create user profile
        const userResponse = await fetch(`${API_BASE_URL}/users`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
                account_id: accountData.id,
                name,
                address
            })
        });

        if (!userResponse.ok) {
            throw new Error('Failed to create user profile');
        }

        const userData = await userResponse.json();
        
        currentUser.account = accountData;
        currentUser.profile = userData;
        saveSession();
        
        showToast('Registration successful!', 'success');
        showApp();
    } catch (error) {
        showToast(error.message, 'error');
    }
}

function handleLogout() {
    clearSession();
    authSection.style.display = 'block';
    appSection.style.display = 'none';
    userInfo.style.display = 'none';
    
    // Reset forms
    document.getElementById('loginForm').reset();
    document.getElementById('registerForm').reset();
    
    showToast('Logged out successfully', 'success');
}

// App Display
function showApp() {
    authSection.style.display = 'none';
    appSection.style.display = 'block';
    userInfo.style.display = 'flex';
    welcomeMessage.textContent = `Welcome, ${currentUser.profile.name}!`;
    
    loadRestaurants();
}

// Load Restaurants
async function loadRestaurants() {
    try {
        const response = await fetch(`${API_BASE_URL}/restaurants`);
        const restaurants = await response.json();
        
        const restaurantsGrid = document.getElementById('restaurantsList');
        restaurantsGrid.innerHTML = '';
        
        restaurants.forEach(restaurant => {
            const card = document.createElement('div');
            card.className = 'restaurant-card';
            card.innerHTML = `
                <h3>${restaurant.name}</h3>
                <p>📍 ${restaurant.address}</p>
                <span class="cuisine">${restaurant.cuisine}</span>
            `;
            card.addEventListener('click', () => showMenu(restaurant));
            restaurantsGrid.appendChild(card);
        });
    } catch (error) {
        showToast('Failed to load restaurants', 'error');
    }
}

// Show Menu
async function showMenu(restaurant) {
    try {
        const response = await fetch(`${API_BASE_URL}/restaurants/${restaurant.id}/foods`);
        const foods = await response.json();
        
        document.getElementById('restaurantName').textContent = restaurant.name;
        document.getElementById('restaurantDetails').textContent = 
            `${restaurant.cuisine} • ${restaurant.address}`;
        
        const foodsList = document.getElementById('foodsList');
        foodsList.innerHTML = '';
        
        if (foods && foods.length > 0) {
            foods.forEach(food => {
                const foodItem = document.createElement('div');
                foodItem.className = 'food-item';
                foodItem.innerHTML = `
                    <div class="food-info">
                        <h4>${food.name}</h4>
                        <span class="category">${food.category}</span>
                    </div>
                    <div style="display: flex; align-items: center;">
                        <span class="food-price">$${food.price.toFixed(2)}</span>
                        <button class="btn btn-success btn-sm" onclick="orderFood(${food.id}, ${restaurant.id}, ${food.price}, '${food.name}')">
                            Order
                        </button>
                    </div>
                `;
                foodsList.appendChild(foodItem);
            });
        } else {
            foodsList.innerHTML = '<div class="empty-state"><h3>No menu items available</h3></div>';
        }
        
        document.getElementById('menuModal').style.display = 'block';
    } catch (error) {
        showToast('Failed to load menu', 'error');
    }
}

// Order Food
async function orderFood(foodId, restaurantId, price, foodName) {
    try {
        const response = await fetch(`${API_BASE_URL}/orders`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
                account_id: currentUser.account.id,
                food_id: foodId,
                restaurant_id: restaurantId,
                total_price: price
            })
        });

        if (!response.ok) {
            throw new Error('Failed to place order');
        }

        showToast(`Order placed for ${foodName}!`, 'success');
        document.getElementById('menuModal').style.display = 'none';
    } catch (error) {
        showToast(error.message, 'error');
    }
}

// Load Orders
async function loadOrders() {
    try {
        const response = await fetch(`${API_BASE_URL}/orders/account/${currentUser.account.id}`);
        const orders = await response.json();
        
        const ordersList = document.getElementById('ordersList');
        ordersList.innerHTML = '';
        
        if (orders && orders.length > 0) {
            orders.forEach(order => {
                const orderCard = document.createElement('div');
                orderCard.className = 'order-card';
                
                const date = new Date(order.created_at);
                const formattedDate = date.toLocaleDateString() + ' ' + date.toLocaleTimeString();
                
                orderCard.innerHTML = `
                    <div class="order-header">
                        <span class="order-id">Order ID: ${order.id}</span>
                        <span class="order-date">${formattedDate}</span>
                    </div>
                    <div class="order-details">
                        <div class="order-detail-item">
                            <span class="order-detail-label">Food ID</span>
                            <span class="order-detail-value">#${order.food_id}</span>
                        </div>
                        <div class="order-detail-item">
                            <span class="order-detail-label">Restaurant ID</span>
                            <span class="order-detail-value">#${order.restaurant_id}</span>
                        </div>
                        <div class="order-detail-item">
                            <span class="order-detail-label">Total Price</span>
                            <span class="order-detail-value order-total">$${order.total_price.toFixed(2)}</span>
                        </div>
                    </div>
                `;
                ordersList.appendChild(orderCard);
            });
        } else {
            ordersList.innerHTML = `
                <div class="empty-state">
                    <h3>No orders yet</h3>
                    <p>Start ordering from your favorite restaurants!</p>
                </div>
            `;
        }
    } catch (error) {
        showToast('Failed to load orders', 'error');
    }
}

// Load Profile
function loadProfile() {
    document.getElementById('profileName').value = currentUser.profile.name;
    document.getElementById('profileAddress').value = currentUser.profile.address;
    document.getElementById('profileEmail').value = currentUser.account.email;
}

// Update Profile
async function handleProfileUpdate(e) {
    e.preventDefault();
    const name = document.getElementById('profileName').value;
    const address = document.getElementById('profileAddress').value;

    try {
        const response = await fetch(`${API_BASE_URL}/users/${currentUser.profile.id}`, {
            method: 'PUT',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ name, address })
        });

        if (!response.ok) {
            throw new Error('Failed to update profile');
        }

        const updatedProfile = await response.json();
        currentUser.profile = updatedProfile;
        saveSession();
        welcomeMessage.textContent = `Welcome, ${currentUser.profile.name}!`;
        
        showToast('Profile updated successfully!', 'success');
    } catch (error) {
        showToast(error.message, 'error');
    }
}

// Toast Notification
function showToast(message, type = 'success') {
    const toast = document.getElementById('toast');
    toast.textContent = message;
    toast.className = `toast ${type} show`;
    
    setTimeout(() => {
        toast.classList.remove('show');
    }, 3000);
}
