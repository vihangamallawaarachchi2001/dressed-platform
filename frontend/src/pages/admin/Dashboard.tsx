/* eslint-disable @typescript-eslint/no-explicit-any */
// src/pages/admin/Dashboard.tsx
import { useState, useEffect } from 'react';
import { fetchAdminStats } from '../../services/adminService';

interface AdminStats {
  totalUsers: number;
  totalDesigns: number;
  totalQuotes: number;
  totalOrders: number;
  designs: any[];
  quotes: any[];
  orders: any[];
}

const StatCard = ({ title, value }: { title: string; value: number }) => (
  <div className="bg-white p-6 rounded-xl shadow-sm border border-gray-200">
    <p className="text-gray-600">{title}</p>
    <p className="text-3xl font-bold text-indigo-700">{value}</p>
  </div>
);

const AdminDashboard = () => {
  const [stats, setStats] = useState<AdminStats>({
    totalUsers: 0,
    totalDesigns: 0,
    totalQuotes: 0,
    totalOrders: 0,
    designs: [],
    quotes: [],
    orders: [],
  });
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const loadStats = async () => {
      const data = await fetchAdminStats();
      setStats(data);
      setLoading(false);
    };
    loadStats();
  }, []);

  return (
    <div className="min-h-screen bg-gray-50 p-4 md:p-6">
      <div className="max-w-7xl mx-auto">
        <h1 className="text-2xl font-bold text-gray-900 mb-6">Dressed™ Admin Portal</h1>

        {/* Stats Overview */}
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
          <StatCard title="Total Users" value={stats.totalUsers} />
          <StatCard title="Total Designs" value={stats.totalDesigns} />
          <StatCard title="Total Quotes" value={stats.totalQuotes} />
          <StatCard title="Total Orders" value={stats.totalOrders} />
        </div>

        {/* Recent Activity */}
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
          {/* Recent Designs */}
          <div className="bg-white p-6 rounded-xl shadow-sm">
            <h2 className="text-lg font-semibold mb-4">Recent Designs</h2>
            {loading ? (
              <p>Loading...</p>
            ) : stats.designs.length === 0 ? (
              <p className="text-gray-500">No designs yet</p>
            ) : (
              <ul className="space-y-3">
                {stats.designs.map((d: any) => (
                  <li key={d.id} className="border-b pb-2">
                    <div className="font-medium">{d.title}</div>
                    <div className="text-sm text-gray-600">{d.category} • {d.status}</div>
                  </li>
                ))}
              </ul>
            )}
          </div>

          {/* Recent Quotes */}
          <div className="bg-white p-6 rounded-xl shadow-sm">
            <h2 className="text-lg font-semibold mb-4">Recent Quotes</h2>
            {loading ? (
              <p>Loading...</p>
            ) : stats.quotes.length === 0 ? (
              <p className="text-gray-500">No quotes yet</p>
            ) : (
              <ul className="space-y-3">
                {stats.quotes.map((q: any) => (
                  <li key={q.id} className="border-b pb-2">
                    <div className="font-medium">${q.price.toFixed(2)}</div>
                    <div className="text-sm text-gray-600">ETA: {q.eta_days} days • {q.status}</div>
                  </li>
                ))}
              </ul>
            )}
          </div>

          {/* Recent Orders */}
          <div className="bg-white p-6 rounded-xl shadow-sm">
            <h2 className="text-lg font-semibold mb-4">Recent Orders</h2>
            {loading ? (
              <p>Loading...</p>
            ) : stats.orders.length === 0 ? (
              <p className="text-gray-500">No orders yet</p>
            ) : (
              <ul className="space-y-3">
                {stats.orders.map((o: any) => (
                  <li key={o.id} className="border-b pb-2">
                    <div className="font-medium">Order #{o.id.substring(0, 6)}</div>
                    <div className="text-sm text-gray-600">{o.status}</div>
                  </li>
                ))}
              </ul>
            )}
          </div>
        </div>
      </div>
    </div>
  );
};

export default AdminDashboard;