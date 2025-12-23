/* eslint-disable @typescript-eslint/no-explicit-any */
// src/components/supplier/DesignFeed.tsx
import { useState, useEffect } from 'react';
import { getPublicDesigns } from '../../services/supplierService';
import DesignCard from './DesignCard';
import type { Design } from '../../types';

const DesignFeed = () => {
  const [designs, setDesigns] = useState<Design[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchDesigns = async () => {
      try {
        const data = await getPublicDesigns();
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
    <div>
      <h2 className="text-xl font-semibold mb-4">Available Designs</h2>
      {loading ? (
        <p>Loading designs...</p>
      ) : designs.length === 0 ? (
        <p className="text-gray-500">No designs available yet.</p>
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {designs.map((design) => (
            <DesignCard key={design.id} design={design} />
          ))}
        </div>
      )}
    </div>
  );
};

export default DesignFeed;