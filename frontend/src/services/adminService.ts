// src/services/adminService.ts
const API_BASE = '/api';

// Since we don't have a dedicated admin service, we'll fetch from existing services
export const fetchAdminStats = async () => {
  const token = localStorage.getItem('token');
  const headers = { Authorization: `Bearer ${token}` };

  try {
    // Fetch from multiple services
    const [usersRes, designsRes, quotesRes, ordersRes] = await Promise.all([
      // Total users (designers + suppliers)
      fetch(`${API_BASE}/auth/users`, { headers }), // ← You'll add this endpoint
      fetch(`${API_BASE}/designs`, { headers }),    // ← List all designs
      fetch(`${API_BASE}/quotes`, { headers }),     // ← List all quotes
      fetch(`${API_BASE}/orders`, { headers }),     // ← List all orders
    ]);

    const users = usersRes.ok ? await usersRes.json() : [];
    const designs = designsRes.ok ? await designsRes.json() : [];
    const quotes = quotesRes.ok ? await quotesRes.json() : [];
    const orders = ordersRes.ok ? await ordersRes.json() : [];

    return {
      totalUsers: users.length,
      totalDesigns: designs.length,
      totalQuotes: quotes.length,
      totalOrders: orders.length,
      designs: designs.slice(0, 5),   // Recent 5
      quotes: quotes.slice(0, 5),
      orders: orders.slice(0, 5),
    };
  } catch (err) {
    console.error('Failed to fetch admin data:', err);
    return {
      totalUsers: 0,
      totalDesigns: 0,
      totalQuotes: 0,
      totalOrders: 0,
      designs: [],
      quotes: [],
      orders: [],
    };
  }
};