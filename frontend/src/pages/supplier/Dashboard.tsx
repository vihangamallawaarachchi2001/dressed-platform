// src/pages/supplier/Dashboard.tsx
import DesignFeed from '../../components/supplier/DesignFeed';

const SupplierDashboard = () => {
  return (
    <div className="min-h-screen bg-gray-50 p-4 md:p-6">
      <div className="max-w-7xl mx-auto">
        <h1 className="text-2xl font-bold text-gray-900 mb-6">Supplier Dashboard</h1>
        <div className="bg-white p-6 rounded-xl shadow-md">
          <DesignFeed />
        </div>
      </div>
    </div>
  );
};

export default SupplierDashboard;