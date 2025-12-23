/* eslint-disable @typescript-eslint/no-explicit-any */
// src/components/designer/DesignCard.tsx
import { submitDesign } from '../../services/designService';
import { useState } from 'react';
import type { Design } from '../../types';

const DesignCard = ({ design }: { design: Design }) => {
  const [submitting, setSubmitting] = useState(false);
  const [submitted, setSubmitted] = useState(design.status === 'SUBMITTED');

  const handleSubmit = async () => {
    if (submitting || submitted) return;
    setSubmitting(true);
    try {
      await submitDesign(design.id);
      setSubmitted(true);
    } catch (err: any) {
      alert('Failed to submit design: ' + err.message);
    } finally {
      setSubmitting(false);
    }
  };

  return (
    <div className="border border-gray-200 rounded-lg overflow-hidden">
      <div className="p-4">
        <h3 className="font-bold text-lg">{design.title}</h3>
        <p className="text-gray-600 text-sm">{design.category}</p>
        <p className="mt-1 text-sm">{design.description}</p>
        <div className="mt-3 flex justify-between items-center">
          <span className={`px-2 py-1 text-xs rounded ${
            design.status === 'SUBMITTED'
              ? 'bg-green-100 text-green-800'
              : 'bg-yellow-100 text-yellow-800'
          }`}>
            {design.status}
          </span>
          {design.status === 'DRAFT' && (
            <button
              onClick={handleSubmit}
              disabled={submitting}
              className="text-sm bg-blue-600 text-white px-3 py-1 rounded hover:bg-blue-700 disabled:opacity-50"
            >
              {submitting ? 'Submitting...' : 'Submit'}
            </button>
          )}
        </div>
      </div>
    </div>
  );
};

export default DesignCard;