/* eslint-disable @typescript-eslint/no-explicit-any */
// src/pages/designer/Dashboard.tsx
import { useState, useEffect } from 'react';
import { getMyDesigns } from '../../services/designService';
import DesignUploadForm from '../../components/designer/DesignUploadForm';
import DesignCard from '../../components/designer/DesignCard';
import QuoteList from '../../components/designer/QuoteList';
import type { Design } from '../../types';

const DesignerDashboard = () => {
  const [designs, setDesigns] = useState<Design[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchDesigns = async () => {
      try {
        const data = await getMyDesigns();
        setDesigns(data);
      } catch (err: any) {
        alert('Failed to load designs: ' + err.message);
      } finally {
        setLoading(false);
      }
    };
    fetchDesigns();
  }, []);

  return (
    <div className="min-h-screen bg-gray-50 p-4 md:p-6">
      <div className="max-w-7xl mx-auto">
        <h1 className="text-2xl font-bold text-gray-900 mb-6">Designer Dashboard</h1>

        <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
          {/* Upload Form - Full width on mobile, 1/3 on desktop */}
          <div className="lg:col-span-1">
            <DesignUploadForm />
          </div>

          {/* Designs & Quotes - 2/3 on desktop */}
          <div className="lg:col-span-2 space-y-6">
            {/* My Designs */}
            <div>
              <h2 className="text-xl font-semibold mb-4">My Designs</h2>
              {loading ? (
                <p>Loading designs...</p>
              ) : designs.length === 0 ? (
                <p className="text-gray-500">No designs yet. Upload your first design!</p>
              ) : (
                <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                  {designs.map((design) => (
                    <DesignCard key={design.id} design={design} />
                  ))}
                </div>
              )}
            </div>

            {/* Quotes */}
            <QuoteList />
          </div>
        </div>
      </div>
    </div>
  );
};

export default DesignerDashboard;