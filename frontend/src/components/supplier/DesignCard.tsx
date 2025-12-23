// src/components/supplier/DesignCard.tsx
import { useState } from 'react';
import QuoteModal from './QuoteModal';
import type { Design } from '../../types';

const DesignCard = ({ design }: { design: Design }) => {
  const [showModal, setShowModal] = useState(false);

  return (
    <div className="bg-white rounded-xl shadow-sm border border-gray-200 overflow-hidden">
      {/* Design Image Preview */}
      <div className="h-48 bg-gray-100 flex items-center justify-center">
        {design.filePath.endsWith('.pdf') ? (
          <div className="text-center">
            <svg xmlns="http://www.w3.org/2000/svg" className="h-12 w-12 text-gray-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
            </svg>
            <p className="mt-1 text-sm text-gray-600">PDF Design</p>
          </div>
        ) : (
          <img
            src={`http://localhost:8000${design.filePath}`}
            alt={design.title}
            className="h-full w-full object-contain"
            onError={(e) => (e.currentTarget.src = 'https://via.placeholder.com/150?text=No+Preview')}
          />
        )}
      </div>

      {/* Design Info */}
      <div className="p-4">
        <div className="flex justify-between items-start">
          <div>
            <h3 className="font-bold text-lg">{design.title}</h3>
            <span className="inline-block px-2 py-1 bg-blue-100 text-blue-800 text-xs rounded mt-1">
              {design.category}
            </span>
          </div>
        </div>
        <p className="text-gray-600 text-sm mt-2">{design.description}</p>
        <button
          onClick={() => setShowModal(true)}
          className="mt-4 w-full bg-emerald-600 text-white py-2 rounded hover:bg-emerald-700 transition"
        >
          Submit Quote
        </button>
      </div>

      {showModal && (
        <QuoteModal
          designId={design.id}
          designTitle={design.title}
          onClose={() => setShowModal(false)}
        />
      )}
    </div>
  );
};

export default DesignCard;